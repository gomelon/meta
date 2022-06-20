package meta

import (
	"fmt"
	"github.com/antonmedv/expr"
	"os"
	"testing"
)

var tplText = `
{{range $itf := interfaces|filterByMeta "sql:table"}}

	{{ $decorator := print $itf.Name "Impl" }}
	type {{$decorator}} struct{
	}

	{{range $method := $itf|interfaceMethods}}
	func (_d *{{$decorator}}) {{$method|declare}}{
	}
	{{end}}
{{end}}
`

func TestTemplateGen(t *testing.T) {

	workdir, _ := os.Getwd()
	path := workdir + "/testdata"
	metas := []Meta{&Table{}}
	generator, err := NewTemplateGenerator(path, tplText, WithMetas(metas))
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

func TestName(t *testing.T) {
	program, _ := expr.Compile("1==1")
	output, _ := expr.Run(program, map[string]any{})
	fmt.Println(output == true)
}
