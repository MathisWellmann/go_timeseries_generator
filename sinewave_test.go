package timeseries_generator

import (
	"fmt"
	"testing"
)

func TestSineWave(t *testing.T) {
	sine := SineWave(1024)

	filename := fmt.Sprintf("img/sinewave.png")
	err := Plt(sine, filename)
	if err != nil {
		t.Error(err)
	}
}
