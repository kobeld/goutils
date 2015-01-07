package goutils

import (
	"log"
	"runtime/debug"
)

// Set true to enable the stack tracker
var enabled = false

func EnablePrintStack() {
	enabled = true
}

func PrintStackAndError(err error) {
	if enabled {
		log.Printf("********** Debug Error message: %+v ***********\n", err)
		debug.PrintStack()
	}
}
