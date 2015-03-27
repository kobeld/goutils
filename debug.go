package goutils

import (
	"log"
	"runtime/debug"
)

// Set true to enable the stack tracker
var enabled = true

func DisablePrintStack() {
	enabled = false
}

func PrintStackAndError(err error) {
	if enabled {
		log.Printf("********** Debug Error message: %+v ***********\n", err)
		debug.PrintStack()
	}
}

func HasErrorAndPrintStack(err error) bool {
	if err == nil {
		return false
	}

	PrintStackAndError(err)
	return true
}

func PrintStackButSwallowError(err error) {
	if err == nil {
		return
	}

	PrintStackAndError(err)
	err = nil
	return
}
