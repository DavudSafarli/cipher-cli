package logic

import "testing"

type test struct {
	name     string
	input    string
	expected string
}

func TestRot13(t *testing.T) {
	tests := []test{
		{name: "Test with ascii charset only", input: "qwertyuiopasdfghjklzxcvbnm", expected: "djreglhvbcnfqstuwxymkpioaz"},
		{name: "Test with non-ascii charset", input: "əüöğ汉字仮名!;%:?*()_'ş", expected: "əüöğ汉字仮名!;%:?*()_'ş"},
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
