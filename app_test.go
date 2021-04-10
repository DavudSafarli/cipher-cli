package main

import (
	"testing"

	"github.com/DavudSafarli/encrypt/logic"
	"github.com/DavudSafarli/encrypt/printer"
	"github.com/stretchr/testify/require"
)

type fakeWriter struct {
	lastOutput string
}

func (c *fakeWriter) Write(p []byte) (n int, err error) {
	c.lastOutput = string(p)
	return len(p), nil
}

func TestApp(t *testing.T) {
	t.Run(`Reads input and outputs ROT-13 transformed string`, func(t *testing.T) {
		// arrange
		input := "Hello World!"
		writer := &fakeWriter{}
		console := printer.NewPrinter(writer)

		// act
		app := NewApp(console)
		app.Run(input)

		// assert
		require.Equal(t, logic.Rot13(input), writer.lastOutput)
	})
}
