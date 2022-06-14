package meta

type Type uint32

const (
	TypeUnknown      = 0
	TypeConst   Type = 1 << iota
	TypeVar
	TypeInterface
	TypeStruct
	TypeField
	TypeInterfaceMethod
	TypeStructMethod
	TypeFunc
	TypeFuncVar //func/method param or result, because can't distinguish on named result
)
