package dialect

import "reflect"

/*
	实现将go语言中的数据类型转换为数据库中的数据类型
*/

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string // 将数据转换为该数据库中的数据类型
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect： 向全局注册dialect实例
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// GetDialect： 获取dialect实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
