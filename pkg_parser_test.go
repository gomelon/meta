package meta

import (
	"testing"
)

func TestLoad(t *testing.T) {
	type args struct {
		loadPackagePaths []string
		findPackagePath  string
		findTypeName     string
	}
	tests := []struct {
		name      string
		args      args
		wantFound bool
		wantErr   bool
	}{
		{
			name: "Load",
			args: args{
				loadPackagePaths: []string{"github.com/gomelon/meta/testdata"},
				findPackagePath:  "github.com/gomelon/meta/testdata",
				findTypeName:     "IntVar",
			},
			wantFound: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkgParser := NewPkgParser()
			err := pkgParser.Load(tt.args.loadPackagePaths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := pkgParser.Object(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("Load() then Object(%v,%v), want %v",
					tt.args.findPackagePath, tt.args.findTypeName, tt.wantFound)
			}
		})
	}
}

func TestPackagesHelper_FindType(t *testing.T) {
	type args struct {
		loadPackagePaths []string
		findPackagePath  string
		findTypeName     string
	}
	tests := []struct {
		name      string
		args      args
		wantFound bool
	}{
		{
			name: "Should Find Import type",
			args: args{
				loadPackagePaths: []string{"github.com/gomelon/meta/testdata"},
				findPackagePath:  "time",
				findTypeName:     "Time",
			},
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkgParser := NewPkgParser()
			_ = pkgParser.Load(tt.args.loadPackagePaths...)
			got := pkgParser.Object(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("Object(%v,%v) wantFound %v",
					tt.args.findPackagePath, tt.args.findTypeName, tt.wantFound)
			}
		})
	}
}

func TestPkgParser_AssignableTo(t *testing.T) {
	type args struct {
		vPkgPath, vName string
		tPkgPath, tName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "struct implement interface",
			args: args{
				vPkgPath: "github.com/gomelon/meta/testdata",
				vName:    "SimpleStruct",
				tPkgPath: "github.com/gomelon/meta/testdata",
				tName:    "SimpleInterface",
			},
			want: true,
		},
		{
			name: "struct not implement interface",
			args: args{
				vPkgPath: "github.com/gomelon/meta/testdata",
				vName:    "SimpleStruct",
				tPkgPath: "github.com/gomelon/meta/testdata",
				tName:    "UserDao",
			},
			want: false,
		},
		{
			name: "same interface",
			args: args{
				vPkgPath: "github.com/gomelon/meta/testdata",
				vName:    "SimpleInterface",
				tPkgPath: "github.com/gomelon/meta/testdata",
				tName:    "SimpleInterface",
			},
			want: true,
		},
		{
			name: "context to context",
			args: args{
				vPkgPath: "context",
				vName:    "Context",
				tPkgPath: "context",
				tName:    "Context",
			},
			want: true,
		},
		{
			name: "var context to context",
			args: args{
				vPkgPath: "github.com/gomelon/meta/testdata",
				vName:    "varCtx",
				tPkgPath: "context",
				tName:    "Context",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkgParser := NewPkgParser()
			vType := pkgParser.Object(tt.args.vPkgPath, tt.args.vName).Type()
			tType := pkgParser.Object(tt.args.tPkgPath, tt.args.tName).Type()
			if got := pkgParser.AssignableTo(vType, tType); got != tt.want {
				t.Errorf("AssignableTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
