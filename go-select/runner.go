package goselect

import (
	"net/http"
	"time"
)

func Runner(a, b string) (winner string) {
	durationA := calculateResponseTime(a)
	durationB := calculateResponseTime(b)

	if durationA < durationB {
		return a
	}

	return b
}

func calculateResponseTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	return time.Since(start)
}