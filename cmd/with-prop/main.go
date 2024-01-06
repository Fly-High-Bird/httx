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
		var v any
		if err := json.Unmarshal([]byte(os.Args[2]), &v); err != nil {
			v = os.Args[2]
		}
		ctx.Props[key] = v
	}

	if err := stdout.Encode(&ctx); err != nil {
		log.Fatal(err)
	}
}
