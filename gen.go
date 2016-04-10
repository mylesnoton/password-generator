package passgen

import (
	"crypto/rand"
	//"fmt"
	"errors"
	"log"
	"math/big"
	"time"
)

func NewPassword(length int, use_numbers bool, use_special bool) (string, error) {
	return generate(length, use_numbers, use_special)
}

func generate(length int, use_numbers bool, use_special bool) (string, error) {
	start := time.Now()

	password := ""
	char := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '"', '#', '$', '%', '&'}
	//special := []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", ",", "+", "-", ".", "/", "_", "[", "]", ":", ";", "<", ">", "=", "@", "?", "{", "}", "|", "~", "^"}

	for _, value := range char {
		password += string(value)
	}

	elapsed := time.Since(start)
	log.Printf("Exec took %s", elapsed)
	return password, nil
}

func generateRandomNumber() (int64, error) {
	//max := *big.NewInt(92)
	max := *big.NewInt(26)
	n, err := rand.Int(rand.Reader, &max)

	if err != nil {
		return 0, errors.New("Couldn't reliably generate a random number")
	}

	return n.Int64(), nil
}
