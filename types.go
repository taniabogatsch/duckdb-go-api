package duckdb

/*
#include <go_duckdb.h>
*/
import "C"
import "unsafe"

/*
TODO
//! DuckDB's index type.
typedef uint64_t idx_t;

//! Used for threading, contains a task state. Must be destroyed with `duckdb_destroy_state`.
typedef void *duckdb_task_state;

typedef struct _duckdb_extension_info {
	void *internal_ptr;
} * duckdb_extension_info;

//! Returns the aggregate state size
typedef idx_t (*duckdb_aggregate_state_size)(duckdb_function_info info);
//! Initialize the aggregate state
typedef void (*duckdb_aggregate_init_t)(duckdb_function_info info, duckdb_aggregate_state state);
//! Destroy aggregate state (optional)
typedef void (*duckdb_aggregate_destroy_t)(duckdb_aggregate_state *states, idx_t count);
//! Update a set of aggregate states with new values
typedef void (*duckdb_aggregate_update_t)(duckdb_function_info info, duckdb_data_chunk input,
                                          duckdb_aggregate_state *states);
//! Combine aggregate states
typedef void (*duckdb_aggregate_combine_t)(duckdb_function_info info, duckdb_aggregate_state *source,
                                           duckdb_aggregate_state *target, idx_t count);
//! Finalize aggregate states into a result vector
typedef void (*duckdb_aggregate_finalize_t)(duckdb_function_info info, duckdb_aggregate_state *source,
                                            duckdb_vector result, idx_t count, idx_t offset);

//! The bind function of the table function.
typedef void (*duckdb_table_function_bind_t)(duckdb_bind_info info);

//! The (possibly thread-local) init function of the table function.
typedef void (*duckdb_table_function_init_t)(duckdb_init_info info);

//! The main function of the table function.
typedef void (*duckdb_table_function_t)(duckdb_function_info info, duckdb_data_chunk output);

typedef bool (*duckdb_cast_function_t)(duckdb_function_info info, idx_t count, duckdb_vector input, duckdb_vector output);


//! A replacement scan function that can be added to a database.
typedef void (*duckdb_replacement_callback_t)(duckdb_replacement_scan_info info, const char *table_name, void *data);


struct duckdb_extension_access {
	//! Indicate that an error has occured
	void (*set_error)(duckdb_extension_info info, const char *error);
	//! Fetch the database from duckdb to register extensions to
	duckdb_database *(*get_database)(duckdb_extension_info info);
	//! Fetch the API
	void *(*get_api)(duckdb_extension_info info, const char *version);
};
*/

type TaskState unsafe.Pointer

type Date struct {
	c C.duckdb_date
}

type DateStruct struct {
	c C.duckdb_date_struct
}

type Time struct {
	c C.duckdb_time
}

type TimeStruct struct {
	c C.duckdb_time_struct
}

type TimeTZ struct {
	c C.duckdb_time_tz
}

type TimeTZStruct struct {
	c C.duckdb_time_tz_struct
}

type Timestamp struct {
	c C.duckdb_timestamp
}

type TimestampStruct struct {
	c C.duckdb_timestamp_struct
}

type Interval struct {
	c C.duckdb_interval
}

type Hugeint struct {
	c C.duckdb_hugeint
}

type UHugeint struct {
	c C.duckdb_uhugeint
}

type Decimal struct {
	c C.duckdb_decimal
}

type QueryProgressType struct {
	c C.duckdb_query_progress_type
}

type StringT struct {
	c C.duckdb_string_t
}

type ListEntry struct {
	c C.duckdb_list_entry
}

type Column struct {
	c C.duckdb_column
}

type Vector struct {
	c C.duckdb_vector
}

func (v *Vector) Set(ptr unsafe.Pointer) {
	v.c = C.duckdb_vector(ptr)
}

type String struct {
	c C.duckdb_string
}

type Blob struct {
	c C.duckdb_blob
}

type Result struct {
	c C.duckdb_result
}

type Database struct {
	c C.duckdb_database
}

type Connection struct {
	c C.duckdb_connection
}

type PreparedStatement struct {
	c C.duckdb_prepared_statement
}

type ExtractedStatements struct {
	c C.duckdb_extracted_statements
}

type PendingResult struct {
	c C.duckdb_pending_result
}

type Appender struct {
	c C.duckdb_appender
}

type TableDescription struct {
	c C.duckdb_table_description
}

type Config struct {
	c C.duckdb_config
}

type LogicalType struct {
	c C.duckdb_logical_type
}

type CreateTypeInfo struct {
	c C.duckdb_create_type_info
}

type DataChunk struct {
	c C.duckdb_data_chunk
}

func (chunk *DataChunk) Set(ptr unsafe.Pointer) {
	chunk.c = C.duckdb_data_chunk(ptr)
}

type Value struct {
	c C.duckdb_value
}

type ProfilingInfo struct {
	c C.duckdb_profiling_info
}

type FunctionInfo struct {
	c C.duckdb_function_info
}

type ScalarFunction struct {
	c C.duckdb_scalar_function
}

type ScalarFunctionSet struct {
	c C.duckdb_scalar_function_set
}

type AggregateFunction struct {
	c C.duckdb_aggregate_function
}

type AggregateFunctionSet struct {
	c C.duckdb_aggregate_function_set
}

type AggregateState struct {
	c C.duckdb_aggregate_state
}

type TableFunction struct {
	c C.duckdb_table_function
}

type BindInfo struct {
	c C.duckdb_bind_info
}

type InitInfo struct {
	c C.duckdb_init_info
}

type CastFunction struct {
	c C.duckdb_cast_function
}

type ReplacementScanInfo struct {
	c C.duckdb_replacement_scan_info
}

type Arrow struct {
	c C.duckdb_arrow
}

type ArrowStream struct {
	c C.duckdb_arrow_stream
}

type ArrowSchema struct {
	c C.duckdb_arrow_schema
}

type ArrowArray struct {
	c C.duckdb_arrow_array
}
