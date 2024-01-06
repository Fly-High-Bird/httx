package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fly-high-bird/httx"
	"github.com/joho/godotenv"
)

var path, addr string

func main() {
	flag.StringVar(&path, "path", ".", "path to root dir")
	flag.StringVar(&addr, "http", ":8080", "http address")

	flag.Parse()

	log.Println("Loading any environment variables")
	env, err := godotenv.Read()
	if err != nil {
		log.Println("No .env file found locally")
	}

	log.Printf("Mounting to directory: %s", path)
	h := httx.Mount(path, env)
	if err := h.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening at http://%s", addr)
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}
