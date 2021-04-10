package printer

import (
	"io"
)

type Printer struct {
	writer io.Writer
}

func NewPrinter(writer io.Writer) Printer {
	return Printer{
		writer: writer,
	}
}

func (p Printer) Print(output string) {
	p.writer.Write([]byte(output))
}
