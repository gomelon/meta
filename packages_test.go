package meta

import (
	"encoding/json"
	"os"
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
				loadPackagePaths: []string{"github.com/gomelon/meta/internal/testdata"},
				findPackagePath:  "github.com/gomelon/meta/internal/testdata",
				findTypeName:     "IntVar",
			},
			wantFound: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workdir, _ := os.Getwd()
			packages := NewPackages(workdir)
			err := packages.Load(tt.args.loadPackagePaths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := packages.FindType(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("Load() then FindType(%v,%v), want %v",
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
				loadPackagePaths: []string{"github.com/gomelon/meta/internal/testdata"},
				findPackagePath:  "time",
				findTypeName:     "Time",
			},
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workdir, _ := os.Getwd()
			packages := NewPackages(workdir)
			_ = packages.Load(tt.args.loadPackagePaths...)
			got := packages.FindType(tt.args.findPackagePath, tt.args.findTypeName)
			if (got != nil) != tt.wantFound {
				t.Errorf("FindType(%v,%v) wantFound %v",
					tt.args.findPackagePath, tt.args.findTypeName, tt.wantFound)
			}
		})
	}
}

type inputElement struct {
	PkgPath string
	Name    string
	Parent  string
	Metas   map[string][]Meta
}

func TestPackages_FindByMeta(t *testing.T) {
	type args struct {
		loadPackagePaths []string
		metas            []Meta
	}
	tests := []struct {
		name              string
		args              args
		wantInputElements []*inputElement
		wantErr           bool
	}{
		{
			name: "",
			args: args{
				loadPackagePaths: []string{"/home/kimloong/GolandProjects/gomelon/meta/internal/testdata"},
				metas:            []Meta{&Table{}},
			},
			wantInputElements: []*inputElement{
				{
					PkgPath: "github.com/gomelon/meta/internal/testdata",
					Name:    "UserDao",
					Metas:   map[string][]Meta{MetaSqlTable: {&Table{Value: "user"}}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workdir, _ := os.Getwd()
			packages := NewPackages(workdir)
			_ = packages.Load(tt.args.loadPackagePaths...)
			gotInputs, err := packages.FindByMeta(tt.args.metas...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			actualInputElements := extractInputElements(gotInputs)
			isInputElementsEquals(t, tt.wantInputElements, actualInputElements)
		})
	}
}

func extractInputElements(inputs []*Input) []*inputElement {
	inputElements := make([]*inputElement, 0, len(inputs))
	for _, input := range inputs {
		interfaces := input.Interfaces
		for _, i := range interfaces {
			inputElements = append(inputElements, &inputElement{
				PkgPath: input.PkgPath(),
				Name:    i.Name(),
				Metas:   i.Metas,
			})
		}
	}
	return inputElements
}

func isInputElementsEquals(t *testing.T, want []*inputElement, actual []*inputElement) {
	if len(want) != len(actual) {
		marshalWant, _ := json.Marshal(want)
		marshalActual, _ := json.Marshal(actual)
		t.Errorf("want input element lenth=%v, actual lenth=%v,want=%v,actual=%v",
			len(want), len(actual), string(marshalWant), string(marshalActual))
	}
	for _, wantElement := range want {
		found := false
		for _, actualElement := range actual {
			if wantElement.PkgPath == actualElement.PkgPath &&
				wantElement.Name == actualElement.Name &&
				wantElement.Parent == actualElement.Parent &&
				len(wantElement.Metas) == len(actualElement.Metas) {
				found = true
				break
			}
		}
		if !found {
			marshalWantElement, _ := json.Marshal(wantElement)
			marshalActual, _ := json.Marshal(actual)
			t.Errorf("want input element but not found \nwantElement=%v, \nactual=%v",
				string(marshalWantElement), string(marshalActual))
		}
	}
}

func TestName(t *testing.T) {
	println(TargetUnsupported)
}
