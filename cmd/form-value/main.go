package main

import (
	"fmt"
	"kokodo"
	"os"
)

func main() {
	ctx := kokodo.NewContext()
	ctx.SetBody(os.Args[1])
	fmt.Print(ctx.FormValue(os.Args[2]))
}
