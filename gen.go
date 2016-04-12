package passgen

import (
	"crypto/rand"
	"errors"
	//"fmt"
	"log"
	"math/big"
	"time"
)

func NewPassword(length int) (string, error) {
	return generate(length)
}

func generate(length int) (string, error) {
	start := time.Now()

	if length < 8 || length > 100 {
		return "", errors.New("Your password should be between 8 and 100 characters long")
	}

	password := ""
	char := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*,+-./_[]:;<>=@?{}|~^")

	for i := 0; i < length; i++ {
		num, _ := generateRandomNumber()
		password += string(char[num])
	}

	elapsed := time.Since(start)
	log.Printf("Exec took %s", elapsed)
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
