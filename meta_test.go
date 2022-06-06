package meta

import (
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"testing"
)

func TestObjectTarget(t *testing.T) {
	type args struct {
		objectName string
	}
	tests := []struct {
		name string
		args args
		want Target
	}{
		{
			name: "Const",
			args: args{
				objectName: "IntConst",
			},
			want: TargetConst,
		},
		{
			name: "Var",
			args: args{
				objectName: "IntVar",
			},
			want: TargetVar,
		},
		{
			name: "Func",
			args: args{
				objectName: "SimpleFunc",
			},
			want: TargetFunc,
		},
		{
			name: "Struct",
			args: args{
				objectName: "SimpleStruct",
			},
			want: TargetStruct,
		},
		{
			name: "Interface",
			args: args{
				objectName: "SimpleInterface",
			},
			want: TargetInterface,
		},
	}
	for _, tt := range tests {
		scope := getScope(t)
		object := scope.Lookup(tt.args.objectName)
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectTarget(object); got != tt.want {
				t.Errorf("ObjectTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectTarget_StructField(t *testing.T) {
	type args struct {
		objectName string
	}
	tests := []struct {
		name string
		args args
		want Target
	}{
		{
			name: "Member",
			args: args{
				objectName: "SimpleStruct",
			},
			want: TargetField,
		},
	}
	for _, tt := range tests {
		scope := getScope(t)
		object := scope.Lookup(tt.args.objectName)
		structObject := object.Type().Underlying().(*types.Struct)
		fieldObject := structObject.Field(0)
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectTarget(fieldObject); got != tt.want {
				t.Errorf("ObjectTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectTarget_StructMethod(t *testing.T) {
	type args struct {
		objectName string
		methodName string
	}
	tests := []struct {
		name string
		args args
		want Target
	}{
		{
			name: "PointerStructMethod",
			args: args{
				objectName: "PointerSimpleStruct",
				methodName: "StructMethod",
			},
			want: TargetStructMethod,
		},
		{
			name: "StructMethod",
			args: args{
				objectName: "SimpleStruct",
				methodName: "StructMethod",
			},
			want: TargetStructMethod,
		},
	}
	for _, tt := range tests {
		scope := getScope(t)
		object := scope.Lookup(tt.args.objectName)
		namedObject := object.Type().(*types.Named)
		var methodObject *types.Func
		for i := 0; i < namedObject.NumMethods(); i++ {
			if namedObject.Method(i).Name() == tt.args.methodName {
				methodObject = namedObject.Method(0)
				break
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectTarget(methodObject); got != tt.want {
				t.Errorf("ObjectTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectTarget_InterfaceMethod(t *testing.T) {
	type args struct {
		objectName string
		methodName string
	}
	tests := []struct {
		name string
		args args
		want Target
	}{
		{
			name: "InterfaceMethod",
			args: args{
				objectName: "SimpleInterface",
				methodName: "InterfaceMethod",
			},
			want: TargetInterfaceMethod,
		},
	}
	for _, tt := range tests {
		scope := getScope(t)
		object := scope.Lookup(tt.args.objectName)
		interfaceMethod := object.Type().Underlying().(*types.Interface)
		var methodObject *types.Func
		for i := 0; i < interfaceMethod.NumMethods(); i++ {
			if interfaceMethod.Method(i).Name() == tt.args.methodName {
				methodObject = interfaceMethod.Method(0)
				break
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectTarget(methodObject); got != tt.want {
				t.Errorf("ObjectTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getScope(t *testing.T) *types.Scope {
	workdir, _ := os.Getwd()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Tests: false,
	}

	packageList, err := packages.Load(cfg, workdir+"/internal/testdata")
	if err != nil {
		t.Errorf("ObjectTarget() Load package fail, err=%v", err)
	}

	scope := packageList[0].Types.Scope()
	return scope
}

const MetaSqlTable = "sql:table"
const MetaSqlQuery = "sql:query"

type Table struct {
	Value string
}

func (t *Table) Target() Target {
	return TargetInterface
}

func (t *Table) Name() string {
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

func (q *Query) Target() Target {
	return TargetInterfaceMethod
}

func (q *Query) Name() string {
	return MetaSqlQuery
}

func (q *Query) Repeatable() bool {
	return false
}
