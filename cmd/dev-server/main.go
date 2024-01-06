package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	path = flag.String("path", ".", "path to root dir")
	addr = flag.String("addr", ":8080", "http address")
)

func main() {
	flag.Parse()

	log.Println("Loading any environment variables")
	env, err := godotenv.Read()
	if err != nil {
		log.Println("No .env file found locally")
	}

	log.Printf("Mounting to directory: %s", *path)
	h := .Mount(*path, env)
	if err := h.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening at http://%s", *addr)
	if err := http.ListenAndServe(*addr, h); err != nil {
		log.Fatal(err)
	}
}
