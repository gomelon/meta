package meta

import (
	"fmt"
	"os"
	"strings"
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
	pkgPath := "github.com/gomelon/meta/internal/testdata"
	metas := []Meta{&Table{}}
	funcMap := map[string]any{
		"short": func(name string) string {
			return strings.ToLower(string(name[0]))
		},
	}
	generator, err := NewTemplateGenerator(workdir, pkgPath, metas, tplText, funcMap)
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
