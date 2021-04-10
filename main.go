package main

import (
	"os"
	"strings"

	"github.com/DavudSafarli/encrypt/printer"
)

func main() {
	p := printer.NewPrinter(os.Stdout)
	app := NewApp(p)
	input := os.Args[1:]
	app.Run(strings.Join(input, " "))
}
