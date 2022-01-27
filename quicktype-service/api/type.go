package api

type type1 struct {
	Name      string
	TableName string
	Fields    []*field1
}

type field1 struct {
	Name       string
	TypeName   string
	JsonName   string
	ColumnName string
	Order      int
	IsRepeated bool
}
