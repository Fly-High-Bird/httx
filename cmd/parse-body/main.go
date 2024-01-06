package main

import (
	"fmt"

	"github.com/fly-high-bird/httx"
)

func main() {
	ctx := httx.NewContext()
	ctx.ParseStdin()
	fmt.Print(string(ctx.Body))
}
