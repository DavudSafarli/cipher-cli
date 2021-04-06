package logic

import (
	"strings"
)

var rune_a int = int(rune('a'))
var rune_A int = int(rune('A'))
var rune_n int = int(rune('n'))
var rune_N int = int(rune('N'))
var rune_z int = int(rune('z'))
var rune_Z int = int(rune('Z'))

func Rot13(input string) string {
	sb := strings.Builder{}

	for _, ch := range input {
		ascii := int(ch)
		if (rune_a <= ascii && ascii < rune_n) || (rune_A <= ascii && ascii < rune_N) {
			_, err := sb.WriteRune(rune(ascii + 13))
			panicOnError(err)
			continue
		}
		if (rune_n <= ascii && ascii >= rune_z) || (rune_N <= ascii && ascii > rune_Z) {
			_, err := sb.WriteRune(rune(ascii - 13))
			panicOnError(err)
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
