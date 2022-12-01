package credential

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// 3DES加密
func TripleDesEncrypt(origData, des3key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(des3key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, des3key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, des3key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(des3key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, des3key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后⼀一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
