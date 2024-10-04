package duckdb

/*
#include <go_duckdb.h>
*/
import "C"
import "unsafe"

type DeleteCallbackT struct {
	c C.duckdb_delete_callback_t
}

type ScalarFunctionT struct {
	c C.duckdb_scalar_function_t
}

func (f *ScalarFunctionT) Set(ptr unsafe.Pointer) {
	f.c = C.duckdb_scalar_function_t(ptr)
}
