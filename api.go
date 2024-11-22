package duckdb

/*
#cgo CFLAGS: -DDUCKDB_API_EXCLUDE_FUNCTIONS=0 -DDUCKDB_EXTENSION_API_VERSION_DEV=1
#include <string.h>
#include <duckdb_go_extension.h>
#include <magic.h>
*/
import "C"
import (
	"unsafe"
)

const apiSize = unsafe.Sizeof([1]C.duckdb_ext_api_v0{})

var apiPtr unsafe.Pointer

type API struct {
	info   C.duckdb_extension_info
	access *C._duckdb_extension_access
}

func (api *API) Database() Database {
	db := C._get_database(api.access.get_database, api.info)
	return Database{unsafe.Pointer(db)}
}

func (api *API) SetError(msg string) {
	err := C.CString(msg)
	defer C.duckdb_free(unsafe.Pointer(err))
	C._set_error(api.access.set_error, api.info, err)
}

func Init(minVersion string, infoPtr unsafe.Pointer, accessPtr unsafe.Pointer) (API, error) {
	version := C.CString(minVersion)
	defer C.duckdb_free(unsafe.Pointer(version))

	info := (C.duckdb_extension_info)(infoPtr)
	access := (*C._duckdb_extension_access)(accessPtr)
	api := API{info: info, access: access}

	ptr := (*C.duckdb_ext_api_v0)(C._get_api(access.get_api, info, version))
	if ptr == nil {
		api.SetError(getAPIError.Error())
		return api, getAPIError
	}

	apiPtr = unsafe.Pointer(C.malloc(C.size_t(apiSize)))
	C.memcpy(apiPtr, unsafe.Pointer(ptr), C.size_t(apiSize))
	C.duckdb_ext_api = (*C.duckdb_ext_api_v0)(apiPtr)
	return api, nil
}
