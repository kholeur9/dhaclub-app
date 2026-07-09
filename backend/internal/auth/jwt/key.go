package jwt

import (
	"crypto/rand"
	"fmt"
)

func generateKeyToken() (*[]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%x", b)
	return &b, nil
}