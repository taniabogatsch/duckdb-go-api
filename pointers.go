package duckdb

/*
#include <duckdb.h>
*/
import "C"
import "unsafe"

/*
Unsure about these:
- duckdb_extension_info
- duckdb_extension_access
*/

type (
	DeleteCallbackT unsafe.Pointer
	TaskState       unsafe.Pointer
)

type Column struct {
	Ptr unsafe.Pointer
}

func (col *Column) data() C.duckdb_column {
	var column C.duckdb_column
	column.internal_data = col.Ptr
	return column
}

type Vector struct {
	Ptr unsafe.Pointer
}

func (v *Vector) data() C.duckdb_vector {
	return C.duckdb_vector(v.Ptr)
}

type Result struct {
	Ptr unsafe.Pointer
}

func (r *Result) data() C.duckdb_result {
	var result C.duckdb_result
	result.internal_data = r.Ptr
	return result
}

type Database struct {
	Ptr unsafe.Pointer
}

func (db *Database) data() C.duckdb_database {
	return C.duckdb_database(db.Ptr)
}

type Connection struct {
	Ptr unsafe.Pointer
}

func (conn *Connection) data() C.duckdb_connection {
	return C.duckdb_connection(conn.Ptr)
}

type PreparedStatement struct {
	Ptr unsafe.Pointer
}

func (stmt *PreparedStatement) data() C.duckdb_prepared_statement {
	return (C.duckdb_prepared_statement)(stmt.Ptr)
}

type ExtractedStatements struct {
	Ptr unsafe.Pointer
}

func (stmts *ExtractedStatements) data() C.duckdb_extracted_statements {
	return (C.duckdb_extracted_statements)(stmts.Ptr)
}

type PendingResult struct {
	Ptr unsafe.Pointer
}

func (r *PendingResult) data() C.duckdb_pending_result {
	return (C.duckdb_pending_result)(r.Ptr)
}

type Appender struct {
	Ptr unsafe.Pointer
}

func (a *Appender) data() C.duckdb_appender {
	return (C.duckdb_appender)(a.Ptr)
}

type TableDescription struct {
	Ptr unsafe.Pointer
}

func (desc *TableDescription) data() C.duckdb_table_description {
	return (C.duckdb_table_description)(desc.Ptr)
}

type Config struct {
	Ptr unsafe.Pointer
}

func (config *Config) data() C.duckdb_config {
	return (C.duckdb_config)(config.Ptr)
}

type LogicalType struct {
	Ptr unsafe.Pointer
}

func (t *LogicalType) data() C.duckdb_logical_type {
	return C.duckdb_logical_type(t.Ptr)
}

type CreateTypeInfo struct {
	Ptr unsafe.Pointer
}

func (info *CreateTypeInfo) data() C.duckdb_create_type_info {
	return (C.duckdb_create_type_info)(info.Ptr)
}

type DataChunk struct {
	Ptr unsafe.Pointer
}

func (chunk *DataChunk) data() C.duckdb_data_chunk {
	return C.duckdb_data_chunk(chunk.Ptr)
}

type Value struct {
	Ptr unsafe.Pointer
}

func (v *Value) data() C.duckdb_value {
	return (C.duckdb_value)(v.Ptr)
}

type ProfilingInfo struct {
	Ptr unsafe.Pointer
}

func (info *ProfilingInfo) data() C.duckdb_profiling_info {
	return (C.duckdb_profiling_info)(info.Ptr)
}

type FunctionInfo struct {
	Ptr unsafe.Pointer
}

func (info *FunctionInfo) data() C.duckdb_function_info {
	return (C.duckdb_function_info)(info.Ptr)
}

type ScalarFunction struct {
	Ptr unsafe.Pointer
}

func (f *ScalarFunction) data() C.duckdb_scalar_function {
	return C.duckdb_scalar_function(f.Ptr)
}

type ScalarFunctionSet struct {
	Ptr unsafe.Pointer
}

func (set *ScalarFunctionSet) data() C.duckdb_scalar_function_set {
	return (C.duckdb_scalar_function_set)(set.Ptr)
}

type ScalarFunctionT unsafe.Pointer

type AggregateFunction struct {
	Ptr unsafe.Pointer
}

func (f *AggregateFunction) data() C.duckdb_aggregate_function {
	return (C.duckdb_aggregate_function)(f.Ptr)
}

type AggregateFunctionSet struct {
	Ptr unsafe.Pointer
}

func (set *AggregateFunctionSet) data() C.duckdb_aggregate_function_set {
	return (C.duckdb_aggregate_function_set)(set.Ptr)
}

type AggregateState struct {
	Ptr unsafe.Pointer
}

func (state *AggregateState) data() C.duckdb_aggregate_state {
	return (C.duckdb_aggregate_state)(state.Ptr)
}

type (
	AggregateStateSize unsafe.Pointer
	AggregateInitT     unsafe.Pointer
	AggregateDestroyT  unsafe.Pointer
	AggregateUpdateT   unsafe.Pointer
	AggregateCombineT  unsafe.Pointer
	AggregateFinalizeT unsafe.Pointer
)

type TableFunction struct {
	Ptr unsafe.Pointer
}

func (f *TableFunction) data() C.duckdb_table_function {
	return (C.duckdb_table_function)(f.Ptr)
}

type BindInfo struct {
	Ptr unsafe.Pointer
}

func (info *BindInfo) data() C.duckdb_bind_info {
	return (C.duckdb_bind_info)(info.Ptr)
}

type InitInfo struct {
	Ptr unsafe.Pointer
}

func (info *InitInfo) data() C.duckdb_init_info {
	return (C.duckdb_init_info)(info.Ptr)
}

type (
	TableFunctionBindT unsafe.Pointer
	TableFunctionInitT unsafe.Pointer
	TableFunctionT     unsafe.Pointer
)

type CastFunction struct {
	Ptr unsafe.Pointer
}

func (f *CastFunction) data() C.duckdb_cast_function {
	return (C.duckdb_cast_function)(f.Ptr)
}

type CastFunctionT unsafe.Pointer

type ReplacementScanInfo struct {
	Ptr unsafe.Pointer
}

func (info *ReplacementScanInfo) data() C.duckdb_replacement_scan_info {
	return (C.duckdb_replacement_scan_info)(info.Ptr)
}

type ReplacementScanCallbackT unsafe.Pointer

type Arrow struct {
	Ptr unsafe.Pointer
}

func (a *Arrow) data() C.duckdb_arrow {
	return (C.duckdb_arrow)(a.Ptr)
}

type ArrowStream struct {
	Ptr unsafe.Pointer
}

func (stream *ArrowStream) data() C.duckdb_arrow_stream {
	return (C.duckdb_arrow_stream)(stream.Ptr)
}

type ArrowSchema struct {
	Ptr unsafe.Pointer
}

func (schema *ArrowSchema) data() C.duckdb_arrow_schema {
	return (C.duckdb_arrow_schema)(schema.Ptr)
}

type ArrowArray struct {
	Ptr unsafe.Pointer
}

func (array *ArrowArray) data() C.duckdb_arrow_array {
	return (C.duckdb_arrow_array)(array.Ptr)
}
