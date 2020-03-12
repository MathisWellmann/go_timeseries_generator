package timeseries_generator

import "math"

// GenerateSineWave will output the sine wave up to 2*pi on the x axis over the given length
func SineWave(length int) []float64 {
	out := make([]float64, length)

	for i := 0; i < len(out); i++ {
		out[i] = 0.99 * math.Sin(float64(i)/float64(len(out))*2.0*math.Pi)
	}
	return out
}
