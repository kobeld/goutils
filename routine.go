package goutils

import (
	"fmt"
)

func CoveredGo(funcs ...func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%+v", r)
				PrintStackAndError(err)
			}
		}()

		for _, f := range funcs {
			f()
		}
	}()
}
