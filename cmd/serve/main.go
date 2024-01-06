package main

import (
	"encoding/json"
	"kokodo"
	"os"
)

func main() {
	var (
		ctx    = kokodo.NewContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	stdout.SetEscapeHTML(false)
	stdout.Encode(&kokodo.Response{
		Headers: ctx.Headers,
		Body:    string(ctx.Body),
	})
}
