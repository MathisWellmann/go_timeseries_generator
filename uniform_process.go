package go_timeseries_generator

import (
	"math"
	"math/rand"
)

// UniformProcess will return a randomly generated positive path with uniformly sampled random values
func UniformProcess(length int) []float64 {
	out := make([]float64, length)
	out[0] = rand.Float64()
	for i := 1; i < len(out); i++ {
		change := RandChange()
		if out[i-1]+change <= 0 {
			change = math.Abs(change)
		}
		out[i] = out[i-1] + change
	}
	return out
}
