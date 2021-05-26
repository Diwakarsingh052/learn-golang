package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

type HMAC struct {
	hmac hash.Hash
}

func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New,[]byte(key)) // accepts a hashing func , a key on basis of which it will
	return HMAC{hmac: h} // assigning the instance of hash to our struct
}

func (h HMAC) Hash (input string) string {
	h.hmac.Reset() // reset any previous underlying hash
	h.hmac.Write([]byte(input)) // write the hash in bytes
	b:= h.hmac.Sum(nil) // it will return the hash of token or input

	return base64.URLEncoding.EncodeToString(b) // it will make sure everything maps to string and is url safe



}