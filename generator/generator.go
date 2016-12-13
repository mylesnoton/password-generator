package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	chars    = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	specials = []byte("!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")
	numbers  = []byte("0123456789")
)

// NewPassword creates a new random password string based on the arguments it is passed.
func NewPassword(length int, noSpecial bool) (string, error) {
	if length < 6 || length > 200 {
		err := errors.New("Your password length should be between 8 and 200 characters long")
		return "", err
	}

	passwordBytes := generatePassword(length, noSpecial)

	return string(passwordBytes), nil
}

func generatePassword(length int, noSpecial bool) []byte {
	passwordBytes := make([]byte, length)
	finalChars := make([]byte, 93)

	if noSpecial == true {
		finalChars = append(finalChars, chars...)
	} else {
		copy(finalChars, chars)
		finalChars = append(finalChars, specials...)
	}

	for i := 0; i < 10; i++ {
		for ii := 0; ii < length*4; ii++ {
			position := randomNumber(len(finalChars))
			placement := randomNumber(len(passwordBytes))

			if passwordBytes[placement] == 0 {
				passwordBytes[placement] = finalChars[position]
			}
		}
	}

	if !validatePassword(passwordBytes, noSpecial) {
		generatePassword(length, noSpecial)
	}

	return passwordBytes
}

func validatePassword(passwordBytes []byte, noSpecial bool) bool {
	foundNumber := false
	foundSpecial := false

	for _, passChar := range passwordBytes {
		for _, lookupChar := range numbers {
			if int(lookupChar) == int(passChar) {
				foundNumber = true
			}
		}

		for _, lookupChar := range specials {
			if string(lookupChar) == string(passChar) {
				foundSpecial = true
			}
		}
	}

	// Pretend this was found if we don't even want it
	if noSpecial {
		foundSpecial = true
	}

	if !foundNumber || !foundSpecial {
		return false
	}

	return true
}

func randomNumber(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
