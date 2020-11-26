package users

import (
	"testing"
)

func TestCreateUserS(t *testing.T) {
	if user, err := CreateUser("Nicholas", "nico@eiya.com", "xxxxx"); err != nil {
		t.Errorf("Hello() = %q, want %q", user.Username, err)
	}
}
