package goutils

import (
	"log"
	"time"
)

func TrackTime(trackFor string, startAt time.Time) {
	var (
		endAt   = time.Now()
		latency = endAt.Sub(startAt)
	)

	log.Printf("| %10v | %s", latency, trackFor)
}
