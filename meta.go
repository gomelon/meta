package meta

import "go/types"

type Meta interface {
	Target() Target
	Name() string
	Repeatable() bool
}

type Target uint16

const (
	TargetUnsupported Target = 1 << iota
	TargetConst
	TargetVar
	TargetInterface
	TargetStruct
	TargetField
	TargetInterfaceMethod
	TargetStructMethod
	TargetFunc
)

type Metas []Meta

func (m Metas) Target() Target {
	var target Target
	for _, meta := range m {
		target = target | meta.Target()
	}
	return target
}

var objectToTarget = map[types.Object]Target{}

func ObjectTarget(object types.Object) Target {
	target, ok := objectToTarget[object]
	if ok {
		return target
	}
	switch object := object.(type) {
	case *types.Const:
		target = TargetConst
	case *types.Var:
		if object.IsField() {
			target = TargetField
		} else {
			target = TargetVar
		}
	case *types.Func:
		signature := object.Type().(*types.Signature)
		receiver := signature.Recv()
		if receiver == nil {
			target = TargetFunc
		} else {
			receiverPointer, ok := receiver.Type().(*types.Pointer)
			var receiverType types.Type
			if ok {
				receiverType = receiverPointer.Elem().Underlying()
			} else {
				receiverType = receiver.Type().Underlying()
			}
			switch receiverType.(type) {
			case *types.Struct:
				target = TargetStructMethod
			case *types.Interface:
				target = TargetInterfaceMethod
			default:
				target = TargetUnsupported
			}
		}
	case *types.TypeName:
		switch object.Type().Underlying().(type) {
		case *types.Interface:
			target = TargetInterface
		case *types.Struct:
			target = TargetStruct
		default:
			target = TargetUnsupported
		}
	default:
		target = TargetUnsupported
	}
	objectToTarget[object] = target
	return target
}
