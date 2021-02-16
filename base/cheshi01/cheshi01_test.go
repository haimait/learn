package cheshi01

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right!")
	}
}
func TestAdd2(t *testing.T) {
	if Add(2, 3) != 6 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right!")
	}
}
