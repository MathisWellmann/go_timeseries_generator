package go_timeseries_generator

import (
	"fmt"
	"testing"
)

func TestUniformProcess(t *testing.T) {
	vals := UniformProcess(1024)

	filename := fmt.Sprintf("img/uniform_process.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
