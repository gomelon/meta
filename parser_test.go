package meta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParser_ObjectMetaGroups(t *testing.T) {
	type args struct {
		pkgPath    string
		objectName string
		metaNames  []string
	}
	tests := []struct {
		name                 string
		args                 args
		wantParsedMetaGroups map[string]Group
	}{
		{
			name: "Name",
			args: args{
				pkgPath:    "github.com/gomelon/meta/testdata",
				objectName: "SimpleFunc",
				metaNames:  []string{"testdata.SomeMeta"},
			},
			wantParsedMetaGroups: map[string]Group{
				"testdata.SomeMeta": []*Meta{
					{
						qualifyName: "testdata.SomeMeta",
						properties:  map[string]any{"some": int64(1)},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkgParser := NewPkgParser()
			err := pkgParser.Load(tt.args.pkgPath)
			if err != nil {
				t.Errorf("ObjectMetaGroups() load pkg fail")
			}
			object := pkgParser.Object(tt.args.pkgPath, tt.args.objectName)
			fmt.Println(object.Name(), object.Pos())

			objects := pkgParser.Functions(tt.args.pkgPath)
			for _, o := range objects {
				fmt.Println(o.Name(), o.Pos())
			}
			parser := NewParser(pkgParser)
			gotParsedMetaGroups := parser.ObjectMetaGroups(object, tt.args.metaNames...)
			if !reflect.DeepEqual(gotParsedMetaGroups, tt.wantParsedMetaGroups) {
				t.Errorf("ObjectMetaGroups() = %v, want %v", gotParsedMetaGroups, tt.wantParsedMetaGroups)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		qualifyName string
		comment     string
	}
	tests := []struct {
		name  string
		args  args
		want  *Meta
		want1 bool
	}{
		{
			name: "normal",
			args: args{
				qualifyName: "demo.Demo",
				comment:     "+demo.Demo Int32Value=1 StringValue=\"Hi\" BoolValue=true",
			},
			want: &Meta{
				qualifyName: "demo.Demo",
				properties: map[string]any{
					"Int32Value":  int64(1),
					"StringValue": "Hi",
					"BoolValue":   true,
				},
			},
			want1: true,
		},
		{
			name: "empty string value",
			args: args{
				qualifyName: "demo.Demo",
				comment:     "+demo.Demo Int32Value=1 StringValue=\"\" BoolValue=true",
			},
			want: &Meta{
				qualifyName: "demo.Demo",
				properties: map[string]any{
					"Int32Value":  int64(1),
					"StringValue": "",
					"BoolValue":   true,
				},
			},
			want1: true,
		},
		{
			name: "not spec meta 1",
			args: args{
				qualifyName: "demo.Dem",
				comment:     "+demo.Demo Int32Value=1 StringValue=\"Hi\" BoolValue=true",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "not spec meta 2",
			args: args{
				qualifyName: "demo.Demo1",
				comment:     "+demo.Demo Int32Value=1 StringValue=\"Hi\" BoolValue=true",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "not spec meta 2",
			args: args{
				qualifyName: "demo.Demo1",
				comment:     "+demo.Demo Int32Value=1 StringValue=\"Hi\" BoolValue=true",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "invalid string value",
			args: args{
				qualifyName: "demo.Demo",
				comment:     "+demo.Demo Int32Value=1 StringValue=Hi\" BoolValue=true",
			},
			want: &Meta{
				qualifyName: "demo.Demo",
				properties: map[string]any{
					"Int32Value": int64(1),
					"Hi":         true,
					"BoolValue":  true,
				},
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.qualifyName, tt.args.comment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
