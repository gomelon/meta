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
				metaNames:  []string{"SimpleFunc"},
			},
			wantParsedMetaGroups: map[string]Group{
				"SimpleFunc": []*Meta{{name: "SimpleFunc", properties: map[string]string{"some": "1"}}},
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
