package goutils

import (
	"fmt"
	"strconv"
	"time"
)

// Compare to the offical time.Parse, this ignore the error message and
// return a blank time.Time{} object when value has something wrong
func ParseTime(layout, value string) time.Time {
	if value == "" {
		return time.Time{}
	}

	r, err := time.Parse(layout, value)
	if err != nil {
		PrintStackAndError(err)
		return time.Time{}
	}

	return r
}

// Compare to the offical time.Format, it will return blank string when
// the time is zero, rather than return "0001-01-01 00:00"
func FormatTime(theTime time.Time, layout string) (r string) {
	if theTime.IsZero() {
		return ""
	}

	return theTime.Format(layout)
}

// Millisecond e.g. 1445485125599
func MillisecondToTime(ms string) (theTime time.Time, err error) {

	msInt, err := strconv.ParseInt(ms, 10, 64)
	if HasErrorAndPrintStack(err) {
		return
	}

	theTime = time.Unix(0, msInt*int64(time.Millisecond))
	return
}

func TimeToMillisecond(theTime time.Time) string {
	return fmt.Sprintf("%d", theTime.UnixNano()/int64(time.Millisecond))
}
