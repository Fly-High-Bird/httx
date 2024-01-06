package main

import (
	"bytes"
	"fmt"
	"html/template"
	"kokodo"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	var (
		ctx       kokodo.Context
		b         bytes.Buffer
		name      = os.Args[1]
		tmpl, err = template.ParseFiles("_templates/" + name + ".html")
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse tmpl"))
	}

	tmpl.ParseGlob("_templates/**/*.html")
	ctx = kokodo.LoadContext()

	if err = tmpl.Execute(&b, ctx.Props); err != nil {
		log.Fatal(errors.Wrap(err, "failed to render"))
	}

	fmt.Print(b.String())
}
