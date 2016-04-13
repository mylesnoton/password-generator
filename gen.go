package passgen

import (
	"crypto/rand"
	"errors"
	//"fmt"
	"math/big"
)

func NewPassword(length int) (string, error) {
	return generate(length)
}

func generate(length int) (string, error) {
	if length < 8 || length > 50 {
		return "", errors.New("Your password should be between 8 and 50 characters long")
	}

	password := ""
	char := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")

	for i := 0; i < length; i++ {
		num, _ := generateRandomNumber()
		password += string(char[num])
	}

	return password, nil
}

func generateRandomNumber() (int64, error) {
	max := *big.NewInt(92)

	n, err := rand.Int(rand.Reader, &max)

	if err != nil {
		return 0, errors.New("Couldn't reliably generate a random number")
	}

	return n.Int64(), nil
}
