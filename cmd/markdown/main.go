package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	_ "github.com/gomarkdown/markdown"
	"github.com/pkg/errors"
)

func main() {

	var (
		name         = os.Args[1]
		content, err = ioutil.ReadFile(name)
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse tmpl"))
	}

	html := markdown.ToHTML(content, nil, nil)
	fmt.Print(string(html))
}
