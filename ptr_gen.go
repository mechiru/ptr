// +build ignore

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	if err := inner(); err != nil {
		log.Fatal(err)
	}
}

func inner() (err error) {
	file, err := os.OpenFile("ptr.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New("code").
		Funcs(template.FuncMap{"title": strings.Title}).
		Parse(code)
	if err != nil {
		return err
	}
	return tmpl.Execute(file, basicTypes)
}

// https://tour.golang.org/basics/11
// https://github.com/golang/go/blob/7e01b3b3879593828b89f4ff4a04667a547b22d9/src/go/types/type.go#L26-L42
var basicTypes = []string{
	"bool",
	"string",
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	"byte",
	"rune",
	"float32", "float64",
	"complex64", "complex128",
}

const code = `
// Code generated by main.go. DO NOT EDIT.

package ptr

{{ range . }}
func Deref{{ . | title }}(a *{{ . }}, b {{ . }}) {{ . }} {
	if a != nil {
		return *a
	} else {
		return b
	}
}

func Ref{{ . | title }}(a {{ . }}) *{{ . }} { return &a }
{{ end }}
`