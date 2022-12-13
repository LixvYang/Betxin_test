package timewheel

import (
	"math/rand"
	"strconv"
	"time"
)

var tw = New(time.Second, 3600)

func init() {
	tw.Start()
}

// Delay executes job after waiting the given duration
func Delay(duration time.Duration, key string, job func()) {
	tw.AddJob(duration, key, job)
}

// At executes job at given time
func At(at time.Time, key string, job func()) {
	tw.AddJob(at.Sub(time.Now()), key, job)
}

// Every execfutes job at every time
func Every(every time.Duration, job func()) {
	for {
		// tw.AddJob(every, strconv.Itoa(int(rand.Float64())), job)
		At(time.Now().Add(every), strconv.Itoa(int(rand.Float64())), job)
		time.Sleep(every)
	}
}

// Cancel stops a pending job
func Cancel(key string) {
	tw.RemoveJob(key)
}
