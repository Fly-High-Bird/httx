package main

import (
	"encoding/json"
	"kokodo"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		ctx    = kokodo.LoadContext()
		stdout = json.NewEncoder(os.Stdout)
	)

	if key := os.Args[1]; key != "" {
		val := os.Args[2]
		ctx.Cookies = append(ctx.Cookies, &http.Cookie{
			Name:     key,
			Value:    val,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
	}

	if err := stdout.Encode(&ctx); err != nil {
		log.Fatal(err)
	}
}
