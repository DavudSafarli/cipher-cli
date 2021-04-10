package logic

import "testing"

type test struct {
	name     string
	input    string
	expected string
}

func TestRot13(t *testing.T) {
	tests := []test{
		{name: "Transforms ascii charset", input: "qwertyuiopasdfghjklzxcvbnm", expected: "djreglhvbcnfqstuwxymkpioaz"},
		{name: "Transforms Uppercase letters", input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", expected: "NOPQRSTUVWXYZABCDEFGHIJKLM"},
		{name: "Transforms Lowercase letters", input: "abcdefghijklmnopqrstuvwxyz", expected: "nopqrstuvwxyzabcdefghijklm"},
		{name: "Doesn't change non-ascii charset and symbols", input: "əüöğ汉字仮名!;%:?*()_'ş@#[]", expected: "əüöğ汉字仮名!;%:?*()_'ş@#[]"},
		{name: "Doesn't transform numbers", input: "1234567890", expected: "1234567890"},
		{name: "Doesn't transform emojis", input: "✅🚫🙋", expected: "✅🚫🙋"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Rot13(test.input)
			if actual != test.expected {
				t.Error("not equal", actual, test.expected)
			}
		})
	}
}
