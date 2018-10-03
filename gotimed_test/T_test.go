package gotimed_test

import (
	"testing"

	"github.com/AnimusPEXUS/gotimed"
)

func TestRfc868ToUnix(t *testing.T) {
	r := gotimed.Rfc868ToUnix(1538488003)
	if r != 3747476803 {
		t.Error("invalid time conversion")
	}
}

func TestUnixToRfc(t *testing.T) {
	r := gotimed.UnixToRfc(3747476803)
	if r != 1538488003 {
		t.Error("invalid time conversion")
	}
}
