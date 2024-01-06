package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"kokodo"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	var (
		ctx    kokodo.Context
		b      bytes.Buffer
		name   = os.Args[1]
		stdout = json.NewEncoder(os.Stdout)
		funcs  = map[string]any{
			"raw": func(s string) template.HTML { return template.HTML(s) },
		}
		parts     = strings.Split(name, "/")
		tmpl, err = template.
			// Removing pages/
			New(parts[len(parts)-1] + ".html").
			Funcs(funcs).
			ParseFiles("_templates/" + name + ".html")
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse tmpl"))
	}

	ctx = kokodo.LoadContext()

	tmpl.ParseGlob("_templates/**/*.html")

	if err = tmpl.Execute(&b, ctx.Props); err != nil {
		log.Fatal(errors.Wrap(err, "failed to render"))
	}

	stdout.SetEscapeHTML(false)
	stdout.Encode(&kokodo.Response{
		Headers: ctx.Headers,
		Cookies: ctx.Cookies,
		Body:    b.String(),
	})
}
