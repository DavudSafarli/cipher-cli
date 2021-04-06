package logic

import "testing"

func TestRot13(t *testing.T) {
	input := "qwertyuiopasdfghjklzxcvbnm"
	actual := Rot13(input)
	expected := "djreglhvbcnfqstuwxymkpioaz"

	if actual != expected {
		t.Error("not equal", actual, expected)
	}

}
