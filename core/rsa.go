package core

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
	"strings"
)

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "\n-----END RSA PRIVATE KEY-----"
)

func RsaSign(plaintext string, privateKey string) (ciphertext string, err error) {
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		log.Printf("ParsePrivateKey err: %v", err)
		return
	}

	hash := sha256.New()
	if _, err = hash.Write([]byte(plaintext)); err != nil {
		log.Printf("hash.Write error: %v", err)
		return
	}
	digest := hash.Sum(nil)
	rsaSign, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, digest)
	if err != nil {
		log.Printf("rsa.SignPKCS1v15 error: %v", err)
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(rsaSign)
	return
}

func ParsePrivateKey(privateKey string) (priKey *rsa.PrivateKey, err error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	var block, _ = pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Printf("x509.ParsePKCS1PrivateKey privateKey %s error: %v", privateKey, err)
		return nil, err
	}
	return priKey, nil
}

func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}

type rsaSign struct {
	privateKey string
}

func NewRsaSing(privateKey string) Signer {
	return &rsaSign{privateKey}
}

func (r *rsaSign) Type() string {
	return "rsa"
}

func (r *rsaSign) Sign(params string) (string, error) {
	return RsaSign(params, r.privateKey)
}
