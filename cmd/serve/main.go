package main

import (
	"encoding/json"
	"os"

	"github.com/fly-high-bird/httx"
)

func main() {
	var (
		ctx    = httx.NewContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	stdout.SetEscapeHTML(false)
	stdout.Encode(&httx.Response{
		Headers: ctx.Headers,
		Body:    string(ctx.Body),
	})
}
