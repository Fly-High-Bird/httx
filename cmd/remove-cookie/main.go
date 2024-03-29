package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fly-high-bird/httx"
)

func main() {
	var (
		ctx    = httx.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if key := os.Args[1]; key != "" {
		ctx.Cookies = append(ctx.Cookies, &http.Cookie{
			Name:     key,
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		})
	}

	if err := stdout.Encode(&ctx); err != nil {
		log.Fatal(err)
	}
}
