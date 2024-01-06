package main

import (
	"fmt"
	"kokodo"
)

func main() {
	ctx := kokodo.NewContext()
	ctx.ParseStdin()
	fmt.Print(string(ctx.Body))
}
