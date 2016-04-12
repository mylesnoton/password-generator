package passgen

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	password, _ := NewPassword(10)

	if len(password) != 10 {
		t.Fail()
	}
}

func TestNewPasswordHasLowerChar(t *testing.T) {
	password, _ := NewPassword(10)

	lowers := []byte("abcdefghijklmnopqrstuvwxyz")
	found := false

	for _, find_char := range lowers {
		for _, password_char := range password {
			if string(find_char) == string(password_char) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordHasUpperChar(t *testing.T) {
	password, _ := NewPassword(10)

	uppers := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	found := false

	for _, find_char := range uppers {
		for _, password_char := range password {
			if string(find_char) == string(password_char) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordHasSpecialChar(t *testing.T) {
	password, _ := NewPassword(10)

	specials := []byte("!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")
	found := false

	for _, find_char := range specials {
		for _, password_char := range password {
			if string(find_char) == string(password_char) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestNewPasswordHasNumber(t *testing.T) {
	password, _ := NewPassword(10)

	numbers := []byte("0123456789")
	found := false

	for _, find_number := range numbers {
		for _, password_char := range password {
			if int(find_number) == int(password_char) {
				found = true
			}
		}
	}

	if found != true {
		t.Fail()
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	number, _ := generateRandomNumber()

	if number != int64(number) {
		t.Fail()
	}
}
