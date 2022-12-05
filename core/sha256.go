package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type signSHA256 struct {
	key string
}

func NewSHA256Sign(key string) Signer {
	return &signSHA256{key}
}

func (*signSHA256) Type() string {
	return "sha256"
}

func (s *signSHA256) Sign(params string) (string, error) {
	hash := hmac.New(sha256.New, []byte(s.key))
	hash.Write([]byte(params))
	md := hash.Sum(nil)
	hashStr := hex.EncodeToString(md)
	return hashStr, nil
}
