extern duckdb_ext_api_v0 *duckdb_ext_api;

typedef struct duckdb_extension_access _duckdb_extension_access;
typedef void (*set_error)(duckdb_extension_info info, const char *error);

static void _set_error(set_error f, duckdb_extension_info info, const char *error) {
	return f(info, error);
}

typedef duckdb_database *(*get_database)(duckdb_extension_info info);

static duckdb_database *_get_database(get_database f, duckdb_extension_info info) {
	return f(info);
}

typedef void *(*get_api)(duckdb_extension_info info, const char *version);

static void *_get_api(get_api f, duckdb_extension_info info, const char *version) {
	return f(info, version);
}