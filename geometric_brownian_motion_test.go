package go_timeseries_generator

import "testing"

func TestGeometricBrownianMotion(t *testing.T) {
	vals := GeometricBrownianMotionDefault(1024)
	err := Plt(vals, "img/geometric_brownian_motion.png")
	if err != nil {
		t.Error(err)
	}
}
