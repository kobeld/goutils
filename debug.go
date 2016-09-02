package goutils

import (
	"log"
	"runtime/debug"

	"gopkg.in/mgo.v2"
)

// Set true to enable the stack tracker
var enabled = true

func DisablePrintStack() {
	enabled = false
}

func PrintStackAndError(err error) {
	if enabled {
		log.Printf("********** Debug Error message: %+v ***********\n", err)

		// Don't print the stack for the Not Found error of mgo
		if err != mgo.ErrNotFound {
			debug.PrintStack()
		}
	}
}

func HasErrorAndPrintStack(err error) bool {
	if err == nil {
		return false
	}

	PrintStackAndError(err)
	return true
}

func PrintStackButSwallowError(err *error) {
	if *err == nil {
		return
	}

	PrintStackAndError(*err)
	*err = nil
	return
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
