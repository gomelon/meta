package meta

import (
	"fmt"
	"os"
	"testing"
)

var tplText = `
{{range $itf := interfaces|filterByMeta "sql:table"}}

	{{ $decorator := print $itf.Name "Impl" }}
	type {{$decorator}} struct{
	}

	{{range $method := $itf|methods}}
	func (_d *{{$decorator}}) {{$method|declare}}{
		panic("implement me")
	}
	{{end}}
{{end}}

{{range $struct := structs|filterByMeta "aop:iface"}}
    {{$decorator := print $struct.Name "AOPIface"}}
    type {{$decorator}} interface {
    {{range $method := $struct|methods}}
        {{if $method|exported}}
            {{$method|declare}}
        {{end}}
    {{end}}
    }
{{end}}
`

func TestTmplGenerate(t *testing.T) {
	//generate output file is in ./testdata/zz_testdata_gen.go
	workdir, _ := os.Getwd()
	path := workdir + "/testdata"
	generator, err := NewTmplPkgGen(path, tplText)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//err = generator.Print()
	err = generator.Generate()
	//err = generator.Generate()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestScanToTmplGenerate(t *testing.T) {
	//generate output file is in ./testdata/zz_testdata_gen.go
	err := ScanCurrentMod().
		TemplateText(tplText).
		RegexOr("testdata").
		And().
		Generate()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
