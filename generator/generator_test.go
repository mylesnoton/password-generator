package generator

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	password, _ := NewPassword(50, false)

	if len(password) != 50 {
		t.Fail()
	}
}

func TestRandomNumber(t *testing.T) {
	number := randomNumber(500)

	if number != int(number) {
		t.Fail()
	}
}

func TestNewPasswordHasLowerChar(t *testing.T) {
	password, _ := NewPassword(50, false)

	lower := []byte("abcdefghijklmnopqrstuvwxyz")
	found := false

	for _, findChar := range lower {
		for _, passwordChar := range password {
			if string(findChar) == string(passwordChar) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordHasUpperChar(t *testing.T) {
	password, _ := NewPassword(50, false)

	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	found := false

	for _, findChar := range upper {
		for _, passwordChar := range password {
			if string(findChar) == string(passwordChar) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordHasSpecialChar(t *testing.T) {
	password, _ := NewPassword(50, false)

	special := []byte("!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")
	found := false

	for _, findChar := range special {
		for _, passwordChar := range password {
			if string(findChar) == string(passwordChar) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordDoesNotHaveSpecialChar(t *testing.T) {
	password, _ := NewPassword(50, true)

	special := []byte("!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")
	found := false

	for _, findChar := range special {
		for _, passwordChar := range password {
			if string(findChar) == string(passwordChar) {
				found = true
			}
		}
	}

	if found {
		t.Fail()
	}
}

func TestNewPasswordHasNumber(t *testing.T) {
	password, _ := NewPassword(50, false)

	number := []byte("0123456789")
	found := false

	for _, findNumber := range number {
		for _, passwordChar := range password {
			if int(findNumber) == int(passwordChar) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}
