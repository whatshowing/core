package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type CipherService interface {
	Encrypt(data []byte, keyPhrase string) []byte
	Decrypt(data []byte, keyPhrase string) []byte
}

type cipherService struct{}

func NewCipherService() CipherService {
	return &cipherService{}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (*cipherService) Encrypt(data []byte, keyPhrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(keyPhrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func (*cipherService) Decrypt(data []byte, keyPhrase string) []byte {
	key := []byte(createHash(keyPhrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	clearText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return clearText
}
