package postgresql

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

func genPGCatalog() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = []*catalog.Function{
		{
			Name:       "RI_FKey_cascade_del",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_cascade_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_check_ins",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_check_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_noaction_del",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_noaction_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_restrict_del",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_restrict_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_setdefault_del",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_setdefault_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_setnull_del",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "RI_FKey_setnull_upd",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name: "abbrev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cidr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "abbrev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "aclcontains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem[]"},
				},
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "acldefault",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "aclitem[]"},
		},
		{
			Name: "aclinsert",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem[]"},
				},
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "aclitem[]"},
		},
		{
			Name: "aclitemeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "aclitemin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "aclitem"},
		},
		{
			Name: "aclitemout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "aclremove",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem[]"},
				},
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "aclitem[]"},
		},
		{
			Name: "acos",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "acosd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "acosh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "amvalidate",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "any_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "any"},
		},
		{
			Name: "any_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anyarray_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "anyarray_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anyarray_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "anyelement_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "anyelement_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anyenum_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "anyenum_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anynonarray_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anynonarray"},
		},
		{
			Name: "anynonarray_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anynonarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anyrange_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name: "anyrange_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "anytextcat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anynonarray"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "area",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "area",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "area",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "array_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anynonarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_append",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_cat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_dims",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "array_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_fill",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer[]"},
				},
				{
					Type: &ast.TypeName{Name: "integer[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_fill",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "array_lower",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "array_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_ndims",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "array_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "array_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "array_position",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "array_position",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "array_positions",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer[]"},
		},
		{
			Name: "array_prepend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_remove",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_replace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "array_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "array_to_json",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "array_to_json",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "array_to_string",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "array_to_string",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "array_to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "array_upper",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "arraycontained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "arraycontains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "arrayoverlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ascii",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "asin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "asind",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "asinh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "atan",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "atan2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "atan2d",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "atand",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "atanh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "binary_upgrade_create_empty_extension",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "oid[]"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_missing_value",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_array_pg_type_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_heap_pg_class_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_index_pg_class_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_pg_authid_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_pg_enum_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_pg_type_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_toast_pg_class_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_next_toast_pg_type_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "binary_upgrade_set_record_init_privs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "bit_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "bit_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bit_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bit_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bit_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bit_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "bit_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bit_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bit_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "bit_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "bit_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "bitand",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bitcat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit varying"},
		},
		{
			Name: "bitcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "biteq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bitnot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bitor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bitshiftleft",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bitshiftright",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bittypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bittypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "bitxor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "bool",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bool",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bool_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bool_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "booland_statefunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "booleq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boollt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolor_statefunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "boolout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "boolsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "bound_box",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_above",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_above_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_add",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_below",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_below_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "box_contain",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_contain_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_contained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "box_div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_intersect",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "box_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "box_overabove",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_overbelow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_overlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_overleft",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_overright",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_same",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "box_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "box_sub",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "bpchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpchar_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpchar_pattern_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchar_pattern_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchar_pattern_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchar_pattern_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchar_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpcharcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bpchareq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchargt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpchariclike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharicnlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharicregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharicregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "bpcharle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharnlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "bpcharregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bpcharsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "bpchartypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bpchartypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "brin_desummarize_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "brin_summarize_new_values",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "brin_summarize_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "broadcast",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "btarraycmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btboolcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btbpchar_pattern_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btcharcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btfloat48cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btfloat4cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btfloat84cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btfloat8cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint24cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint28cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint2cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint42cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint48cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint4cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint82cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint84cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btint8cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btnamecmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btnametextcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btoidcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btoidvectorcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btrecordcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btrecordimagecmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "btrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "btrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "btrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "bttext_pattern_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bttextcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bttextnamecmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "bttidcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "byteacat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "byteacmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "byteaeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteage",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteagt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteain",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "byteale",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bytealike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "bytealt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteane",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteanlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "byteaout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "byteasend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "cardinality",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "cash_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "cash_div_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cash_div_flt4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_div_flt8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_div_int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_div_int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_div_int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_mul_flt4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_mul_flt8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_mul_int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_mul_int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_mul_int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cash_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "cash_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cash_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "cash_words",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "cashlarger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cashsmaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "cbrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "ceil",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "ceil",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "ceiling",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "ceiling",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "char"},
		},
		{
			Name: "char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "char"},
		},
		{
			Name: "char_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "char_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "character_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "character_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "chareq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "charge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "chargt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "charin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "char"},
		},
		{
			Name: "charle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "charlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "charne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "charout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "charsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "chr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "cideq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cid"},
				},
				{
					Type: &ast.TypeName{Name: "cid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "cidin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cid"},
		},
		{
			Name: "cidout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "cidr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cidr"},
		},
		{
			Name: "cidr_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cidr"},
		},
		{
			Name: "cidr_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cidr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "cidr_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cidr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "cidsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "circle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle_above",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_add_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle_below",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "circle_contain",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_contain_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_contained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "circle_div_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_mul_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name: "circle_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "circle_overabove",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_overbelow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_overlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_overleft",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_overright",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_same",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "circle_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "circle_sub_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "circle"},
		},
		{
			Name:       "clock_timestamp",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "close_lb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_ls",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_lseg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_pb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_ps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_sb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "close_sl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "col_description",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "convert",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "convert_from",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "convert_to",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "corr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cos",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cosd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cosh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cotd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "count",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "covar_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "covar_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "cstring_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "cstring_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "cstring_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name:       "cume_dist",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "current_database",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name:       "current_query",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name:       "current_schema",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "current_schemas",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name[]"},
		},
		{
			Name: "current_setting",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "current_setting",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name:       "current_user",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "currtid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "currtid2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "currval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "cursor_to_xml",
			Args: []*catalog.Argument{
				{
					Name: "cursor",
					Type: &ast.TypeName{Name: "refcursor"},
				},
				{
					Name: "count",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "cursor_to_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "cursor",
					Type: &ast.TypeName{Name: "refcursor"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "database_to_xml",
			Args: []*catalog.Argument{
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "database_to_xml_and_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "database_to_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "date_cmp_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "date_cmp_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "date_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_eq_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_eq_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_ge_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_ge_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_gt_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_gt_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_le_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_le_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_lt_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_lt_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "date_mi_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "date_mii",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_ne_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_ne_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "date_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "date_pl_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "date_pli",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "date_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "date_trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "date_trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "date_trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "date_trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "daterange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "daterange"},
		},
		{
			Name: "daterange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "daterange"},
		},
		{
			Name: "daterange_canonical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "daterange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "daterange"},
		},
		{
			Name: "daterange_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "datetime_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "datetimetz_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "dcbrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "decode",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "degrees",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "dense_rank",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "dexp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "diagonal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "lseg"},
		},
		{
			Name: "diameter",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_cpoint",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_cpoly",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_lb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_pb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_pc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_polyp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_ppath",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_ppoly",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_ps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_sb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dist_sl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "dlog1",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dlog10",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "domain_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "any"},
		},
		{
			Name: "dpow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dround",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dsqrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "dtrunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "elem_contained_by_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "encode",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "enum_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "enum_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_first",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "enum_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "enum_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "enum_last",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "enum_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "enum_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "enum_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "enum_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "enum_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "enum_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "event_trigger_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "event_trigger"},
		},
		{
			Name: "event_trigger_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "event_trigger"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "every",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "exp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "exp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "factorial",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "family",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "fdw_handler_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "fdw_handler"},
		},
		{
			Name: "fdw_handler_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "fdw_handler"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "first_value",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float48div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float48eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float48mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float48ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float48pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float4_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "float4abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float4out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "float4pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "float4smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float4up",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float84div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float84eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float84mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float84ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float84pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "float8_avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_combine",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "float8_corr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_covar_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_covar_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "float8_regr_avgx",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_avgy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_combine",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "float8_regr_intercept",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_r2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_slope",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_sxx",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_sxy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_regr_syy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8_var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "float8out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "float8pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "float8smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "float8up",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "floor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "floor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "flt4_mul_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "flt8_mul_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "fmgr_c_validator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "fmgr_internal_validator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "fmgr_sql_validator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "format",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "format_type",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "generate_series",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "generate_subscripts",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "generate_subscripts",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "get_bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "get_bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "get_byte",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "get_current_ts_config",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "regconfig"},
		},
		{
			Name:       "getdatabaseencoding",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name:       "getpgusername",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "gin_clean_pending_list",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "gin_cmp_tslexeme",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "gin_compare_jsonb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "gtsvectorin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "gtsvector"},
		},
		{
			Name: "gtsvectorout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "gtsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_any_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_column_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_database_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_foreign_data_wrapper_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_function_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_language_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_schema_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_sequence_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_server_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_table_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_tablespace_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "has_type_privilege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "hash_aclitem",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hash_aclitem_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "aclitem"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hash_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hash_array_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hash_numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hash_numeric_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hash_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hash_range_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashbpchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashbpcharextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashcharextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashenum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashenumextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashfloat4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashfloat4extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashfloat8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashfloat8extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashinet",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashinetextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashint2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashint2extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashint4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashint4extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashint8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashint8extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashmacaddr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashmacaddr8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashmacaddr8extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashmacaddrextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashname",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashnameextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashoid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashoidextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashoidvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashoidvectorextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashtext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashtextextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "hashtid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "hashtidextended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "height",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "host",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "hostmask",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "in_range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "index_am_handler_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "index_am_handler"},
		},
		{
			Name: "index_am_handler_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "index_am_handler"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name:       "inet_client_addr",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name:       "inet_client_port",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "inet_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "inet_merge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cidr"},
		},
		{
			Name: "inet_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "inet_same_family",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "inet_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name:       "inet_server_addr",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name:       "inet_server_port",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "inetand",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "inetmi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "inetmi_int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "inetnot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "inetor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "inetpl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "initcap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int24div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int24eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int24mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int24ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int24pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int28div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int28eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int28mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int28ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int28pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int2_avg_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint[]"},
		},
		{
			Name: "int2_avg_accum_inv",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint[]"},
		},
		{
			Name: "int2_mul_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "int2_sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int2abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2int4_sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int2larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int2not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "int2pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "int2shl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2shr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2up",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int2vectorin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int2vector"},
		},
		{
			Name: "int2vectorout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "int2vector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "int2vectorsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "int2vector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "int2xor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int42div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int42eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int42mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int42ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int42pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int48div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int48eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int48mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int48ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int48pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int4_avg_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint[]"},
		},
		{
			Name: "int4_avg_accum_inv",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint[]"},
		},
		{
			Name: "int4_avg_combine",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint[]"},
		},
		{
			Name: "int4_mul_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "int4_sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int4abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4inc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int4not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "int4pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int4range"},
		},
		{
			Name: "int4range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int4range"},
		},
		{
			Name: "int4range_canonical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "int4range"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int4range"},
		},
		{
			Name: "int4range_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "int4send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "int4shl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4shr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4up",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int4xor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int82div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int82eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int82mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int82ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int82pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int84div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int84eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int84mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int84ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int84pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8_avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "int8_mul_cash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "int8_sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "int8abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8dec",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8dec_any",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8inc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8inc_any",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8inc_float8_float8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "int8not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "int8pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8pl_inet",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "int8range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int8range"},
		},
		{
			Name: "int8range",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int8range"},
		},
		{
			Name: "int8range_canonical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "int8range"},
				},
			},
			ReturnType: &ast.TypeName{Name: "int8range"},
		},
		{
			Name: "int8range_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "int8send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "int8shl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8shr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8up",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "int8xor",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "integer_pl_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "inter_lb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "inter_sb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "inter_sl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_accum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval[]"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval[]"},
		},
		{
			Name: "interval_accum_inv",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval[]"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval[]"},
		},
		{
			Name: "interval_avg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "interval_combine",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval[]"},
				},
				{
					Type: &ast.TypeName{Name: "interval[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval[]"},
		},
		{
			Name: "interval_div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "interval_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "interval_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "interval_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "interval_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_pl_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "interval_pl_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "interval_pl_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "interval_pl_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "interval_pl_timetz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "interval_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "interval_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "interval_um",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "intervaltypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "intervaltypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "isclosed",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isempty",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isfinite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isfinite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isfinite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isfinite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ishorizontal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ishorizontal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ishorizontal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isopen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isparallel",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isparallel",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isperp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isperp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isvertical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isvertical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "isvertical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "json_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_array_element",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name: "element_index",
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_array_element_text",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name: "element_index",
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "json_array_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "json_build_array",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name:       "json_build_object",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_object",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_object",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_object_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_object_field",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name: "field_name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_object_field_text",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name: "field_name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "json_object_keys",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "json_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "json_populate_record",
			Args: []*catalog.Argument{
				{
					Name: "base",
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name:       "use_json_as_text",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "json_populate_recordset",
			Args: []*catalog.Argument{
				{
					Name: "base",
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Name:       "use_json_as_text",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "json_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "json_strip_nulls",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "json_to_record",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "json_to_recordset",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "json_to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "json_to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "json_typeof",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonb_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_array_element",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "element_index",
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_array_element_text",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "element_index",
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonb_array_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "jsonb_build_array",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name:       "jsonb_build_object",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "jsonb_concat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_contained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_contains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_delete",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_delete",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_delete_path",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_exists",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_exists_all",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_exists_any",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "jsonb_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "jsonb_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_insert",
			Args: []*catalog.Argument{
				{
					Name: "jsonb_in",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Name: "replacement",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "insert_after",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_object",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_object",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_object_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_object_field",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "field_name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_object_field_text",
			Args: []*catalog.Argument{
				{
					Name: "from_json",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "field_name",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonb_object_keys",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonb_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "jsonb_path_exists",
			Args: []*catalog.Argument{
				{
					Name: "target",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "jsonpath"},
				},
				{
					Name:       "vars",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "silent",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_path_exists_opr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonpath"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_path_match",
			Args: []*catalog.Argument{
				{
					Name: "target",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "jsonpath"},
				},
				{
					Name:       "vars",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "silent",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_path_match_opr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonpath"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "jsonb_path_query",
			Args: []*catalog.Argument{
				{
					Name: "target",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "jsonpath"},
				},
				{
					Name:       "vars",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "silent",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_path_query_array",
			Args: []*catalog.Argument{
				{
					Name: "target",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "jsonpath"},
				},
				{
					Name:       "vars",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "silent",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_path_query_first",
			Args: []*catalog.Argument{
				{
					Name: "target",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "jsonpath"},
				},
				{
					Name:       "vars",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "silent",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_populate_record",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "jsonb_populate_recordset",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "jsonb_pretty",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonb_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "jsonb_set",
			Args: []*catalog.Argument{
				{
					Name: "jsonb_in",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name: "path",
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Name: "replacement",
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Name:       "create_if_missing",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_strip_nulls",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "jsonb_to_record",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "jsonb_to_recordset",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "jsonb_to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "jsonb_to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "jsonb_typeof",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "jsonpath_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonpath"},
		},
		{
			Name: "jsonpath_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonpath"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "jsonpath_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonpath"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "justify_days",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "justify_hours",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "justify_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "lag",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "lag",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "lag",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "language_handler_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "language_handler"},
		},
		{
			Name: "language_handler_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "language_handler"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "last_value",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name:       "lastval",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "lead",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "lead",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "lead",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "like",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "like",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "like",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "like_escape",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "like_escape",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "line",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "line"},
		},
		{
			Name: "line_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "line_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "line_horizontal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "line_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "line"},
		},
		{
			Name: "line_interpt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "line_intersect",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "line_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "line_parallel",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "line_perp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "line_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "line_vertical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ln",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "ln",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "lo_close",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_creat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "lo_create",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "lo_export",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_from_bytea",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "lo_get",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "lo_get",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "lo_import",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "lo_import",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "lo_lseek",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_lseek64",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "lo_open",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_put",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "lo_tell",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_tell64",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "lo_truncate",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_truncate64",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lo_unlink",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "log",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "log",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "log",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "log10",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "log10",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "loread",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "lower",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "lower",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "lower_inc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lower_inf",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lowrite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "lpad",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "lpad",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "lseg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "lseg"},
		},
		{
			Name: "lseg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "lseg"},
		},
		{
			Name: "lseg_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "lseg_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "lseg_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_horizontal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "lseg"},
		},
		{
			Name: "lseg_interpt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "lseg_intersect",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "lseg_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "lseg_parallel",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_perp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "lseg_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "lseg_vertical",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ltrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "ltrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "macaddr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "macaddr8",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr8_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr8_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "macaddr8_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr8_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr8_not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr8_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr8_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "macaddr8_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "macaddr8_set7bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "macaddr_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "macaddr_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "macaddr_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "macaddr_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "macaddr_not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "macaddr_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "macaddr_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "macaddr_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "make_date",
			Args: []*catalog.Argument{
				{
					Name: "year",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "month",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "day",
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "make_interval",
			Args: []*catalog.Argument{
				{
					Name:       "years",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "months",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "weeks",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "days",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "hours",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "mins",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
				{
					Name:       "secs",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "make_time",
			Args: []*catalog.Argument{
				{
					Name: "hour",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "min",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "sec",
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "make_timestamp",
			Args: []*catalog.Argument{
				{
					Name: "year",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "month",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "mday",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "hour",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "min",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "sec",
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "make_timestamptz",
			Args: []*catalog.Argument{
				{
					Name: "year",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "month",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "mday",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "hour",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "min",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "sec",
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "make_timestamptz",
			Args: []*catalog.Argument{
				{
					Name: "year",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "month",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "mday",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "hour",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "min",
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Name: "sec",
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Name: "timezone",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "makeaclitem",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "aclitem"},
		},
		{
			Name: "masklen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "max",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "md5",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "md5",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyenum"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "min",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
		{
			Name: "mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "mode",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "money",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "money",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "money",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "mul_d_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "mxid_age",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "name",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "name",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "name",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "nameconcatoid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "nameeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameeqtext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namege",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namegetext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namegt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namegttext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameiclike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameicnlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameicregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameicregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namein",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "namele",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameletext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namelike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namelt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namelttext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namene",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namenetext",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namenlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "nameregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nameregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "namesend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "netmask",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "network",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cidr"},
		},
		{
			Name: "network_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "network_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "network_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_overlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "network_sub",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_subeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_sup",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "network_supeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "nextval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "notlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "notlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "notlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "now",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "npoints",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "npoints",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "nth_value",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "ntile",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_abs",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_add",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "numeric_div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_div_trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_exp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_fac",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_inc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_ln",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_log",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_mod",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "numeric_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "numeric_power",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "numeric_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_sqrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_sub",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_uminus",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numeric_uplus",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "numerictypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "numerictypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "numnode",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "numrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numrange"},
		},
		{
			Name: "numrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numrange"},
		},
		{
			Name: "numrange_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "obj_description",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "obj_description",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "octet_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "octet_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "octet_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "octet_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "oideq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "oidlarger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "oidle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "oidsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "oidsmaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "oidvectoreq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oidvector"},
		},
		{
			Name: "oidvectorle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "oidvectorout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "oidvectorsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "oidvectortypes",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oidvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "on_pb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "on_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "on_ppath",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "on_ps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "on_sb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "on_sl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
				{
					Type: &ast.TypeName{Name: "line"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "opaque_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "opaque"},
		},
		{
			Name: "opaque_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "opaque"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "overlay",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "parse_ident",
			Args: []*catalog.Argument{
				{
					Name: "str",
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Name:       "strict",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "path",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_add",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_add_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "path_contain_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "path_div_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_inter",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "path_mul_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "path_n_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_n_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_n_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_n_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_n_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "path_npoints",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "path_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "path_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "path_sub_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "pclose",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name:       "percent_rank",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "percentile_cont",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "percentile_cont",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision[]"},
		},
		{
			Name: "percentile_cont",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "percentile_cont",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval[]"},
		},
		{
			Name: "percentile_disc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision[]"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyarray"},
		},
		{
			Name: "percentile_disc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "pg_advisory_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_unlock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_advisory_unlock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_advisory_unlock_all",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_unlock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_advisory_unlock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_advisory_xact_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_xact_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_xact_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_advisory_xact_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name:       "pg_backend_pid",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "pg_backup_start_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_blocking_pids",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer[]"},
		},
		{
			Name: "pg_cancel_backend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_char_to_encoding",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "pg_client_encoding",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "pg_collation_actual_version",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_collation_for",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_collation_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_column_is_updatable",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_column_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "pg_conf_load_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_conversion_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_create_restore_point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_current_logfile",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name:       "pg_current_logfile",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name:       "pg_current_wal_flush_lsn",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_current_wal_insert_lsn",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_current_wal_lsn",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_database_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_database_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_ddl_command_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_ddl_command"},
		},
		{
			Name: "pg_ddl_command_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_ddl_command"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_ddl_command_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_ddl_command"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_dependencies_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_dependencies"},
		},
		{
			Name: "pg_dependencies_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_dependencies"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_dependencies_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_dependencies"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_describe_object",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_drop_replication_slot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_encoding_max_length",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_encoding_to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name:       "pg_event_trigger_table_rewrite_reason",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name:       "pg_export_snapshot",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_extension_config_dump",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_filenode_relation",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name: "pg_function_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_get_constraintdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_constraintdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_expr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_node_tree"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_expr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_node_tree"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_function_arg_default",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_function_arguments",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_function_identity_arguments",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_function_result",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_functiondef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_indexdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_indexdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_partition_constraintdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_partkeydef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_replica_identity_index",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name: "pg_get_ruledef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_ruledef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_serial_sequence",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_statisticsobjdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_triggerdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_triggerdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_userbyid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "pg_get_viewdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_viewdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_viewdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_viewdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_get_viewdef",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_has_role",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_import_system_collations",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regnamespace"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_index_column_has_property",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_index_has_property",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_indexam_has_property",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_indexam_progress_phasename",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_indexes_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_is_in_backup",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_is_in_recovery",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_is_other_temp_schema",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_is_wal_replay_paused",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_isolation_test_session_is_blocked",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_jit_available",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_last_wal_receive_lsn",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_last_wal_replay_lsn",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_last_xact_replay_timestamp",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name:       "pg_listening_channels",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_logical_emit_message",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_logical_emit_message",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_ls_dir",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_ls_dir",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_lsn_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_lsn_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_lsn_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_lsn_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_lsn_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "pg_lsn_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_lsn_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_lsn_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_mcv_list_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_mcv_list"},
		},
		{
			Name: "pg_mcv_list_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_mcv_list"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_mcv_list_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_mcv_list"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name:       "pg_my_temp_schema",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_ndistinct_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_ndistinct"},
		},
		{
			Name: "pg_ndistinct_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_ndistinct"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_ndistinct_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_ndistinct"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_nextoid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_node_tree_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_node_tree"},
		},
		{
			Name: "pg_node_tree_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_node_tree"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "pg_node_tree_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_node_tree"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name:       "pg_notification_queue_usage",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_notify",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_opclass_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_operator_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_opfamily_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_partition_root",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name:       "pg_postmaster_start_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_promote",
			Args: []*catalog.Argument{
				{
					Name:       "wait",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
				{
					Name:       "wait_seconds",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_read_binary_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_read_binary_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_read_binary_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "pg_read_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_read_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_read_file",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_read_file_old",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_relation_filenode",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_relation_filepath",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_relation_is_publishable",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_relation_is_updatable",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_relation_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_relation_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_reload_conf",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_replication_origin_advance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_replication_origin_create",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_replication_origin_drop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_replication_origin_oid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_replication_origin_progress",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_replication_origin_session_is_setup",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_replication_origin_session_progress",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_replication_origin_session_reset",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_replication_origin_session_setup",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name:       "pg_replication_origin_xact_reset",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_replication_origin_xact_setup",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name:       "pg_rotate_logfile",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_rotate_logfile_old",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_safe_snapshot_blocking_pids",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer[]"},
		},
		{
			Name: "pg_sequence_last_value",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_size_bytes",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_size_pretty",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_size_pretty",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_sleep",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_sleep_for",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_sleep_until",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_start_backup",
			Args: []*catalog.Argument{
				{
					Name: "label",
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Name:       "fast",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
				{
					Name:       "exclusive",
					HasDefault: true,
					Type:       &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_stat_clear_snapshot",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_stat_get_analyze_count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_autoanalyze_count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_autovacuum_count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_backend_activity",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_stat_get_backend_activity_start",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_backend_client_addr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "pg_stat_get_backend_client_port",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_stat_get_backend_dbid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name:       "pg_stat_get_backend_idset",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_stat_get_backend_pid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_stat_get_backend_start",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_backend_userid",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_stat_get_backend_wait_event",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_stat_get_backend_wait_event_type",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_stat_get_backend_xact_start",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name:       "pg_stat_get_bgwriter_buf_written_checkpoints",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_bgwriter_buf_written_clean",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_bgwriter_maxwritten_clean",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_bgwriter_requested_checkpoints",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_bgwriter_stat_reset_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name:       "pg_stat_get_bgwriter_timed_checkpoints",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_blocks_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_blocks_hit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_buf_alloc",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_buf_fsync_backend",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_buf_written_backend",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_checkpoint_sync_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "pg_stat_get_checkpoint_write_time",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_db_blk_read_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_db_blk_write_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_db_blocks_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_blocks_hit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_checksum_failures",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_checksum_last_failure",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_db_conflict_all",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_conflict_bufferpin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_conflict_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_conflict_snapshot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_conflict_startup_deadlock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_conflict_tablespace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_deadlocks",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_numbackends",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_stat_get_db_stat_reset_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_db_temp_bytes",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_temp_files",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_tuples_deleted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_tuples_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_tuples_inserted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_tuples_returned",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_tuples_updated",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_xact_commit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_db_xact_rollback",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_dead_tuples",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_function_calls",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_function_self_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_function_total_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_last_analyze_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_last_autoanalyze_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_last_autovacuum_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_last_vacuum_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_live_tuples",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_mod_since_analyze",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_numscans",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_get_snapshot_timestamp",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "pg_stat_get_tuples_deleted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_tuples_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_tuples_hot_updated",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_tuples_inserted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_tuples_returned",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_tuples_updated",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_vacuum_count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_blocks_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_blocks_hit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_function_calls",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_function_self_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_xact_function_total_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pg_stat_get_xact_numscans",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_deleted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_fetched",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_hot_updated",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_inserted",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_returned",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_stat_get_xact_tuples_updated",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_stat_reset",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_stat_reset_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_stat_reset_single_function_counters",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_stat_reset_single_table_counters",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_statistics_obj_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "pg_stop_backup",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name:       "pg_switch_wal",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "pg_lsn"},
		},
		{
			Name: "pg_table_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_table_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_tablespace_databases",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "oid"},
		},
		{
			Name: "pg_tablespace_location",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_tablespace_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_tablespace_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_terminate_backend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_total_relation_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "pg_trigger_depth",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "pg_try_advisory_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_xact_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_xact_lock",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_xact_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_try_advisory_xact_lock_shared",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_ts_config_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_ts_dict_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_ts_parser_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_ts_template_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_type_is_visible",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pg_typeof",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "any"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regtype"},
		},
		{
			Name: "pg_wal_lsn_diff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name:       "pg_wal_replay_pause",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name:       "pg_wal_replay_resume",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "pg_walfile_name",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "pg_lsn"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "pg_xact_commit_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "phraseto_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "phraseto_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name:       "pi",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "plainto_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "plainto_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name:       "plpgsql_call_handler",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "language_handler"},
		},
		{
			Name: "plpgsql_validator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "lseg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_above",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_add",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_below",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "point_div",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_horiz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_mul",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "point_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "point_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "point_sub",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "point_vert",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_above",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_below",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "point"},
		},
		{
			Name: "poly_contain",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_contain_pt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_contained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_distance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "poly_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "polygon"},
		},
		{
			Name: "poly_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_npoints",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "poly_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "poly_overabove",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_overbelow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_overlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_overleft",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_overright",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_same",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "poly_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "polygon",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "polygon"},
		},
		{
			Name: "polygon",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "polygon"},
		},
		{
			Name: "polygon",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "polygon"},
		},
		{
			Name: "polygon",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "polygon"},
		},
		{
			Name: "popen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "path"},
				},
			},
			ReturnType: &ast.TypeName{Name: "path"},
		},
		{
			Name: "position",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "position",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "position",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "bit"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "postgresql_fdw_validator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "pow",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "power",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "power",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "pt_contained_circle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "pt_contained_poly",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "query_to_xml",
			Args: []*catalog.Argument{
				{
					Name: "query",
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "query_to_xml_and_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "query",
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "query_to_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "query",
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "querytree",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "quote_ident",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "quote_literal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "quote_literal",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "quote_nullable",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "quote_nullable",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "radians",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "radius",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "circle"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "random",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "range_adjacent",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_after",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_before",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "range_contained_by",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_contains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_contains_elem",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name: "range_intersect",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name: "range_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_merge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name: "range_minus",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name: "range_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "range_overlaps",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_overleft",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_overright",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "range_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "range_union",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyrange"},
		},
		{
			Name:       "rank",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "record_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_image_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "record_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "record_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "record_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regclass",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name: "regclassin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name: "regclassout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regclasssend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regconfigin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regconfig"},
		},
		{
			Name: "regconfigout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regconfigsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regdictionaryin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regdictionary"},
		},
		{
			Name: "regdictionaryout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regdictionary"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regdictionarysend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regdictionary"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regexp_match",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_match",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_matches",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_matches",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_replace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "regexp_replace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "regexp_split_to_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_split_to_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "regexp_split_to_table",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "regexp_split_to_table",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "regnamespacein",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regnamespace"},
		},
		{
			Name: "regnamespaceout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regnamespace"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regnamespacesend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regnamespace"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regoperatorin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regoperator"},
		},
		{
			Name: "regoperatorout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regoperator"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regoperatorsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regoperator"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regoperin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regoper"},
		},
		{
			Name: "regoperout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regoper"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regopersend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regoper"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regprocedurein",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regprocedure"},
		},
		{
			Name: "regprocedureout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regprocedure"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regproceduresend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regprocedure"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regprocin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regproc"},
		},
		{
			Name: "regprocout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regproc"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regprocsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regproc"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regr_avgx",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_avgy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_count",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "regr_intercept",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_r2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_slope",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_sxx",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_sxy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regr_syy",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "regrolein",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regrole"},
		},
		{
			Name: "regroleout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regrole"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regrolesend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regrole"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "regtypein",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regtype"},
		},
		{
			Name: "regtypeout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regtype"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "regtypesend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regtype"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "repeat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "replace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "reverse",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "round",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "round",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "round",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "row_number",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "row_security_active",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "row_security_active",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "row_to_json",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "row_to_json",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "record"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "rpad",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "rpad",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "rtrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "rtrim",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "scale",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "schema_to_xml",
			Args: []*catalog.Argument{
				{
					Name: "schema",
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "schema_to_xml_and_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "schema",
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "schema_to_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "schema",
					Type: &ast.TypeName{Name: "name"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name:       "session_user",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "name"},
		},
		{
			Name: "set_bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "set_bit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "set_byte",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "set_config",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "set_masklen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cidr"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cidr"},
		},
		{
			Name: "set_masklen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "inet"},
		},
		{
			Name: "setseed",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "setval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "setval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "setweight",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "setweight",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "sha224",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "sha256",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "sha384",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "sha512",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "shell_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "opaque"},
		},
		{
			Name: "shell_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "opaque"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "shobj_description",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "sign",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "sign",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "similar_escape",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "sin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "sind",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "sinh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "slope",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "point"},
				},
				{
					Type: &ast.TypeName{Name: "point"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "spg_poly_quad_compress",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "polygon"},
				},
			},
			ReturnType: &ast.TypeName{Name: "box"},
		},
		{
			Name: "split_part",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "sqrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "sqrt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "starts_with",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "statement_timestamp",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "stddev_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "string_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "string_agg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "string_to_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "string_to_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "strip",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "strpos",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "substr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "substr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "substr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "substr",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bytea"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "substring",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "money"},
				},
			},
			ReturnType: &ast.TypeName{Name: "money"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "sum",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name:       "suppress_redundant_updates_trigger",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name: "table_am_handler_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "table_am_handler"},
		},
		{
			Name: "table_am_handler_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "table_am_handler"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "table_to_xml",
			Args: []*catalog.Argument{
				{
					Name: "tbl",
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "table_to_xml_and_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "tbl",
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "table_to_xmlschema",
			Args: []*catalog.Argument{
				{
					Name: "tbl",
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Name: "nulls",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "tableforest",
					Type: &ast.TypeName{Name: "boolean"},
				},
				{
					Name: "targetns",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "tan",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "tand",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "tanh",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "char"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "inet"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "text_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_pattern_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_pattern_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_pattern_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_pattern_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "text_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "textanycat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "anynonarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "textcat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "texteq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "texteqname",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textgename",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textgtname",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "texticlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "texticnlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "texticregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "texticregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "textlen",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "textlename",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textltname",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textnename",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textnlike",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "textregexeq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textregexne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "textsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "tideq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "tidlarger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "tidle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tidout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "tidsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "tidsmaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tid"},
				},
				{
					Type: &ast.TypeName{Name: "tid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tid"},
		},
		{
			Name: "time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "time_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "time_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "time_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_mi_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time_mi_time",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "time_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "time_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "time_pl_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "time_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "time_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time without time zone"},
		},
		{
			Name: "timedate_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name:       "timeofday",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamp_cmp_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamp_cmp_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamp_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_eq_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_eq_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_ge_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_ge_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_gt_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_gt_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamp_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "timestamp_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_le_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_le_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_lt_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_lt_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "timestamp_mi_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_ne_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_ne_timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamp_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timestamp_pl_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamp_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "timestamp_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timestamptypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamptypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "date"},
				},
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamptz_cmp_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamptz_cmp_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamptz_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_eq_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_eq_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_ge_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_ge_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_gt_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_gt_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_le_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_le_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_lt_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_lt_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_mi",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "interval"},
		},
		{
			Name: "timestamptz_mi_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_ne_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_ne_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timestamptz_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timestamptz_pl_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptz_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "timestamptz_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timestamptztypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timestamptztypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timetypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timetypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timetz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timetz_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timetz_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "timetz_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz_larger",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_mi_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "timetz_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timetz_pl_interval",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "interval"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetz_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "timetz_smaller",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timetzdate_pl",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "date"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timetztypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "timetztypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "time with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "time with time zone"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "timezone",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp without time zone"},
		},
		{
			Name: "to_ascii",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_ascii",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_ascii",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "interval"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_char",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_date",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "date"},
		},
		{
			Name: "to_hex",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_hex",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "to_json",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "to_jsonb",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "to_number",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "to_regclass",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regclass"},
		},
		{
			Name: "to_regnamespace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regnamespace"},
		},
		{
			Name: "to_regoper",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regoper"},
		},
		{
			Name: "to_regoperator",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regoperator"},
		},
		{
			Name: "to_regproc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regproc"},
		},
		{
			Name: "to_regprocedure",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regprocedure"},
		},
		{
			Name: "to_regrole",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regrole"},
		},
		{
			Name: "to_regtype",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "regtype"},
		},
		{
			Name: "to_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "to_timestamp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "to_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "to_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "to_tsvector",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name:       "transaction_timestamp",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "timestamp with time zone"},
		},
		{
			Name: "translate",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "trigger_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name: "trigger_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "trigger"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr"},
		},
		{
			Name: "trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "macaddr8"},
				},
			},
			ReturnType: &ast.TypeName{Name: "macaddr8"},
		},
		{
			Name: "trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "trunc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "ts_delete",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "ts_delete",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "ts_filter",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "char[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "json"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "json"},
		},
		{
			Name: "ts_headline",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "jsonb"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "jsonb"},
		},
		{
			Name: "ts_lexize",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regdictionary"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name: "ts_match_qv",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ts_match_tq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ts_match_tt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ts_match_vq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "ts_rank",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real[]"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real[]"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank_cd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank_cd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real[]"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank_cd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real[]"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rank_cd",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "ts_rewrite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "ts_rewrite",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsm_handler_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsm_handler"},
		},
		{
			Name: "tsm_handler_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsm_handler"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "tsq_mcontained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsq_mcontains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_and",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsquery_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "tsquery_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsquery_not",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsquery_or",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsquery_phrase",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsquery_phrase",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsqueryin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "tsqueryout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "tsquerysend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsquery"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "tsrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsrange"},
		},
		{
			Name: "tsrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsrange"},
		},
		{
			Name: "tsrange_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp without time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "tstzrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tstzrange"},
		},
		{
			Name: "tstzrange",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tstzrange"},
		},
		{
			Name: "tstzrange_subdiff",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
				{
					Type: &ast.TypeName{Name: "timestamp with time zone"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "tsvector_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "tsvector_concat",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "tsvector_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "tsvector_to_array",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text[]"},
		},
		{
			Name:       "tsvector_update_trigger",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name:       "tsvector_update_trigger_column",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name: "tsvectorin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsvector"},
		},
		{
			Name: "tsvectorout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "tsvectorsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "tsvector"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name:       "txid_current",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "txid_current_if_assigned",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name:       "txid_current_snapshot",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "txid_snapshot"},
		},
		{
			Name: "txid_snapshot_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "txid_snapshot"},
		},
		{
			Name: "txid_snapshot_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "txid_snapshot_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "txid_snapshot_xip",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "txid_snapshot_xmax",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "txid_snapshot_xmin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "txid_status",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "txid_visible_in_snapshot",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
				{
					Type: &ast.TypeName{Name: "txid_snapshot"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name:       "unique_key_recheck",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "trigger"},
		},
		{
			Name: "unknownin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "unknown"},
		},
		{
			Name: "unknownout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "unknown"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "unknownsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "unknown"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "unnest",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "upper",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "upper",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "anyelement"},
		},
		{
			Name: "upper_inc",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "upper_inf",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyrange"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "uuid_eq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_hash",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "uuid_hash_extended",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "uuid_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "uuid"},
		},
		{
			Name: "uuid_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_ne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "uuid_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "uuid_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "uuid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "var_pop",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "var_samp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "varbit",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit varying"},
		},
		{
			Name: "varbit_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bit varying"},
		},
		{
			Name: "varbit_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "varbit_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "varbitcmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "varbiteq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbitge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbitgt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbitle",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbitlt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbitne",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
				{
					Type: &ast.TypeName{Name: "bit varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "varbittypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "varbittypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "varchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "name"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character varying"},
		},
		{
			Name: "varchar",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character varying"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
				{
					Type: &ast.TypeName{Name: "boolean"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character varying"},
		},
		{
			Name: "varcharin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
				{
					Type: &ast.TypeName{Name: "oid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "character varying"},
		},
		{
			Name: "varcharout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "varcharsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "character varying"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "varchartypmodin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "varchartypmodout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "smallint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "numeric"},
		},
		{
			Name: "variance",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "real"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name:       "version",
			Args:       []*catalog.Argument{},
			ReturnType: &ast.TypeName{Name: "text"},
		},
		{
			Name: "void_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "void"},
		},
		{
			Name: "void_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "void"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "void_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "void"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "websearch_to_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "websearch_to_tsquery",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regconfig"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "tsquery"},
		},
		{
			Name: "width",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "box"},
				},
			},
			ReturnType: &ast.TypeName{Name: "double precision"},
		},
		{
			Name: "width_bucket",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "width_bucket",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyelement"},
				},
				{
					Type: &ast.TypeName{Name: "anyarray"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "width_bucket",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "double precision"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "xideq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xideqint4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xidin",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xid"},
		},
		{
			Name: "xidneq",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xidneqint4",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
				{
					Type: &ast.TypeName{Name: "integer"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xidout",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "xidsend",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xid"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "xml",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "xml_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "xml_is_well_formed",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xml_is_well_formed_content",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xml_is_well_formed_document",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xml_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "xml_send",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bytea"},
		},
		{
			Name: "xmlagg",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "xmlcomment",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "xmlconcat2",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml"},
		},
		{
			Name: "xmlexists",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xmlvalidate",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "xml"},
				},
				{
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xpath",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml[]"},
		},
		{
			Name: "xpath",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "xml[]"},
		},
		{
			Name: "xpath_exists",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "xpath_exists",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "text"},
				},
				{
					Type: &ast.TypeName{Name: "xml"},
				},
				{
					Type: &ast.TypeName{Name: "text[]"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
	}
	return s
}