package testmeta

import (
	"github.com/gomelon/meta"
)

type Table struct {
	Value string
}

func (t *Table) Target() meta.Type {
	return meta.TypeInterface
}

func (t *Table) Name() string {
	return "sql:table"
}

type Query struct {
	Value     string
	Master    bool
	Omitempty bool
}

func (q *Query) Target() meta.Type {
	return meta.TypeInterfaceMethod
}

func (q *Query) Name() string {
	return "sql:query"
}
