package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

// RandBytes generates n random bytes, n is byte size (lenght).
// Or it will return an error if something was wrong.
// It uses crypto/rand package to generate bytes.
func RandBytes(n int) ([]byte, error) {
	rb := make([]byte, n)
	_, err := rand.Read(rb)
	if err != nil {
		return nil, err
	}
	return rb, nil
}

// RandString turns []bytes of size n into a random string. 
func RandString(n int) (string, error) {
	rs, err := RandBytes(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(rs), nil
}

// RememberToken generates remember tokens of a 64 byte size.
func RememberToken() (string, error) {
	return RandString(64)
}
