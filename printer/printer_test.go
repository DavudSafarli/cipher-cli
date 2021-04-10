package printer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConsole(t *testing.T) {
	t.Run(`It writes when providing an input`, func(t *testing.T) {
		// arrange
		reader, writer, err := os.Pipe()
		require.Nil(t, err)

		// act
		console := NewPrinter(writer)
		input := "text"
		inputSize := len(input)
		console.Print(input)

		// assert
		buf := make([]byte, inputSize*2)
		n, err := reader.Read(buf)

		require.Nil(t, err)
		require.Equal(t, inputSize, n, "reader read more bytes than input")
		output := string(buf[0:inputSize])
		require.Equal(t, input, output)
	})
}
