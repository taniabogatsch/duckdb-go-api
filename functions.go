package duckdb

/*
#include <go_duckdb.h>
#include <duckdb_extension.h>
*/
import "C"
import "unsafe"

// TODO: should I always pass the custom types as pointers?

func Open(path string, db *Database) State {
	cPath := C.CString(path)
	defer C.duckdb_free(unsafe.Pointer(cPath))
	return State(C.duckdb_open(cPath, &db.c))
}

func OpenExt(path string, db *Database, config Config, errMsg *string) State {
	cPath := C.CString(path)
	defer C.duckdb_free(unsafe.Pointer(cPath))
	var err *C.char
	defer C.duckdb_free(unsafe.Pointer(err))

	state := State(C.duckdb_open_ext(cPath, &db.c, config.c, &err))
	if state == STATE_ERROR {
		*errMsg = C.GoString(err)
	}
	return state
}

func Close(db Database) {
	C.duckdb_close(&db.c)
}

func Connect(db Database, conn *Connection) State {
	return State(C.duckdb_connect(db.c, &conn.c))
}

func Interrupt(conn Connection) {
	C.duckdb_interrupt(conn.c)
}

func QueryProgress(conn Connection) QueryProgressType {
	return QueryProgressType{C.duckdb_query_progress(conn.c)}
}

func Disconnect(conn *Connection) {
	C.duckdb_disconnect(&conn.c)
}

func LibraryVersion() string {
	cStr := C.duckdb_library_version()
	return C.GoString(cStr)
}

// Types

func CreateLogicalType(t Type) LogicalType {
	return LogicalType{C.duckdb_create_logical_type(C.duckdb_type(t))}
}

func DestroyLogicalType(t *LogicalType) {
	C.duckdb_destroy_logical_type(&t.c)
}

// Data chunk

func FetchChunk(result Result) DataChunk {
	return DataChunk{C.duckdb_fetch_chunk(result.c)}
}

func CreateDataChunk(types []LogicalType, count uint64) DataChunk {
	var dummy C.duckdb_logical_type
	size := C.size_t(unsafe.Sizeof(dummy))

	ptr := unsafe.Pointer(C.malloc(C.size_t(count) * size))
	slice := (*[1 << 30]C.duckdb_logical_type)(ptr)[:count:count]

	for k, t := range types {
		slice[k] = t.c
	}

	typesPtr := (*C.duckdb_logical_type)(ptr)
	return DataChunk{C.duckdb_create_data_chunk(typesPtr, C.idx_t(count))}
}

func DestroyDataChunk(chunk *DataChunk) {
	C.duckdb_destroy_data_chunk(&chunk.c)
}

func DataChunkReset(chunk DataChunk) {
	C.duckdb_data_chunk_reset(chunk.c)
}

func DataChunkGetColumnCount(chunk DataChunk) uint64 {
	return uint64(C.duckdb_data_chunk_get_column_count(chunk.c))
}

func DataChunkGetVector(chunk DataChunk, colIdx uint64) Vector {
	return Vector{C.duckdb_data_chunk_get_vector(chunk.c, C.idx_t(colIdx))}
}

func DataChunkGetSize(chunk DataChunk) uint64 {
	return uint64(C.duckdb_data_chunk_get_size(chunk.c))
}

func DataChunkSetSize(chunk DataChunk, size uint64) {
	C.duckdb_data_chunk_set_size(chunk.c, C.idx_t(size))
}

// Vector

func VectorGetColumnType(vec Vector) LogicalType {
	return LogicalType{C.duckdb_vector_get_column_type(vec.c)}
}

func VectorGetData(vec Vector) unsafe.Pointer {
	return unsafe.Pointer(C.duckdb_vector_get_data(vec.c))
}

func VectorGetValidity(vec Vector) unsafe.Pointer {
	return unsafe.Pointer(C.duckdb_vector_get_validity(vec.c))
}

func VectorEnsureValidityWritable(vec Vector) {
	C.duckdb_vector_ensure_validity_writable(vec.c)
}

// String vector

// List vector

// Struct vector

// Array vector

// Validity

func ValidityRowIsValid(mask unsafe.Pointer, rowIdx uint64) bool {
	return bool(C.duckdb_validity_row_is_valid((*C.uint64_t)(mask), C.idx_t(rowIdx)))
}

func ValiditySetRowValidity(mask unsafe.Pointer, rowIdx uint64, valid bool) {
	C.duckdb_validity_set_row_validity((*C.uint64_t)(mask), C.idx_t(rowIdx), C.bool(valid))
}

func ValiditySetRowInvalid(mask unsafe.Pointer, rowIdx uint64) {
	C.duckdb_validity_set_row_invalid((*C.uint64_t)(mask), C.idx_t(rowIdx))
}

func ValiditySetRowValid(mask unsafe.Pointer, rowIdx uint64) {
	C.duckdb_validity_set_row_valid((*C.uint64_t)(mask), C.idx_t(rowIdx))
}

// Scalar function set

func CreateScalarFunctionSet(name string) ScalarFunctionSet {
	cName := C.CString(name)
	defer C.duckdb_free(unsafe.Pointer(cName))
	return ScalarFunctionSet{C.duckdb_create_scalar_function_set(cName)}
}

func DestroyScalarFunctionSet(set *ScalarFunctionSet) {
	C.duckdb_destroy_scalar_function_set(&set.c)
}

func AddScalarFunctionToSet(set ScalarFunctionSet, f ScalarFunction) State {
	return State(C.duckdb_add_scalar_function_to_set(set.c, f.c))
}

func RegisterScalarFunctionSet(conn Connection, set ScalarFunctionSet) State {
	return State(C.duckdb_register_scalar_function_set(conn.c, set.c))
}

// Scalar function

func CreateScalarFunction() ScalarFunction {
	return ScalarFunction{C.duckdb_create_scalar_function()}
}

func DestroyScalarFunction(f *ScalarFunction) {
	C.duckdb_destroy_scalar_function(&f.c)
}

func ScalarFunctionSetName(f ScalarFunction, name string) {
	cName := C.CString(name)
	defer C.duckdb_free(unsafe.Pointer(cName))
	C.duckdb_scalar_function_set_name(f.c, cName)
}

func ScalarFunctionAddParameter(f ScalarFunction, t LogicalType) {
	C.duckdb_scalar_function_add_parameter(f.c, t.c)
}

func ScalarFunctionSetReturnType(f ScalarFunction, t LogicalType) {
	C.duckdb_scalar_function_set_return_type(f.c, t.c)
}

func ScalarFunctionSetExtraInfo(f ScalarFunction, info unsafe.Pointer, callback DeleteCallbackT) {
	C.duckdb_scalar_function_set_extra_info(f.c, info, callback.c)
}

func ScalarFunctionSetFunction(f ScalarFunction, callback ScalarFunctionT) {
	C.duckdb_scalar_function_set_function(f.c, callback.c)
}

func RegisterScalarFunction(conn Connection, f ScalarFunction) State {
	return State(C.duckdb_register_scalar_function(conn.c, f.c))
}
