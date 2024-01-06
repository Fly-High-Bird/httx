package main

import (
	"bytes"
	"encoding/json"
	"kokodo"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	var (
		buf    bytes.Buffer
		ctx    = kokodo.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if err := json.NewEncoder(&buf).Encode(ctx.Props); err != nil {
		log.Fatal(errors.Wrap(err, "failed to render"))
	}

	ctx.Headers["Content-Type"] = "application/json"
	stdout.SetEscapeHTML(false)
	stdout.Encode(&kokodo.Response{
		Headers: ctx.Headers,
		Body:    buf.String(),
	})
}
