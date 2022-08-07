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

{{range $struct := structs|filterByMeta "aop:interface"}}
    {{$decorator := print $struct.Name "AOPInterface"}}
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

	workdir, _ := os.Getwd()

	err := ScanFor(workdir).
		TemplateText(tplText).RegexOr("testdata").
		And().
		Generate()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
