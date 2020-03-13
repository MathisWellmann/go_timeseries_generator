package timeseries_generator

import (
	"fmt"
	"testing"
)

func TestGaussianProcess(t *testing.T) {
	vals := GaussianProcess(1024)

	filename := fmt.Sprintf("img/gaussian_process.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
