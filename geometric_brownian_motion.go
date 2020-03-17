package go_timeseries_generator

import (
	"math"
	"math/rand"
)

func GeometricBrownianMotion(length int, initVal, drift, dt, diffusion float64) []float64 {
	out := make([]float64, length)

	out[0] = initVal

	driftFac := 1.0 + drift*dt
	diffusionFac := diffusion * math.Sqrt(dt)

	for i := 1; i < length; i++ {
		rv := driftFac + diffusionFac*rand.NormFloat64()
		prod := rv * out[i-1]
		out[i] = prod
	}
	return out
}

func GeometricBrownianMotionDefault(length int) []float64 {
	return GeometricBrownianMotion(length, 100.0, 0.15, 1.5/365, 0.5)
}
