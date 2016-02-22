package goutils

import (
	"log"
	"time"
)

var isTrackTimeEnabled = true

func DisableTrackTime() {
	isTrackTimeEnabled = false
}

func TrackTime(trackFor string, startAt time.Time) {

	if !isTrackTimeEnabled {
		return
	}

	var (
		endAt   = time.Now()
		latency = endAt.Sub(startAt)
	)

	log.Printf("| %10v | %s", latency, trackFor)
}
