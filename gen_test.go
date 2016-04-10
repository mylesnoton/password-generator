package passgen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	password := generate(10)

	if len(password) != 10 {
		t.Fail()
	}
}
