package timeseries_generator

import (
	"fmt"
	"testing"
)

func TestStepFunction(t *testing.T) {
	vals := StepFunction(1024, 500, 100)
	filename := fmt.Sprintf("img/step_function.png")
	err := Plt(vals, filename)
	if err != nil {
		t.Error(err)
	}
}
