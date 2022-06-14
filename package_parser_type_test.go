package meta

import (
	"go/types"
	"os"
	"testing"
)

func TestObjectTarget(t *testing.T) {
	type args struct {
		pkgPath    string
		objectName string
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "Const",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "IntConst",
			},
			want: TypeConst,
		},
		{
			name: "Var",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "IntVar",
			},
			want: TypeVar,
		},
		{
			name: "Func",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "SimpleFunc",
			},
			want: TypeFunc,
		},
		{
			name: "Struct",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "SimpleStruct",
			},
			want: TypeStruct,
		},
		{
			name: "Interface",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "SimpleInterface",
			},
			want: TypeInterface,
		},
	}
	for _, tt := range tests {
		workdir, _ := os.Getwd()
		packageParser := NewPackageParser(workdir)
		_ = packageParser.Load(tt.args.pkgPath)
		object := packageParser.TypeByPkgPathAndName(tt.args.pkgPath, tt.args.objectName)
		t.Run(tt.name, func(t *testing.T) {
			if got := packageParser.ObjectType(object); got != tt.want {
				t.Errorf("ObjectType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectTarget_StructField(t *testing.T) {
	type args struct {
		pkgPath    string
		objectName string
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "Member",
			args: args{
				pkgPath:    "github.com/gomelon/meta/internal/testdata",
				objectName: "SimpleStruct",
			},
			want: TypeField,
		},
	}
	for _, tt := range tests {
		workdir, _ := os.Getwd()
		packageParser := NewPackageParser(workdir)
		_ = packageParser.Load(tt.args.pkgPath)
		object := packageParser.TypeByPkgPathAndName(tt.args.pkgPath, tt.args.objectName)
		structObject := object.Type().Underlying().(*types.Struct)
		fieldObject := structObject.Field(0)
		t.Run(tt.name, func(t *testing.T) {
			if got := packageParser.ObjectType(fieldObject); got != tt.want {
				t.Errorf("ObjectType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectTarget_StructMethod(t *testing.T) {
	type args struct {
		pkgPath     string
		objectName  string
		methodName  string
		paramIndex  int //-1 means no test param
		resultIndex int //-1 means no test result
	}
	wantMethodType := TypeStructMethod
	wantParamType := TypeFuncVar
	wantResultType := TypeFuncVar
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PointerMethod",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleStruct",
				methodName:  "PointerMethod",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleStruct",
				methodName:  "Method",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "MethodWithParamAndResult",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleStruct",
				methodName:  "MethodWithParamAndResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
		{
			name: "MethodWithParamAndNameResult",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleStruct",
				methodName:  "MethodWithParamAndNameResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		workdir, _ := os.Getwd()
		packageParser := NewPackageParser(workdir)
		_ = packageParser.Load(tt.args.pkgPath)
		object := packageParser.TypeByPkgPathAndName(tt.args.pkgPath, tt.args.objectName)
		namedObject := object.Type().(*types.Named)
		var methodObject *types.Func
		var paramObject *types.Var
		var resultObject *types.Var
		for i := 0; i < namedObject.NumMethods(); i++ {
			methodObject = namedObject.Method(i)
			if methodObject.Name() == tt.args.methodName {
				signature := methodObject.Type().(*types.Signature)
				if tt.args.paramIndex >= 0 {
					paramObject = signature.Params().At(tt.args.paramIndex)
				}
				if tt.args.resultIndex >= 0 {
					resultObject = signature.Results().At(tt.args.resultIndex)
				}
				break
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := packageParser.ObjectType(methodObject); got != wantMethodType {
				t.Errorf("ObjectType(%v) = %v, wantMethodType %v", tt.args.methodName, got, wantMethodType)
			}
			if tt.args.paramIndex >= 0 {
				if got := packageParser.ObjectType(paramObject); got != wantParamType {
					t.Errorf("ObjectType(%v.params[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.paramIndex, got, wantParamType)
				}
			}
			if tt.args.resultIndex >= 0 {
				if got := packageParser.ObjectType(resultObject); got != wantResultType {
					t.Errorf("ObjectType(%v.results[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.resultIndex, got, wantResultType)
				}
			}
		})
	}
}

func TestObjectTarget_InterfaceMethod(t *testing.T) {
	type args struct {
		pkgPath     string
		objectName  string
		methodName  string
		paramIndex  int //-1 means no test param
		resultIndex int //-1 means no test result
	}
	wantMethodType := TypeInterfaceMethod
	wantParamType := TypeFuncVar
	wantResultType := TypeFuncVar
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleInterface",
				methodName:  "Method",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleInterface",
				methodName:  "MethodWithParamAndResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/internal/testdata",
				objectName:  "SimpleInterface",
				methodName:  "MethodWithParamAndNameResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		workdir, _ := os.Getwd()
		packageParser := NewPackageParser(workdir)
		_ = packageParser.Load(tt.args.pkgPath)
		object := packageParser.TypeByPkgPathAndName(tt.args.pkgPath, tt.args.objectName)
		itf := object.Type().Underlying().(*types.Interface)
		var methodObject *types.Func
		var paramObject *types.Var
		var resultObject *types.Var
		for i := 0; i < itf.NumMethods(); i++ {
			methodObject = itf.Method(i)
			if methodObject.Name() == tt.args.methodName {
				signature := methodObject.Type().(*types.Signature)

				if tt.args.paramIndex >= 0 {
					paramObject = signature.Params().At(tt.args.paramIndex)
				}
				if tt.args.resultIndex >= 0 {
					resultObject = signature.Results().At(tt.args.resultIndex)
				}
				break
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := packageParser.ObjectType(methodObject); got != wantMethodType {
				t.Errorf("ObjectType() = %v, want %v", got, wantMethodType)
			}
			if tt.args.paramIndex >= 0 {
				if got := packageParser.ObjectType(paramObject); got != wantParamType {
					t.Errorf("ObjectType(%v.params[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.paramIndex, got, wantParamType)
				}
			}
			if tt.args.resultIndex >= 0 {
				if got := packageParser.ObjectType(resultObject); got != wantResultType {
					t.Errorf("ObjectType(%v.results[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.resultIndex, got, wantResultType)
				}
			}
		})
	}
}
