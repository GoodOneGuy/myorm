package dialect

import (
	"reflect"
)

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(reflect.Value) string
	TableExistSQL(string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (Dialect, bool) {
	dialect, ok := dialectMap[name]
	return dialect, ok
}
