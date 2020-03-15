package go_timeseries_generator

import "math/rand"

// randChange returns a random number in the range (-1, 1)
func RandChange() float64 {
	r := rand.Float64()
	// with 50/50 probability
	if rand.Float64() < 0.5 {
		r *= -1 // make random number negative
	}
	return r
}
