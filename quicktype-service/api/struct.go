package api

import (
	"bytes"
	"html/template"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/xormplus/xorm/schemas"
)

func Table2Struct(t *schemas.Table) (string, error) {
	type1 := tables2Type(t)
	msgTemplate := `package model

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type {{.Name}} struct {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined TODO
	{{.Name}} []{{.TypeName}} ` + "`gorm:\"column:{{.JsonName}}\" json:\"{{.JsonName}}\"`" + `
{{- else}}
	// {{.Name}} defined TODO
	{{.Name}} {{.TypeName}} ` + "`gorm:\"column:{{.JsonName}}\" json:\"{{.JsonName}}\"`" + `
{{- end}}
{{- end}}
}

func (t *{{.Name}}) TableName() string {
	return "{{.TableName}}"
}

func (t *{{.Name}}) Marshal(o interface{}) error {
	err := copier.Copy(o, h)
	if err != nil {
		return err
	}
	return nil
}

func (t *{{.Name}}) Unmarshal(o interface{}) error {
	err := copier.Copy(h, o)
	if err != nil {
		return err
	}
	return nil
}

func {{.Name}}MarshalLst(toValue interface{}, o interface{}) error {
	err := copier.Copy(toValue, o)
	if err != nil {
		return err
	}
	return nil
}

func {{.Name}}UnmarshalLst(o interface{}, toValue interface{}) error {
	err := copier.Copy(toValue, o)
	if err != nil {
		return err
	}
	return nil
}
`
	tmpl, err := template.New("struct").Parse(msgTemplate)
	if err != nil {
		return "", err
	}
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, type1)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

func tables2Type(t *schemas.Table) *type1 {
	msg := &type1{
		Name:      case2camel(t.Name),
		TableName: camel2case(t.Name),
		Fields:    []*field1{},
	}
	lc := len(t.Columns())
	for i := 0; i < lc; i++ {
		f := t.Columns()[i]
		newField := &field1{
			Name:       toStructFieldName(f.Name),
			TypeName:   toStructFieldTypeNameBySql(f.SQLType),
			IsRepeated: false,
			Order:      i + 1,
		}
		newField.JsonName = camel2case(newField.Name)
		newField.ColumnName = camel2case(newField.Name)
		msg.Fields = append(msg.Fields, newField)
	}
	return msg
}

func toStructFieldName(name string) string {
	if len(name) == 2 {
		return strings.ToLower(name)
	}
	r, n := utf8.DecodeRuneInString(name)
	name = string(unicode.ToLower(r)) + name[n:]
	return case2camel(name)
}

func case2camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func toStructFieldTypeNameBySql(f schemas.SQLType) string {
	switch f.Name {
	case "VARCHAR", "TEXT", "LONGTEXT", "CHAR", "MEDIUMTEXT", "TINYTEXT":
		return "string"
	case "DATETIME", "TIMESTAMP", "ENUM", "INT", "SMALLINT", "BIGINT", "TINYINT":
		return "int64"
	case "DECIMAL":
		return "string"
	case "BOOLEAN":
		return "bool"
	case "FLOAT", "DOUBLE":
		return "float64"
	case "MEDIUMBLOB", "BLOB":
		return "string"
	}
	return "string"
}
