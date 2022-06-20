package meta

const MetaSqlTable = "sql:table"
const MetaSqlQuery = "sql:query"

type Table struct {
	Value string
}

func (t *Table) Target() Type {
	return TypeInterface
}

func (t *Table) Directive() string {
	return MetaSqlTable
}

func (t *Table) Repeatable() bool {
	return false
}

type Query struct {
	Value     string
	Master    bool
	Omitempty bool
}

func (q *Query) Target() Type {
	return TypeInterfaceMethod
}

func (q *Query) Name() string {
	return MetaSqlQuery
}

func (q *Query) Repeatable() bool {
	return false
}
