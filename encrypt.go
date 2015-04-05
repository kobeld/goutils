package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func Sha1Encrypt(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Took from https://gist.github.com/manishtpatel/8222606
// Encrypt string to base64 crypto using AES
func Encrypt(keyStr string, text string) (r string, err error) {
	key := []byte(keyStr)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	r = base64.URLEncoding.EncodeToString(ciphertext)
	return
}

// Eecrypt from base64 to decrypted string
func Decrypt(keyStr, cryptoText string) (r string, err error) {
	key := []byte(keyStr)
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	r = fmt.Sprintf("%s", ciphertext)
	return
}
