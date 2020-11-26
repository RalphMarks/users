package users

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, Nicholas o/"
	if got := Hello("Nicholas"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
