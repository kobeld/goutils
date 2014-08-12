package goutils

import (
	"fmt"
)

func LoremImage(width, length int) string {
	return fmt.Sprintf("http://lorempixel.com/%d/%d", width, length)
}
