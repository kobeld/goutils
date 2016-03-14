package goutils

import (
	"strings"
)

func ToLowerAndTrimSpace(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
