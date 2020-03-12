package timeseries_generator

// GenerateStepFunction generates a step function of given length with midpoint and width
func StepFunction(length, midPoint, width int) []float64 {
	out := make([]float64, length)
	for i := 0; i < len(out); i++ {
		if i < midPoint-width || i > midPoint+width {
			out[i] = 0
			continue
		}
		out[i] = 1
	}
	return out
}
