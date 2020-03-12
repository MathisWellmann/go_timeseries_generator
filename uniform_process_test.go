package timeseries_generator

import (
	"fmt"
	"testing"
)

func TestUniformProcess(t *testing.T) {
	vals := UniformProcess(1024)

	filename := fmt.Sprintf("img/TestGenerateTimeseriesGraph.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
