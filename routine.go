package goutils

import (
	"fmt"
)

var isAsync bool = true

func DisableGoroutine() {
	isAsync = false
}

func CoveredGo(funcs ...func()) {
	if isAsync {
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

	} else {
		for _, f := range funcs {
			f()
		}
	}

}
