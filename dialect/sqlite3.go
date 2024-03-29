package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil) // 进行编译时候的类型检查， 确保实现了Dialect接口

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

// DataTypeOf： 将go语言类型映射到数据库类型
func (s sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob" // blob一般存储二进制格式文件
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datatime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind())) // 如果sql类型无效则会panic
}

// TableExistSQL: 返回在数据库中判断数据表是否存在的sql语句
func (s sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
