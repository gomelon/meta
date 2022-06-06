package meta

import (
	"go/types"
)

type Interface struct {
	*types.Interface
	types.Object
	Methods  []*Func
	Comments []string
	Metas    map[string][]Meta
}

type Struct struct {
	*types.Struct
	types.Object
	Methods  []*Func
	Fields   []*Field
	Comments []string
	Metas    map[string][]Meta
}

type Field struct {
	*types.Var
	types.Object
	Comments []string
	Metas    map[string][]Meta
}

// Func represents a function
type Func struct {
	*types.Func
	types.Object
	Params   ParamsSlice
	Results  ParamsSlice
	Comments []string
	Metas    map[string][]Meta

	ReturnsError   bool
	AcceptsContext bool
	hasPtrRecv     bool
}

// ParamsSlice slice of Param
type ParamsSlice []Param

// Param represents function argument or result
type Param struct {
	*types.Var
	types.Object
	Comments []string
	Metas    map[string][]Meta

	Variadic bool
}
