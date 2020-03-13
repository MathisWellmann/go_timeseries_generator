package timeseries_generator

import "math/rand"

// GaussianHL returns two series High and Low based on a common gaussian process
func GaussianHL(length int) ([]float64, []float64) {
	highs := make([]float64, length)
	lows := make([]float64, length)
	vals := GaussianProcess(length)

	for i := 0; i < len(vals); i++ {
		r := rand.Float64()
		highs[i] = vals[i] + r
		lows[i] = vals[i] - r
	}

	return highs, lows
}
