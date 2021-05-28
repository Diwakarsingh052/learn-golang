package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const TokenBytes = 64

func bytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func stringGen(nBytes int) (string, error) {
	b, err := bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil

}
func Token() (string, error) {
	return stringGen(TokenBytes)
}
