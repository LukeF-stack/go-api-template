package model_types

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Column   string
	DataType string
	FK       FK
}

type FK struct {
	Table  string
	Column string
}
