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
			packageParser := NewPackageParser()
			err := packageParser.Load(tt.args.loadPackagePaths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := packageParser.ObjectByPkgPathAndName(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("Load() then ObjectByPkgPathAndName(%v,%v), want %v",
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
			packageParser := NewPackageParser()
			_ = packageParser.Load(tt.args.loadPackagePaths...)
			got := packageParser.ObjectByPkgPathAndName(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("ObjectByPkgPathAndName(%v,%v) wantFound %v",
					tt.args.findPackagePath, tt.args.findTypeName, tt.wantFound)
			}
		})
	}
}
