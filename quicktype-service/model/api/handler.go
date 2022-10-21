package api

import (
	"bytes"
	"html/template"

	"github.com/xormplus/xorm/schemas"
)

func Table2Handler(t *schemas.Table) (string, error) {
	message := tables2Handler(t)
	msgTemplate := `syntax = "proto3";

package proto;

service Handler {
	// Insert{{.Name}} defined todo
	rpc Insert{{.Name}} (Insert{{.Name}}Request) returns (Insert{{.Name}}Response) {}
	// Delete{{.Name}} defined todo
	rpc Delete{{.Name}} (Delete{{.Name}}Request) returns (Delete{{.Name}}Response) {}
	// Update{{.Name}} defined todo
	rpc Update{{.Name}} (Update{{.Name}}Request) returns (Update{{.Name}}Response) {}
	// Query{{.Name}} defined todo
	rpc Query{{.Name}} (Query{{.Name}}Request) returns (Query{{.Name}}Response) {}
	// Query{{.Name}}Detail defined todo
	rpc Query{{.Name}}Detail (Query{{.Name}}DetailRequest) returns (Query{{.Name}}DetailResponse) {}
}

message Insert{{.Name}}Request {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined todo
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	// {{.Name}} defined todo
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}

message Insert{{.Name}}Response {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined todo
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	// {{.Name}} defined todo
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}

message Delete{{.Name}}Request {
	uint32 id = 1 [json_name = "id"];
}

message Delete{{.Name}}Response {
	uint32 id = 1 [json_name = "id"];
}

message Update{{.Name}}Request {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined todo
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	// {{.Name}} defined todo
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}

message Update{{.Name}}Response {
	uint32 id = 1 [json_name = "id"];
}

message Query{{.Name}}Request {
	uint32 id = 1 [json_name = "id"];
	int64 page = 2 [json_name = "page"];
	int64 size = 3 [json_name = "size"];
	string order = 4 [json_name = "order"];
}

message Query{{.Name}}ResponseItem {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined todo
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	// {{.Name}} defined todo
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}

message Query{{.Name}}Response {
    repeated Query{{.Name}}ResponseItem data = 1 [json_name = "data"];
	int64 page = 2 [json_name = "page"];
	int64 size = 3 [json_name = "size"];
	int64 total_count = 4 [json_name = "total_count"];
}

message Query{{.Name}}DetailRequest {
	uint32 id = 1 [json_name = "id"];
}

message Query{{.Name}}DetailResponse {
{{- range .Fields}}
{{- if .IsRepeated}}
	// {{.Name}} defined todo
	repeated {{.TypeName}} {{.Name}} = {{.Order}};
{{- else}}
	// {{.Name}} defined todo
	{{.TypeName}} {{.Name}} = {{.Order}} [json_name = "{{.Name}}"];
{{- end}}
{{- end}}
}
`
	tmpl, err := template.New("handler").Parse(msgTemplate)
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

func tables2Handler(t *schemas.Table) *message {
	msg := &message{
		Name:   case2camel(t.Name),
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
