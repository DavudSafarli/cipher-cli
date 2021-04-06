package logic

import "testing"

func TestRot13(t *testing.T) {

	t.Run(`Test with ascii charset only`, func(t *testing.T) {
		input := "qwertyuiopasdfghjklzxcvbnm"
		actual := Rot13(input)
		expected := "djreglhvbcnfqstuwxymkpioaz"

		if actual != expected {
			t.Error("not equal", actual, expected)
		}
	})

	t.Run(`Test with non-ascii charset`, func(t *testing.T) {
		input := "əüöğ汉字仮名!;%:?*()_'ş"
		actual := Rot13(input)
		expec := "əüöğ汉字仮名!;%:?*()_'ş"

		if actual != expec {
			t.Error("not equal", actual, expec)
		}
	})

}
