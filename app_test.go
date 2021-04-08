package main

import (
	"testing"

	"github.com/DavudSafarli/encrypt/logic"
	"github.com/stretchr/testify/require"
)

type fakePrinter struct {
	lastOutput string
}

func (p *fakePrinter) Print(output string) {
	p.lastOutput = output
}

func TestAe(t *testing.T) {
	fake := &fakePrinter{}
	app := NewApp(fake)
	app.Run("input")
	require.Equal(t, logic.Rot13("input"), fake.lastOutput)
}
