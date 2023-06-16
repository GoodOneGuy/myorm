package schema

import (
	"github.com/GoodOneGuy/myorm/dialect"
	"go/ast"
	"reflect"
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

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	s := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			if v, ok := p.Tag.Lookup("myorm"); ok {
				field.Tag = v
			}

			s.Fields = append(s.Fields, field)
			s.FieldNames = append(s.FieldNames, field.Name)
			s.fieldMap[p.Name] = field
		}
	}

	return s
}

func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var values []interface{}

	for _, filed := range s.Fields {
		values = append(values, destValue.FieldByName(filed.Name).Interface())
	}

	return values
}
