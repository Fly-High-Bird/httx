package main

import (
	"encoding/json"
	"kokodo"
	"log"
	"os"
)

func main() {
	var (
		ctx    = kokodo.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if key := os.Args[1]; key != "" {
		ctx.Headers[key] = os.Args[2]
	}

	if err := stdout.Encode(&ctx); err != nil {
		log.Fatal(err)
	}
}
