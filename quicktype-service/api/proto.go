package api

import (
	"bytes"
	"html/template"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/xormplus/xorm/schemas"
)

func Table2Proto(t *schemas.Table) (string, error) {
	message := tables2Message(t)
	msgTemplate := `syntax = "proto3";

package proto;

message {{.Name}} {
{{- range .Fields}}
{{- if .IsRepeated}}
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}
`
	tmpl, err := template.New("proto").Parse(msgTemplate)
	if err != nil {
		return "", err
	}
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, message)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

func toProtoFieldTypeNameBySql(f schemas.SQLType) string {
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

func tables2Message(t *schemas.Table) *message {
	msg := &message{
		Name:   t.Name,
		Fields: []*field{},
	}
	lc := len(t.Columns())
	for i := 0; i < lc; i++ {
		f := t.Columns()[i]
		newField := &field{
			Name:       toProtoFieldName(f.Name),
			TypeName:   toProtoFieldTypeNameBySql(f.SQLType),
			IsRepeated: false,
			Order:      i + 1,
		}
		msg.Fields = append(msg.Fields, newField)
	}
	return msg
}

func toProtoFieldName(name string) string {
	if len(name) == 2 {
		return strings.ToLower(name)
	}
	r, n := utf8.DecodeRuneInString(name)
	name = string(unicode.ToLower(r)) + name[n:]
	return camel2case(name)
}

func camel2case(name string) string {
	buffer := newbuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}
