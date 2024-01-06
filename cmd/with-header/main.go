package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fly-high-bird/httx"
)

func main() {
	var (
		ctx    = httx.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if key := os.Args[1]; key != "" {
		ctx.Headers[key] = os.Args[2]
	}

	if err := stdout.Encode(&ctx); err != nil {
		log.Fatal(err)
	}
}
