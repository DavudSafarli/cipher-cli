package main

import (
	"github.com/DavudSafarli/encrypt/logic"
	"github.com/DavudSafarli/encrypt/printer"
)

type App struct {
	printer printer.Printer
}

func NewApp(printer printer.Printer) App {
	return App{
		printer: printer,
	}
}

func (app App) Run(input string) {
	output := logic.Rot13(input)
	app.printer.Print(output)
}
