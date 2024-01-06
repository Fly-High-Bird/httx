package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/fly-high-bird/httx"
)

var (
	htmx = flag.Bool("htmx", false, "Should Hx-Location header be included")
)

func main() {
	flag.Parse()

	var (
		args   = flag.Args()
		ctx    = httx.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if len(args) == 0 {
		args = []string{"/"}
	}

	if *htmx {
		ctx.Headers["Hx-Location"] = args[0]
	}

	stdout.Encode(&httx.Response{
		Headers:  ctx.Headers,
		Cookies:  ctx.Cookies,
		Redirect: args[0],
	})
}
