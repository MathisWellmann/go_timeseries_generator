package go_timeseries_generator

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

type (
	Values struct {
		Xs []float64
		Ys []float64
	}
)

func (v *Values) Len() int {
	return len(v.Xs)
}

func (v *Values) XY(index int) (float64, float64) {
	return v.Xs[index], v.Ys[index]
}

// Plt will create a line plot with given values and filename
func Plt(vals []float64, filename string) error {
	err := checkImgDirectory()
	if err != nil {
		return err
	}

	xs := make([]float64, len(vals))
	ys := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		xs[i] = float64(i)
		ys[i] = vals[i]
	}
	values := &Values{
		Xs: xs,
		Ys: ys,
	}

	p := plot.New()
	
	p.Title.Text = filename

	line, err := plotter.NewLine(values)
	if err != nil {
		return err
	}
	p.Add(line)

	if err := p.Save(297*vg.Millimeter, 210*vg.Millimeter, filename); err != nil {
		return err
	}
	return nil
}

// checkImgDirectory will make sure the img directory is available for any tests using it in the file path
func checkImgDirectory() error {
	// check if img directory already exists
	_, err := os.Stat("./img")
	if err != nil {
		if os.IsNotExist(err) {
			// create directory img
			err := os.Mkdir("./img", os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
