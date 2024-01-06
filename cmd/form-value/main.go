package main

import (
	"fmt"
	"os"

	"github.com/fly-high-bird/httx"
)

func main() {
	ctx := httx.NewContext()
	ctx.SetBody(os.Args[1])
	fmt.Print(ctx.FormValue(os.Args[2]))
}
