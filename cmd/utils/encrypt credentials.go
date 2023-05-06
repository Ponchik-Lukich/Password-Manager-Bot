package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(text string) (string, error) {
	key := []byte(os.Getenv("ENCRYPT_KEY"))
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherT := make([]byte, aes.BlockSize+len(text))
	iv := cipherT[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(c, iv)
	stream.XORKeyStream(cipherT[aes.BlockSize:], []byte(text))
	return base64.URLEncoding.EncodeToString(cipherT), nil
}

func Decrypt(text string) (string, error) {
	key := []byte(os.Getenv("ENCRYPT_KEY"))
	tBytes, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(tBytes) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}
	iv := tBytes[:aes.BlockSize]
	tBytes = tBytes[aes.BlockSize:]
	streamLogin := cipher.NewCFBDecrypter(block, iv)
	streamLogin.XORKeyStream(tBytes, tBytes)
	return string(tBytes), nil
}
