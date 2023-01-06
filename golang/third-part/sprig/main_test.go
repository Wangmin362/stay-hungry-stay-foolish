package sprig

import (
	"bytes"
	"fmt"
	sprigV3 "github.com/Masterminds/sprig/v3"
	"github.com/fatih/structs"
	"html/template"
	"testing"
)

func TestEmpty(t *testing.T) {
	p := map[string]interface{}{
		"Spec": map[string]interface{}{
			"A": map[string]interface{}{
				"B": 45,
			},
		},
	}

	tem := `sdfsfjkl {{if empty .Spec.A.B }}111{{else}}{{.Spec.A.B}}{{end}}  sdfsfsdf`
	tmpl, err := template.New("sdfsdf").Funcs(sprigV3.GenericFuncMap()).Parse(tem)
	if err != nil {
		t.Fatal(err)
	}
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, &p); err != nil {
		t.Fatal(err)
	}

	fmt.Println(buffer.String())

}

// 结构体必须转为Map许多判断才有用
func TestEmpty2(t *testing.T) {
	type Iner11 struct {
		D string
	}

	type Iner struct {
		C Iner11
	}

	type person struct {
		Spec Iner
		Id   string
	}
	p := person{Spec: Iner{C: Iner11{D: "SDSF"}}, Id: "45"}

	tem := `sdfsfjkl {{or .Spec.C.D  "0sdfsdf"}}  sdfsfsdf`
	tmpl, err := template.New("sdfsdf").Funcs(sprigV3.GenericFuncMap()).Parse(tem)
	if err != nil {
		t.Fatal(err)
	}
	buffer := new(bytes.Buffer)
	m := structs.Map(p)
	if err := tmpl.Execute(buffer, m); err != nil {
		t.Fatal(err)
	}

	fmt.Println(buffer.String())

}
