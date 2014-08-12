package goutils

import (
	"fmt"
)

func CoveredGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%+v", r)
				PrintStackAndError(err)
			}
		}()

		f()
	}()
}
