package main

import (
	"github.com/DavudSafarli/encrypt/logic"
)

type Printer interface {
	Print(output string)
}

type App struct {
	printer Printer
}

func NewApp(printer Printer) App {
	return App{
		printer: printer,
	}
}

func (app App) Run(input string) {
	output := logic.Rot13(input)
	app.printer.Print(output)
}
