package meta

import (
	"go/types"
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
		want Place
	}{
		{
			name: "Const",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "IntConst",
			},
			want: PlaceConst,
		},
		{
			name: "Var",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "IntVar",
			},
			want: PlaceVar,
		},
		{
			name: "Func",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "SimpleFunc",
			},
			want: PlaceFunc,
		},
		{
			name: "Struct",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "SimpleStruct",
			},
			want: PlaceStruct,
		},
		{
			name: "Interface",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "SimpleInterface",
			},
			want: PlaceInterface,
		},
	}
	for _, tt := range tests {
		pkgParser := NewPkgParser()
		_ = pkgParser.Load(tt.args.pkgPath)
		object := pkgParser.Object(tt.args.pkgPath, tt.args.objectName)
		t.Run(tt.name, func(t *testing.T) {
			if got := pkgParser.ObjectPlace(object); got != tt.want {
				t.Errorf("ObjectPlace() = %v, want %v", got, tt.want)
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
		want Place
	}{
		{
			name: "Field",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "SimpleStruct",
			},
			want: PlaceField,
		},
	}
	for _, tt := range tests {
		pkgParser := NewPkgParser()
		_ = pkgParser.Load(tt.args.pkgPath)
		object := pkgParser.Object(tt.args.pkgPath, tt.args.objectName)
		structObject := object.Type().Underlying().(*types.Struct)
		fieldObject := structObject.Field(0)
		t.Run(tt.name, func(t *testing.T) {
			if got := pkgParser.ObjectPlace(fieldObject); got != tt.want {
				t.Errorf("ObjectPlace() = %v, want %v", got, tt.want)
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
	wantMethodType := PlaceStructMethod
	wantParamType := PlaceFuncVar
	wantResultType := PlaceFuncVar
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PointerMethod",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleStruct",
				methodName:  "PointerMethod",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleStruct",
				methodName:  "Method",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "MethodWithParamAndResult",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleStruct",
				methodName:  "MethodWithParamAndResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
		{
			name: "MethodWithParamAndNameResult",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleStruct",
				methodName:  "MethodWithParamAndNameResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		pkgParser := NewPkgParser()
		_ = pkgParser.Load(tt.args.pkgPath)
		object := pkgParser.Object(tt.args.pkgPath, tt.args.objectName)
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
			if got := pkgParser.ObjectPlace(methodObject); got != wantMethodType {
				t.Errorf("ObjectPlace(%v) = %v, wantMethodType %v", tt.args.methodName, got, wantMethodType)
			}
			if tt.args.paramIndex >= 0 {
				if got := pkgParser.ObjectPlace(paramObject); got != wantParamType {
					t.Errorf("ObjectPlace(%v.params[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.paramIndex, got, wantParamType)
				}
			}
			if tt.args.resultIndex >= 0 {
				if got := pkgParser.ObjectPlace(resultObject); got != wantResultType {
					t.Errorf("ObjectPlace(%v.results[%v]) = %v, wantMethodType %v",
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
	wantMethodType := PlaceInterfaceMethod
	wantParamType := PlaceFuncVar
	wantResultType := PlaceFuncVar
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleInterface",
				methodName:  "Method",
				paramIndex:  -1,
				resultIndex: -1,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleInterface",
				methodName:  "MethodWithParamAndResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
		{
			name: "Method",
			args: args{
				pkgPath:     "github.com/gomelon/meta/testdata",
				objectName:  "SimpleInterface",
				methodName:  "MethodWithParamAndNameResult",
				paramIndex:  0,
				resultIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		pkgParser := NewPkgParser()
		_ = pkgParser.Load(tt.args.pkgPath)
		object := pkgParser.Object(tt.args.pkgPath, tt.args.objectName)
		iface := object.Type().Underlying().(*types.Interface)
		var methodObject *types.Func
		var paramObject *types.Var
		var resultObject *types.Var
		for i := 0; i < iface.NumMethods(); i++ {
			methodObject = iface.Method(i)
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
			if got := pkgParser.ObjectPlace(methodObject); got != wantMethodType {
				t.Errorf("ObjectPlace() = %v, want %v", got, wantMethodType)
			}
			if tt.args.paramIndex >= 0 {
				if got := pkgParser.ObjectPlace(paramObject); got != wantParamType {
					t.Errorf("ObjectPlace(%v.params[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.paramIndex, got, wantParamType)
				}
			}
			if tt.args.resultIndex >= 0 {
				if got := pkgParser.ObjectPlace(resultObject); got != wantResultType {
					t.Errorf("ObjectPlace(%v.results[%v]) = %v, wantMethodType %v",
						tt.args.methodName, tt.args.resultIndex, got, wantResultType)
				}
			}
		})
	}
}
