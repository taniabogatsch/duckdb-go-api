package duckdb

/*
#include <stdlib.h>

// TODO: we do not need the functions in duckdb.h, and we need to add additional magic back into the generator script (ifdefs, etc.).
#include <duckdb_extension.h>
#include <magic.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var errState = errors.New("DuckDB Error")

type API struct {
	ptr    unsafe.Pointer
	info   C.duckdb_extension_info
	access *C._duckdb_extension_access
}

func (api *API) Close() {}

type DB struct {
	db *C.duckdb_database
}

func (api *API) DB() DB {
	db := C._get_database(api.access.get_database, api.info)
	return DB{db: db}
}

type Conn struct {
	conn C.duckdb_connection
}

func (c *Conn) Close() {
	C.duckdb_disconnect(&c.conn)
}

func (api *API) Connect(db DB, conn Conn) error {
	if C.duckdb_connect(*db.db, &conn.conn) == C.DuckDBError {
		return errState
	}
	return nil
}

func (api *API) SetError(msg string) {
	err := C.CString(msg)
	C._set_error(api.access.set_error, api.info, err)
	C.free(unsafe.Pointer(err))
}

func InitAPI(minVersion string, infoPtr unsafe.Pointer, accessPtr unsafe.Pointer) (API, error) {
	version := C.CString(minVersion)
	defer C.free(unsafe.Pointer(version))

	info := (C.duckdb_extension_info)(infoPtr)
	access := (*C._duckdb_extension_access)(accessPtr)
	api := API{info: info, access: access}

	C.duckdb_ext_api = (*C.duckdb_ext_api_v0)(C._get_api(access.get_api, info, version))
	if C.duckdb_ext_api == nil {
		api.SetError("failed to get the API")
		return api, errors.New("failed to get the API")
	}

	fmt.Println("got the API")
	return api, nil
}
