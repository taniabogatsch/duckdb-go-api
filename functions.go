package duckdb

/*
#include <go_duckdb.h>
#include <duckdb_extension.h>
*/
import "C"
import "unsafe"

// TODO: should I always pass the custom types as pointers?

//func Open(path string, db *Database) State {
//	cPath := C.CString(path)
//	defer C.duckdb_free(unsafe.Pointer(cPath))
//
//	state := State(C.duckdb_open(cPath, db.data))
//	db.Ptr = unsafe.Pointer(db.data)
//	return state
//}
//
//func OpenExt(path string, db *Database, config Config, errMsg *string) State {
//	cPath := C.CString(path)
//	defer C.duckdb_free(unsafe.Pointer(cPath))
//
//	var err *C.char
//	defer C.duckdb_free(unsafe.Pointer(err))
//
//	state := State(C.duckdb_open_ext(cPath, db.data, *config.data, &err))
//	db.Ptr = unsafe.Pointer(db.data)
//	*errMsg = C.GoString(err)
//	return state
//}
//
//func Close(db *Database) {
//	C.duckdb_close(db.data)
//	db.Ptr = nil
//}

func Connect(db Database, conn *Connection) State {
	var c C.duckdb_connection
	state := State(C.duckdb_connect(db.data(), &c))
	conn.Ptr = unsafe.Pointer(c)
	return state
}

//func Interrupt(conn Connection) {
//	C.duckdb_interrupt(*conn.data)
//}
//
//func QueryProgress(conn Connection) QueryProgressType {
//	p := C.duckdb_query_progress(*conn.data)
//
//	return QueryProgressType{
//		Percentage:         float64(p.percentage),
//		RowsProcessed:      uint64(p.rows_processed),
//		TotalRowsToProcess: uint64(p.total_rows_to_process),
//	}
//}

func Disconnect(c *Connection) {
	data := c.data()
	C.duckdb_disconnect(&data)
	c.Ptr = nil
}

//func LibraryVersion() string {
//	cStr := C.duckdb_library_version()
//	defer C.duckdb_free(unsafe.Pointer(cStr))
//	return C.GoString(cStr)
//}

// Config

/*
static duckdb_state duckdb_create_config(duckdb_config *out_config) {
	return duckdb_ext_api->duckdb_create_config(out_config);
}

static size_t duckdb_config_count() {
	return duckdb_ext_api->duckdb_config_count();
}

static duckdb_state duckdb_get_config_flag(size_t index, const char **out_name, const char **out_description) {
	return duckdb_ext_api->duckdb_get_config_flag(index, out_name, out_description);
}

static duckdb_state duckdb_set_config(duckdb_config config, const char *name, const char *option) {
	return duckdb_ext_api->duckdb_set_config(config, name, option);
}

static void duckdb_destroy_config(duckdb_config *config) {
	return duckdb_ext_api->duckdb_destroy_config(config);
}
*/

// Result

/*
static duckdb_state duckdb_query(duckdb_connection connection, const char *query, duckdb_result *out_result) {
	return duckdb_ext_api->duckdb_query(connection, query, out_result);
}

static void duckdb_destroy_result(duckdb_result *result) {
	return duckdb_ext_api->duckdb_destroy_result(result);
}

static const char *duckdb_column_name(duckdb_result *result, idx_t col) {
	return duckdb_ext_api->duckdb_column_name(result, col);
}

static duckdb_type duckdb_column_type(duckdb_result *result, idx_t col) {
	return duckdb_ext_api->duckdb_column_type(result, col);
}

static duckdb_statement_type duckdb_result_statement_type(duckdb_result result) {
	return duckdb_ext_api->duckdb_result_statement_type(result);
}

static duckdb_logical_type duckdb_column_logical_type(duckdb_result *result, idx_t col) {
	return duckdb_ext_api->duckdb_column_logical_type(result, col);
}

static idx_t duckdb_column_count(duckdb_result *result) {
	return duckdb_ext_api->duckdb_column_count(result);
}

static idx_t duckdb_rows_changed(duckdb_result *result) {
	return duckdb_ext_api->duckdb_rows_changed(result);
}

static const char *duckdb_result_error(duckdb_result *result) {
	return duckdb_ext_api->duckdb_result_error(result);
}
*/

// Utility

/*
static void *duckdb_malloc(size_t size) {
	return duckdb_ext_api->duckdb_malloc(size);
}

static void duckdb_free(void *ptr) {
	return duckdb_ext_api->duckdb_free(ptr);
}

static idx_t duckdb_vector_size() {
	return duckdb_ext_api->duckdb_vector_size();
}
*/

// Type utility

/*
static bool duckdb_string_is_inlined(duckdb_string_t string) {
	return duckdb_ext_api->duckdb_string_is_inlined(string);
}

static duckdb_date_struct duckdb_from_date(duckdb_date date) {
	return duckdb_ext_api->duckdb_from_date(date);
}

static duckdb_date duckdb_to_date(duckdb_date_struct date) {
	return duckdb_ext_api->duckdb_to_date(date);
}

static bool duckdb_is_finite_date(duckdb_date date) {
	return duckdb_ext_api->duckdb_is_finite_date(date);
}

static duckdb_time_struct duckdb_from_time(duckdb_time time) {
	return duckdb_ext_api->duckdb_from_time(time);
}

static duckdb_time_tz duckdb_create_time_tz(int64_t micros, int32_t offset) {
	return duckdb_ext_api->duckdb_create_time_tz(micros, offset);
}

static duckdb_time_tz_struct duckdb_from_time_tz(duckdb_time_tz micros) {
	return duckdb_ext_api->duckdb_from_time_tz(micros);
}

static duckdb_time duckdb_to_time(duckdb_time_struct time) {
	return duckdb_ext_api->duckdb_to_time(time);
}

static duckdb_timestamp_struct duckdb_from_timestamp(duckdb_timestamp ts) {
	return duckdb_ext_api->duckdb_from_timestamp(ts);
}

static duckdb_timestamp duckdb_to_timestamp(duckdb_timestamp_struct ts) {
	return duckdb_ext_api->duckdb_to_timestamp(ts);
}

static bool duckdb_is_finite_timestamp(duckdb_timestamp ts) {
	return duckdb_ext_api->duckdb_is_finite_timestamp(ts);
}

static double duckdb_hugeint_to_double(duckdb_hugeint val) {
	return duckdb_ext_api->duckdb_hugeint_to_double(val);
}

static duckdb_hugeint duckdb_double_to_hugeint(double val) {
	return duckdb_ext_api->duckdb_double_to_hugeint(val);
}

static double duckdb_uhugeint_to_double(duckdb_uhugeint val) {
	return duckdb_ext_api->duckdb_uhugeint_to_double(val);
}

static duckdb_uhugeint duckdb_double_to_uhugeint(double val) {
	return duckdb_ext_api->duckdb_double_to_uhugeint(val);
}

static duckdb_decimal duckdb_double_to_decimal(double val, uint8_t width, uint8_t scale) {
	return duckdb_ext_api->duckdb_double_to_decimal(val, width, scale);
}

static double duckdb_decimal_to_double(duckdb_decimal val) {
	return duckdb_ext_api->duckdb_decimal_to_double(val);
}
*/

// Prepared and extracted statements

/*
static duckdb_state duckdb_prepare(duckdb_connection connection, const char *query,
                                   duckdb_prepared_statement *out_prepared_statement) {
	return duckdb_ext_api->duckdb_prepare(connection, query, out_prepared_statement);
}

static void duckdb_destroy_prepare(duckdb_prepared_statement *prepared_statement) {
	return duckdb_ext_api->duckdb_destroy_prepare(prepared_statement);
}

static const char *duckdb_prepare_error(duckdb_prepared_statement prepared_statement) {
	return duckdb_ext_api->duckdb_prepare_error(prepared_statement);
}

static idx_t duckdb_nparams(duckdb_prepared_statement prepared_statement) {
	return duckdb_ext_api->duckdb_nparams(prepared_statement);
}

static const char *duckdb_parameter_name(duckdb_prepared_statement prepared_statement, idx_t index) {
	return duckdb_ext_api->duckdb_parameter_name(prepared_statement, index);
}

static duckdb_type duckdb_param_type(duckdb_prepared_statement prepared_statement, idx_t param_idx) {
	return duckdb_ext_api->duckdb_param_type(prepared_statement, param_idx);
}

static duckdb_state duckdb_clear_bindings(duckdb_prepared_statement prepared_statement) {
	return duckdb_ext_api->duckdb_clear_bindings(prepared_statement);
}

static duckdb_statement_type duckdb_prepared_statement_type(duckdb_prepared_statement statement) {
	return duckdb_ext_api->duckdb_prepared_statement_type(statement);
}

static duckdb_state duckdb_bind_value(duckdb_prepared_statement prepared_statement, idx_t param_idx, duckdb_value val) {
	return duckdb_ext_api->duckdb_bind_value(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_parameter_index(duckdb_prepared_statement prepared_statement, idx_t *param_idx_out,
                                                const char *name) {
	return duckdb_ext_api->duckdb_bind_parameter_index(prepared_statement, param_idx_out, name);
}

static duckdb_state duckdb_bind_boolean(duckdb_prepared_statement prepared_statement, idx_t param_idx, bool val) {
	return duckdb_ext_api->duckdb_bind_boolean(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_int8(duckdb_prepared_statement prepared_statement, idx_t param_idx, int8_t val) {
	return duckdb_ext_api->duckdb_bind_int8(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_int16(duckdb_prepared_statement prepared_statement, idx_t param_idx, int16_t val) {
	return duckdb_ext_api->duckdb_bind_int16(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_int32(duckdb_prepared_statement prepared_statement, idx_t param_idx, int32_t val) {
	return duckdb_ext_api->duckdb_bind_int32(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_int64(duckdb_prepared_statement prepared_statement, idx_t param_idx, int64_t val) {
	return duckdb_ext_api->duckdb_bind_int64(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_hugeint(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                        duckdb_hugeint val) {
	return duckdb_ext_api->duckdb_bind_hugeint(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_uhugeint(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                         duckdb_uhugeint val) {
	return duckdb_ext_api->duckdb_bind_uhugeint(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_decimal(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                        duckdb_decimal val) {
	return duckdb_ext_api->duckdb_bind_decimal(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_uint8(duckdb_prepared_statement prepared_statement, idx_t param_idx, uint8_t val) {
	return duckdb_ext_api->duckdb_bind_uint8(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_uint16(duckdb_prepared_statement prepared_statement, idx_t param_idx, uint16_t val) {
	return duckdb_ext_api->duckdb_bind_uint16(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_uint32(duckdb_prepared_statement prepared_statement, idx_t param_idx, uint32_t val) {
	return duckdb_ext_api->duckdb_bind_uint32(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_uint64(duckdb_prepared_statement prepared_statement, idx_t param_idx, uint64_t val) {
	return duckdb_ext_api->duckdb_bind_uint64(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_float(duckdb_prepared_statement prepared_statement, idx_t param_idx, float val) {
	return duckdb_ext_api->duckdb_bind_float(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_double(duckdb_prepared_statement prepared_statement, idx_t param_idx, double val) {
	return duckdb_ext_api->duckdb_bind_double(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_date(duckdb_prepared_statement prepared_statement, idx_t param_idx, duckdb_date val) {
	return duckdb_ext_api->duckdb_bind_date(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_time(duckdb_prepared_statement prepared_statement, idx_t param_idx, duckdb_time val) {
	return duckdb_ext_api->duckdb_bind_time(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_timestamp(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                          duckdb_timestamp val) {
	return duckdb_ext_api->duckdb_bind_timestamp(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_timestamp_tz(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                             duckdb_timestamp val) {
	return duckdb_ext_api->duckdb_bind_timestamp_tz(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_interval(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                         duckdb_interval val) {
	return duckdb_ext_api->duckdb_bind_interval(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_varchar(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                        const char *val) {
	return duckdb_ext_api->duckdb_bind_varchar(prepared_statement, param_idx, val);
}

static duckdb_state duckdb_bind_varchar_length(duckdb_prepared_statement prepared_statement, idx_t param_idx,
                                               const char *val, idx_t length) {
	return duckdb_ext_api->duckdb_bind_varchar_length(prepared_statement, param_idx, val, length);
}

static duckdb_state duckdb_bind_blob(duckdb_prepared_statement prepared_statement, idx_t param_idx, const void *data,
                                     idx_t length) {
	return duckdb_ext_api->duckdb_bind_blob(prepared_statement, param_idx, data, length);
}

static duckdb_state duckdb_bind_null(duckdb_prepared_statement prepared_statement, idx_t param_idx) {
	return duckdb_ext_api->duckdb_bind_null(prepared_statement, param_idx);
}

static duckdb_state duckdb_execute_prepared(duckdb_prepared_statement prepared_statement, duckdb_result *out_result) {
	return duckdb_ext_api->duckdb_execute_prepared(prepared_statement, out_result);
}

static idx_t duckdb_extract_statements(duckdb_connection connection, const char *query,
                                       duckdb_extracted_statements *out_extracted_statements) {
	return duckdb_ext_api->duckdb_extract_statements(connection, query, out_extracted_statements);
}

static duckdb_state duckdb_prepare_extracted_statement(duckdb_connection connection,
                                                       duckdb_extracted_statements extracted_statements, idx_t index,
                                                       duckdb_prepared_statement *out_prepared_statement) {
	return duckdb_ext_api->duckdb_prepare_extracted_statement(connection, extracted_statements, index,
	                                                          out_prepared_statement);
}

static const char *duckdb_extract_statements_error(duckdb_extracted_statements extracted_statements) {
	return duckdb_ext_api->duckdb_extract_statements_error(extracted_statements);
}

static void duckdb_destroy_extracted(duckdb_extracted_statements *extracted_statements) {
	return duckdb_ext_api->duckdb_destroy_extracted(extracted_statements);
}

static duckdb_state duckdb_pending_prepared(duckdb_prepared_statement prepared_statement,
                                            duckdb_pending_result *out_result) {
	return duckdb_ext_api->duckdb_pending_prepared(prepared_statement, out_result);
}
*/

// Pending result

/*
static void duckdb_destroy_pending(duckdb_pending_result *pending_result) {
	return duckdb_ext_api->duckdb_destroy_pending(pending_result);
}

static const char *duckdb_pending_error(duckdb_pending_result pending_result) {
	return duckdb_ext_api->duckdb_pending_error(pending_result);
}

static duckdb_pending_state duckdb_pending_execute_task(duckdb_pending_result pending_result) {
	return duckdb_ext_api->duckdb_pending_execute_task(pending_result);
}

static duckdb_pending_state duckdb_pending_execute_check_state(duckdb_pending_result pending_result) {
	return duckdb_ext_api->duckdb_pending_execute_check_state(pending_result);
}

static duckdb_state duckdb_execute_pending(duckdb_pending_result pending_result, duckdb_result *out_result) {
	return duckdb_ext_api->duckdb_execute_pending(pending_result, out_result);
}

static bool duckdb_pending_execution_is_finished(duckdb_pending_state pending_state) {
	return duckdb_ext_api->duckdb_pending_execution_is_finished(pending_state);
}
*/

// Types

/*
static void duckdb_destroy_value(duckdb_value *value) {
	return duckdb_ext_api->duckdb_destroy_value(value);
}

static duckdb_value duckdb_create_varchar(const char *text) {
	return duckdb_ext_api->duckdb_create_varchar(text);
}

static duckdb_value duckdb_create_varchar_length(const char *text, idx_t length) {
	return duckdb_ext_api->duckdb_create_varchar_length(text, length);
}

static duckdb_value duckdb_create_int64(int64_t val) {
	return duckdb_ext_api->duckdb_create_int64(val);
}

static duckdb_value duckdb_create_struct_value(duckdb_logical_type type, duckdb_value *values) {
	return duckdb_ext_api->duckdb_create_struct_value(type, values);
}

static duckdb_value duckdb_create_list_value(duckdb_logical_type type, duckdb_value *values, idx_t value_count) {
	return duckdb_ext_api->duckdb_create_list_value(type, values, value_count);
}

static duckdb_value duckdb_create_array_value(duckdb_logical_type type, duckdb_value *values, idx_t value_count) {
	return duckdb_ext_api->duckdb_create_array_value(type, values, value_count);
}

static char *duckdb_get_varchar(duckdb_value value) {
	return duckdb_ext_api->duckdb_get_varchar(value);
}

static int64_t duckdb_get_int64(duckdb_value val) {
	return duckdb_ext_api->duckdb_get_int64(val);
}
*/

func CreateLogicalType(t Type) LogicalType {
	internalType := C.duckdb_type(t)
	logicalType := C.duckdb_create_logical_type(internalType)
	return LogicalType{unsafe.Pointer(logicalType)}
}

/*
static char *duckdb_logical_type_get_alias(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_logical_type_get_alias(type);
}

static duckdb_logical_type duckdb_create_list_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_create_list_type(type);
}

static duckdb_logical_type duckdb_create_array_type(duckdb_logical_type type, idx_t array_size) {
	return duckdb_ext_api->duckdb_create_array_type(type, array_size);
}

static duckdb_logical_type duckdb_create_map_type(duckdb_logical_type key_type, duckdb_logical_type value_type) {
	return duckdb_ext_api->duckdb_create_map_type(key_type, value_type);
}

static duckdb_logical_type duckdb_create_union_type(duckdb_logical_type *member_types, const char **member_names,
                                                    idx_t member_count) {
	return duckdb_ext_api->duckdb_create_union_type(member_types, member_names, member_count);
}

static duckdb_logical_type duckdb_create_struct_type(duckdb_logical_type *member_types, const char **member_names,
                                                     idx_t member_count) {
	return duckdb_ext_api->duckdb_create_struct_type(member_types, member_names, member_count);
}

static duckdb_logical_type duckdb_create_enum_type(const char **member_names, idx_t member_count) {
	return duckdb_ext_api->duckdb_create_enum_type(member_names, member_count);
}

static duckdb_logical_type duckdb_create_decimal_type(uint8_t width, uint8_t scale) {
	return duckdb_ext_api->duckdb_create_decimal_type(width, scale);
}

static duckdb_type duckdb_get_type_id(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_get_type_id(type);
}

static uint8_t duckdb_decimal_width(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_decimal_width(type);
}

static uint8_t duckdb_decimal_scale(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_decimal_scale(type);
}

static duckdb_type duckdb_decimal_internal_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_decimal_internal_type(type);
}

static duckdb_type duckdb_enum_internal_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_enum_internal_type(type);
}

static uint32_t duckdb_enum_dictionary_size(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_enum_dictionary_size(type);
}

static char *duckdb_enum_dictionary_value(duckdb_logical_type type, idx_t index) {
	return duckdb_ext_api->duckdb_enum_dictionary_value(type, index);
}

static duckdb_logical_type duckdb_list_type_child_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_list_type_child_type(type);
}

static duckdb_logical_type duckdb_array_type_child_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_array_type_child_type(type);
}

static idx_t duckdb_array_type_array_size(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_array_type_array_size(type);
}

static duckdb_logical_type duckdb_map_type_key_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_map_type_key_type(type);
}

static duckdb_logical_type duckdb_map_type_value_type(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_map_type_value_type(type);
}

static idx_t duckdb_struct_type_child_count(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_struct_type_child_count(type);
}

static char *duckdb_struct_type_child_name(duckdb_logical_type type, idx_t index) {
	return duckdb_ext_api->duckdb_struct_type_child_name(type, index);
}

static duckdb_logical_type duckdb_struct_type_child_type(duckdb_logical_type type, idx_t index) {
	return duckdb_ext_api->duckdb_struct_type_child_type(type, index);
}

static idx_t duckdb_union_type_member_count(duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_union_type_member_count(type);
}

static char *duckdb_union_type_member_name(duckdb_logical_type type, idx_t index) {
	return duckdb_ext_api->duckdb_union_type_member_name(type, index);
}

static duckdb_logical_type duckdb_union_type_member_type(duckdb_logical_type type, idx_t index) {
	return duckdb_ext_api->duckdb_union_type_member_type(type, index);
}
*/

func DestroyLogicalType(t *LogicalType) {
	data := t.data()
	C.duckdb_destroy_logical_type(&data)
	t.Ptr = nil
}

// Data chunk

//func FetchChunk(result Result) DataChunk {
//	chunk := C.duckdb_fetch_chunk(result.data)
//	return DataChunk{}
//}
//
//func CreateDataChunk(types []LogicalType, count uint64) DataChunk {
//	var dummy C.duckdb_logical_type
//	size := C.size_t(unsafe.Sizeof(dummy))
//
//	ptr := unsafe.Pointer(C.malloc(C.size_t(count) * size))
//	slice := (*[1 << 30]C.duckdb_logical_type)(ptr)[:count:count]
//
//	for k, t := range types {
//		slice[k] = t.c
//	}
//
//	typesPtr := (*C.duckdb_logical_type)(ptr)
//	return DataChunk{C.duckdb_create_data_chunk(typesPtr, C.idx_t(count))}
//}
//
//func DestroyDataChunk(chunk *DataChunk) {
//	C.duckdb_destroy_data_chunk(&chunk.c)
//}
//
//func DataChunkReset(chunk DataChunk) {
//	C.duckdb_data_chunk_reset(chunk.c)
//}
//
//func DataChunkGetColumnCount(chunk DataChunk) uint64 {
//	return uint64(C.duckdb_data_chunk_get_column_count(chunk.c))
//}

func DataChunkGetVector(chunk DataChunk, colIdx uint64) Vector {
	v := C.duckdb_data_chunk_get_vector(chunk.data(), C.idx_t(colIdx))
	return Vector{unsafe.Pointer(v)}
}

func DataChunkGetSize(chunk DataChunk) uint64 {
	size := C.duckdb_data_chunk_get_size(chunk.data())
	return uint64(size)
}

//func DataChunkSetSize(chunk DataChunk, size uint64) {
//	C.duckdb_data_chunk_set_size(chunk.c, C.idx_t(size))
//}

// Vector

//func VectorGetColumnType(v Vector) LogicalType {
//	ptr := C.duckdb_vector(v)
//	return LogicalType{C.duckdb_vector_get_column_type(ptr)}
//}

func VectorGetData(v Vector) unsafe.Pointer {
	data := C.duckdb_vector_get_data(v.data())
	return unsafe.Pointer(data)
}

func VectorGetValidity(v Vector) unsafe.Pointer {
	mask := C.duckdb_vector_get_validity(v.data())
	return unsafe.Pointer(mask)
}

func VectorEnsureValidityWritable(v Vector) {
	C.duckdb_vector_ensure_validity_writable(v.data())
}

// String vector

/*
static void duckdb_vector_assign_string_element(duckdb_vector vector, idx_t index, const char *str) {
	return duckdb_ext_api->duckdb_vector_assign_string_element(vector, index, str);
}

static void duckdb_vector_assign_string_element_len(duckdb_vector vector, idx_t index, const char *str, idx_t str_len) {
	return duckdb_ext_api->duckdb_vector_assign_string_element_len(vector, index, str, str_len);
}
*/

// List vector

/*
static duckdb_vector duckdb_list_vector_get_child(duckdb_vector vector) {
	return duckdb_ext_api->duckdb_list_vector_get_child(vector);
}

static idx_t duckdb_list_vector_get_size(duckdb_vector vector) {
	return duckdb_ext_api->duckdb_list_vector_get_size(vector);
}

static duckdb_state duckdb_list_vector_set_size(duckdb_vector vector, idx_t size) {
	return duckdb_ext_api->duckdb_list_vector_set_size(vector, size);
}

static duckdb_state duckdb_list_vector_reserve(duckdb_vector vector, idx_t required_capacity) {
	return duckdb_ext_api->duckdb_list_vector_reserve(vector, required_capacity);
}
*/

// Struct vector

/*
static duckdb_vector duckdb_struct_vector_get_child(duckdb_vector vector, idx_t index) {
	return duckdb_ext_api->duckdb_struct_vector_get_child(vector, index);
}
*/

// Array vector

/*
static duckdb_vector duckdb_array_vector_get_child(duckdb_vector vector) {
	return duckdb_ext_api->duckdb_array_vector_get_child(vector);
}
*/

// Validity

func ValidityRowIsValid(mask unsafe.Pointer, rowIdx uint64) bool {
	return bool(C.duckdb_validity_row_is_valid((*C.uint64_t)(mask), C.idx_t(rowIdx)))
}

//func ValiditySetRowValidity(mask unsafe.Pointer, rowIdx uint64, valid bool) {
//	C.duckdb_validity_set_row_validity((*C.uint64_t)(mask), C.idx_t(rowIdx), C.bool(valid))
//}

func ValiditySetRowInvalid(mask unsafe.Pointer, rowIdx uint64) {
	C.duckdb_validity_set_row_invalid((*C.uint64_t)(mask), C.idx_t(rowIdx))
}

//func ValiditySetRowValid(mask unsafe.Pointer, rowIdx uint64) {
//	C.duckdb_validity_set_row_valid((*C.uint64_t)(mask), C.idx_t(rowIdx))
//}

// Scalar function

func CreateScalarFunction() ScalarFunction {
	data := C.duckdb_create_scalar_function()
	return ScalarFunction{unsafe.Pointer(data)}
}

func DestroyScalarFunction(f *ScalarFunction) {
	data := f.data()
	C.duckdb_destroy_scalar_function(&data)
	f.Ptr = nil
}

func ScalarFunctionSetName(f ScalarFunction, name string) {
	cName := C.CString(name)
	defer C.duckdb_free(unsafe.Pointer(cName))
	C.duckdb_scalar_function_set_name(f.data(), cName)
}

func ScalarFunctionAddParameter(f ScalarFunction, t LogicalType) {
	C.duckdb_scalar_function_add_parameter(f.data(), t.data())
}

func ScalarFunctionSetReturnType(f ScalarFunction, t LogicalType) {
	C.duckdb_scalar_function_set_return_type(f.data(), t.data())
}

//func ScalarFunctionSetExtraInfo(f ScalarFunction, info unsafe.Pointer, deleteFun unsafe.Pointer) {
//	df := C.duckdb_delete_callback_t(deleteFun)
//	C.duckdb_scalar_function_set_extra_info(f.c, info, df)
//}

func ScalarFunctionSetFunction(f ScalarFunction, scalarFun unsafe.Pointer) {
	sf := C.duckdb_scalar_function_t(scalarFun)
	C.duckdb_scalar_function_set_function(f.data(), sf)
}

func RegisterScalarFunction(c Connection, f ScalarFunction) State {
	return State(C.duckdb_register_scalar_function(c.data(), f.data()))
}

// Table function

/*
static duckdb_table_function duckdb_create_table_function() {
	return duckdb_ext_api->duckdb_create_table_function();
}

static void duckdb_destroy_table_function(duckdb_table_function *table_function) {
	return duckdb_ext_api->duckdb_destroy_table_function(table_function);
}

static void duckdb_table_function_set_name(duckdb_table_function table_function, const char *name) {
	return duckdb_ext_api->duckdb_table_function_set_name(table_function, name);
}

static void duckdb_table_function_add_parameter(duckdb_table_function table_function, duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_table_function_add_parameter(table_function, type);
}

static void duckdb_table_function_add_named_parameter(duckdb_table_function table_function, const char *name,
                                                      duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_table_function_add_named_parameter(table_function, name, type);
}

static void duckdb_table_function_set_extra_info(duckdb_table_function table_function, void *extra_info,
                                                 duckdb_delete_callback_t destroy) {
	return duckdb_ext_api->duckdb_table_function_set_extra_info(table_function, extra_info, destroy);
}

static void duckdb_table_function_set_bind(duckdb_table_function table_function, duckdb_table_function_bind_t bind) {
	return duckdb_ext_api->duckdb_table_function_set_bind(table_function, bind);
}

static void duckdb_table_function_set_init(duckdb_table_function table_function, duckdb_table_function_init_t init) {
	return duckdb_ext_api->duckdb_table_function_set_init(table_function, init);
}

static void duckdb_table_function_set_local_init(duckdb_table_function table_function,
                                                 duckdb_table_function_init_t init) {
	return duckdb_ext_api->duckdb_table_function_set_local_init(table_function, init);
}

static void duckdb_table_function_set_function(duckdb_table_function table_function, duckdb_table_function_t function) {
	return duckdb_ext_api->duckdb_table_function_set_function(table_function, function);
}

static void duckdb_table_function_supports_projection_pushdown(duckdb_table_function table_function, bool pushdown) {
	return duckdb_ext_api->duckdb_table_function_supports_projection_pushdown(table_function, pushdown);
}

static duckdb_state duckdb_register_table_function(duckdb_connection con, duckdb_table_function function) {
	return duckdb_ext_api->duckdb_register_table_function(con, function);
}
*/

// Bind

/*
static void *duckdb_bind_get_extra_info(duckdb_bind_info info) {
	return duckdb_ext_api->duckdb_bind_get_extra_info(info);
}

static void duckdb_bind_add_result_column(duckdb_bind_info info, const char *name, duckdb_logical_type type) {
	return duckdb_ext_api->duckdb_bind_add_result_column(info, name, type);
}

static idx_t duckdb_bind_get_parameter_count(duckdb_bind_info info) {
	return duckdb_ext_api->duckdb_bind_get_parameter_count(info);
}

static duckdb_value duckdb_bind_get_parameter(duckdb_bind_info info, idx_t index) {
	return duckdb_ext_api->duckdb_bind_get_parameter(info, index);
}

static duckdb_value duckdb_bind_get_named_parameter(duckdb_bind_info info, const char *name) {
	return duckdb_ext_api->duckdb_bind_get_named_parameter(info, name);
}

static void duckdb_bind_set_bind_data(duckdb_bind_info info, void *bind_data, duckdb_delete_callback_t destroy) {
	return duckdb_ext_api->duckdb_bind_set_bind_data(info, bind_data, destroy);
}

static void duckdb_bind_set_cardinality(duckdb_bind_info info, idx_t cardinality, bool is_exact) {
	return duckdb_ext_api->duckdb_bind_set_cardinality(info, cardinality, is_exact);
}

static void duckdb_bind_set_error(duckdb_bind_info info, const char *error) {
	return duckdb_ext_api->duckdb_bind_set_error(info, error);
}
*/

// Init

/*
static void *duckdb_init_get_extra_info(duckdb_init_info info) {
	return duckdb_ext_api->duckdb_init_get_extra_info(info);
}

static void *duckdb_init_get_bind_data(duckdb_init_info info) {
	return duckdb_ext_api->duckdb_init_get_bind_data(info);
}

static void duckdb_init_set_init_data(duckdb_init_info info, void *init_data, duckdb_delete_callback_t destroy) {
	return duckdb_ext_api->duckdb_init_set_init_data(info, init_data, destroy);
}

static idx_t duckdb_init_get_column_count(duckdb_init_info info) {
	return duckdb_ext_api->duckdb_init_get_column_count(info);
}

static idx_t duckdb_init_get_column_index(duckdb_init_info info, idx_t column_index) {
	return duckdb_ext_api->duckdb_init_get_column_index(info, column_index);
}

static void duckdb_init_set_max_threads(duckdb_init_info info, idx_t max_threads) {
	return duckdb_ext_api->duckdb_init_set_max_threads(info, max_threads);
}

static void duckdb_init_set_error(duckdb_init_info info, const char *error) {
	return duckdb_ext_api->duckdb_init_set_error(info, error);
}
*/

// Function

/*
static void *duckdb_function_get_extra_info(duckdb_function_info info) {
	return duckdb_ext_api->duckdb_function_get_extra_info(info);
}

static void *duckdb_function_get_bind_data(duckdb_function_info info) {
	return duckdb_ext_api->duckdb_function_get_bind_data(info);
}

static void *duckdb_function_get_init_data(duckdb_function_info info) {
	return duckdb_ext_api->duckdb_function_get_init_data(info);
}

static void *duckdb_function_get_local_init_data(duckdb_function_info info) {
	return duckdb_ext_api->duckdb_function_get_local_init_data(info);
}

static void duckdb_function_set_error(duckdb_function_info info, const char *error) {
	return duckdb_ext_api->duckdb_function_set_error(info, error);
}
*/

// Replacement scan

/*
static void duckdb_add_replacement_scan(duckdb_database db, duckdb_replacement_callback_t replacement, void *extra_data,
                                        duckdb_delete_callback_t delete_callback) {
	return duckdb_ext_api->duckdb_add_replacement_scan(db, replacement, extra_data, delete_callback);
}

static void duckdb_replacement_scan_set_function_name(duckdb_replacement_scan_info info, const char *function_name) {
	return duckdb_ext_api->duckdb_replacement_scan_set_function_name(info, function_name);
}

static void duckdb_replacement_scan_add_parameter(duckdb_replacement_scan_info info, duckdb_value parameter) {
	return duckdb_ext_api->duckdb_replacement_scan_add_parameter(info, parameter);
}

static void duckdb_replacement_scan_set_error(duckdb_replacement_scan_info info, const char *error) {
	return duckdb_ext_api->duckdb_replacement_scan_set_error(info, error);
}
*/

// Appender

/*
static duckdb_state duckdb_appender_create(duckdb_connection connection, const char *schema, const char *table,
                                           duckdb_appender *out_appender) {
	return duckdb_ext_api->duckdb_appender_create(connection, schema, table, out_appender);
}

static idx_t duckdb_appender_column_count(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_column_count(appender);
}

static duckdb_logical_type duckdb_appender_column_type(duckdb_appender appender, idx_t col_idx) {
	return duckdb_ext_api->duckdb_appender_column_type(appender, col_idx);
}

static const char *duckdb_appender_error(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_error(appender);
}

static duckdb_state duckdb_appender_flush(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_flush(appender);
}

static duckdb_state duckdb_appender_close(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_close(appender);
}

static duckdb_state duckdb_appender_destroy(duckdb_appender *appender) {
	return duckdb_ext_api->duckdb_appender_destroy(appender);
}

static duckdb_state duckdb_appender_begin_row(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_begin_row(appender);
}

static duckdb_state duckdb_appender_end_row(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_appender_end_row(appender);
}

static duckdb_state duckdb_append_default(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_append_default(appender);
}

static duckdb_state duckdb_append_bool(duckdb_appender appender, bool value) {
	return duckdb_ext_api->duckdb_append_bool(appender, value);
}

static duckdb_state duckdb_append_int8(duckdb_appender appender, int8_t value) {
	return duckdb_ext_api->duckdb_append_int8(appender, value);
}

static duckdb_state duckdb_append_int16(duckdb_appender appender, int16_t value) {
	return duckdb_ext_api->duckdb_append_int16(appender, value);
}

static duckdb_state duckdb_append_int32(duckdb_appender appender, int32_t value) {
	return duckdb_ext_api->duckdb_append_int32(appender, value);
}

static duckdb_state duckdb_append_int64(duckdb_appender appender, int64_t value) {
	return duckdb_ext_api->duckdb_append_int64(appender, value);
}

static duckdb_state duckdb_append_hugeint(duckdb_appender appender, duckdb_hugeint value) {
	return duckdb_ext_api->duckdb_append_hugeint(appender, value);
}

static duckdb_state duckdb_append_uint8(duckdb_appender appender, uint8_t value) {
	return duckdb_ext_api->duckdb_append_uint8(appender, value);
}

static duckdb_state duckdb_append_uint16(duckdb_appender appender, uint16_t value) {
	return duckdb_ext_api->duckdb_append_uint16(appender, value);
}

static duckdb_state duckdb_append_uint32(duckdb_appender appender, uint32_t value) {
	return duckdb_ext_api->duckdb_append_uint32(appender, value);
}

static duckdb_state duckdb_append_uint64(duckdb_appender appender, uint64_t value) {
	return duckdb_ext_api->duckdb_append_uint64(appender, value);
}

static duckdb_state duckdb_append_uhugeint(duckdb_appender appender, duckdb_uhugeint value) {
	return duckdb_ext_api->duckdb_append_uhugeint(appender, value);
}

static duckdb_state duckdb_append_float(duckdb_appender appender, float value) {
	return duckdb_ext_api->duckdb_append_float(appender, value);
}

static duckdb_state duckdb_append_double(duckdb_appender appender, double value) {
	return duckdb_ext_api->duckdb_append_double(appender, value);
}

static duckdb_state duckdb_append_date(duckdb_appender appender, duckdb_date value) {
	return duckdb_ext_api->duckdb_append_date(appender, value);
}

static duckdb_state duckdb_append_time(duckdb_appender appender, duckdb_time value) {
	return duckdb_ext_api->duckdb_append_time(appender, value);
}

static duckdb_state duckdb_append_timestamp(duckdb_appender appender, duckdb_timestamp value) {
	return duckdb_ext_api->duckdb_append_timestamp(appender, value);
}

static duckdb_state duckdb_append_interval(duckdb_appender appender, duckdb_interval value) {
	return duckdb_ext_api->duckdb_append_interval(appender, value);
}

static duckdb_state duckdb_append_varchar(duckdb_appender appender, const char *val) {
	return duckdb_ext_api->duckdb_append_varchar(appender, val);
}

static duckdb_state duckdb_append_varchar_length(duckdb_appender appender, const char *val, idx_t length) {
	return duckdb_ext_api->duckdb_append_varchar_length(appender, val, length);
}

static duckdb_state duckdb_append_blob(duckdb_appender appender, const void *data, idx_t length) {
	return duckdb_ext_api->duckdb_append_blob(appender, data, length);
}

static duckdb_state duckdb_append_null(duckdb_appender appender) {
	return duckdb_ext_api->duckdb_append_null(appender);
}

static duckdb_state duckdb_append_data_chunk(duckdb_appender appender, duckdb_data_chunk chunk) {
	return duckdb_ext_api->duckdb_append_data_chunk(appender, chunk);
}
*/

// Task state

/*
static void duckdb_execute_tasks(duckdb_database database, idx_t max_tasks) {
	return duckdb_ext_api->duckdb_execute_tasks(database, max_tasks);
}

static duckdb_task_state duckdb_create_task_state(duckdb_database database) {
	return duckdb_ext_api->duckdb_create_task_state(database);
}

static void duckdb_execute_tasks_state(duckdb_task_state state) {
	return duckdb_ext_api->duckdb_execute_tasks_state(state);
}

static idx_t duckdb_execute_n_tasks_state(duckdb_task_state state, idx_t max_tasks) {
	return duckdb_ext_api->duckdb_execute_n_tasks_state(state, max_tasks);
}

static void duckdb_finish_execution(duckdb_task_state state) {
	return duckdb_ext_api->duckdb_finish_execution(state);
}

static bool duckdb_task_state_is_finished(duckdb_task_state state) {
	return duckdb_ext_api->duckdb_task_state_is_finished(state);
}

static void duckdb_destroy_task_state(duckdb_task_state state) {
	return duckdb_ext_api->duckdb_destroy_task_state(state);
}

static bool duckdb_execution_is_finished(duckdb_connection con) {
	return duckdb_ext_api->duckdb_execution_is_finished(con);
}
*/

// Profiling

/*
static duckdb_profiling_info duckdb_get_profiling_info(duckdb_connection connection) {
	return duckdb_ext_api->duckdb_get_profiling_info(connection);
}

static duckdb_value duckdb_profiling_info_get_value(duckdb_profiling_info info, const char *key) {
	return duckdb_ext_api->duckdb_profiling_info_get_value(info, key);
}

static idx_t duckdb_profiling_info_get_child_count(duckdb_profiling_info info) {
	return duckdb_ext_api->duckdb_profiling_info_get_child_count(info);
}

static duckdb_profiling_info duckdb_profiling_info_get_child(duckdb_profiling_info info, idx_t index) {
	return duckdb_ext_api->duckdb_profiling_info_get_child(info, index);
}

static duckdb_value duckdb_profiling_info_get_metrics(duckdb_profiling_info info) {
	return duckdb_ext_api->duckdb_profiling_info_get_metrics(info);
}
*/

// ...

// Scalar function set

//func CreateScalarFunctionSet(name string) ScalarFunctionSet {
//	cName := C.CString(name)
//	defer C.duckdb_free(unsafe.Pointer(cName))
//
//	return ScalarFunctionSet{C.duckdb_create_scalar_function_set(cName)}
//}
//
//func DestroyScalarFunctionSet(set *ScalarFunctionSet) {
//	C.duckdb_destroy_scalar_function_set(&set.c)
//}
//
//func AddScalarFunctionToSet(set ScalarFunctionSet, f ScalarFunction) State {
//	return State(C.duckdb_add_scalar_function_to_set(set.c, f.c))
//}
//
//func RegisterScalarFunctionSet(conn Connection, set ScalarFunctionSet) State {
//	return State(C.duckdb_register_scalar_function_set(conn.c, set.c))
//}
