package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC s little wrapper of crypto/hmac package.
type HMAC struct {
	hmac hash.Hash
}

// HashString will hash provided string with hmac.
func (h HMAC) HashString(s string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(s))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}


// NewHMAC returns new HMAC object with hashed secret key.
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{hmac: h}
}
