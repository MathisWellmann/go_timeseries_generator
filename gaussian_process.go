package go_timeseries_generator

import (
	"math"
	"math/rand"
)

// GaussianProcess is the same underlying process of UniformProcess,
// only with tanh activation function for random change which transforms the random values to gaussian pdf
func GaussianProcess(length int) []float64 {
	out := make([]float64, length)
	out[0] = rand.Float64()
	for i := 1; i < len(out); i++ {
		change := RandChange()
		if out[i-1]+change <= 0 {
			change = math.Abs(change)
		}
		out[i] = out[i-1] + math.Tanh(change)
	}
	return out
}
