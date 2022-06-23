package meta

const MetaSqlTable = "sql:table"
const MetaSqlQuery = "sql:query"

type Table struct {
	Value string
}

func (t *Table) PlaceAt() Place {
	return PlaceInterface
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

func (q *Query) Target() Place {
	return PlaceInterfaceMethod
}

func (q *Query) Name() string {
	return MetaSqlQuery
}

func (q *Query) Repeatable() bool {
	return false
}
