package timeseries_generator

import "math/rand"

type (
	OHLCV struct {
		Open   float64
		High   float64
		Low    float64
		Close  float64
		Volume float64
	}
)

// GaussianHL returns two series High and Low based on a common gaussian process
// length will be output length of array
// volatilityScale is multiplier by which to scale random variable. recommended is 0.01
// volumeScale defines how much volume to put into each candle weighted with volatility of candle
func GaussianOHLCV(length int, volatilityScale, volumeScale float64) []*OHLCV {
	out := make([]*OHLCV, length)
	vals := GaussianProcess(length)

	for i := 0; i < len(vals); i++ {
		// random volatility scaled by value at time i. max 10% change
		vol := volatilityScale * rand.Float64() * vals[i]
		r0 := RandChange() * vol
		r1 := RandChange() * vol
		out[i] = &OHLCV{
			Open:   vals[i] + r0,
			High:   vals[i] + (vol / 2), // divide by two to keep original volatility of candle
			Low:    vals[i] - (vol / 2),
			Close:  vals[i] + r1,
			Volume: vol * volumeScale,
		}
	}

	return out
}

// GaussianOHLCVDefault wraps the function with default values
func GaussianOHLCVDefault(length int) []*OHLCV {
	return GaussianOHLCV(length, 0.01, 1000)
}
