package goutils

import (
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
func MillisecondToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}
