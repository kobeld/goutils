package goutils

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Encrypt(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
