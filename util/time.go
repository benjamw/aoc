package util

import (
	"log"
	"time"
)

// Usage: inside function:
// func foo() {
//     defer util.Duration(util.Track("foo"))
//
//     ... code ...
// }

// Track begin tracking a function call
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Duration output the duration of a function call that was tracked
func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
