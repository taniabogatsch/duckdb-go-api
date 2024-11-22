package duckdb

/*
#include <duckdb.h>
*/
import "C"

/*
Unsure about these:

- typedef uint64_t idx_t;
- typedef struct {
	union {
		struct {
			uint32_t length;
			char prefix[4];
			char *ptr;
		} pointer;
		struct {
			uint32_t length;
			char inlined[12];
		} inlined;
	} value;
} duckdb_string_t;
*/

type Date struct {
	Days int32
}

type DateStruct struct {
	Year  int32
	Month int8
	Day   int8
}

type Time struct {
	Micros int64
}

type TimeStruct struct {
	Hour   int8
	Min    int8
	Sec    int8
	Micros int32
}

type TimeTZ struct {
	Bits uint64
}

type TimeTZStruct struct {
	Time   TimeStruct
	Offset int32
}

type Timestamp struct {
	Micros int64
}

type TimestampStruct struct {
	Date DateStruct
	Time TimeStruct
}

type Interval struct {
	Months int32
	Days   int32
	Micros int64
}

type Hugeint struct {
	Lower uint64
	Upper int64
}

type UHugeint struct {
	Lower uint64
	Upper uint64
}

type Decimal struct {
	Width uint8
	Scale uint8
	Value Hugeint
}

type QueryProgressType struct {
	Percentage         float64
	RowsProcessed      uint64
	TotalRowsToProcess uint64
}

type ListEntry struct {
	Offset uint64
	Length uint64
}
