package duckdb

/*
#include <duckdb_extension.h>
#include <magic.h>
*/
import "C"
import (
	"unsafe"
)

type API struct {
	ptr    *C.duckdb_ext_api_v0
	info   C.duckdb_extension_info
	access *C._duckdb_extension_access
}

func (api *API) Database() Database {
	db := C._get_database(api.access.get_database, api.info)
	return Database{c: *db}
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

	C.duckdb_ext_api = (*C.duckdb_ext_api_v0)(C._get_api(access.get_api, info, version))
	if C.duckdb_ext_api == nil {
		api.SetError(getAPIError.Error())
		return api, getAPIError
	}

	api.ptr = C.duckdb_ext_api
	return api, nil
}
