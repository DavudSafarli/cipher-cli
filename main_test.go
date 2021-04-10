package main

import (
	"os"
	"strings"
	"testing"

	"github.com/DavudSafarli/encrypt/logic"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	t.Run(`Application reads command line arguments, and writes ROT13 transformation to stdout`, func(t *testing.T) {
		// arrange
		r, w, _ := os.Pipe()
		os.Stdout = w

		// act
		input := []string{"Hello", "World", "!"}
		os.Args = append([]string{"main.exe"}, input...)
		main()

		// assert
		inputAsStr := strings.Join(input, " ")
		inputLen := len(inputAsStr)
		expectedOutput := logic.Rot13(inputAsStr)

		buf := make([]byte, inputLen*2)
		n, err := r.Read(buf)
		require.Nil(t, err)
		require.Equal(t, inputLen, n)
		require.Equal(t, expectedOutput, string(buf[0:inputLen]))
	})
}
