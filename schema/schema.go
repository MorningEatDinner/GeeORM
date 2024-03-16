package schema

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/xiaorui/geeorm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

// GetField： 获取某个字段的名字
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse: 将对象解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type() //reflect.Indirect得到的是传入的值所指向的实际的对象
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 是否是匿名字段， 即嵌入字段  ast.IsExported(p.Name)判断是否可以导出， 即是否首字母是大写的
		if !p.Anonymous && ast.IsExported(p.Name) {
			fmt.Println(reflect.Indirect(reflect.New(p.Type)))
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))), // 将这个字段的类型转换为数据库中对应的类型， 当然，这里仅仅是记录而已
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}

	return schema
}

// RecordValues: 将对象示例转换为
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
