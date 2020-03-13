# timeseries generator
Generate common uniformly sampled timeseries using Go

## Requirements
- Gonum for plotting in tests:
``
go get -u gonum.org/v1/gonum/...
``

## Installation:
``
go get github.com/MathisWellmann/timeseries-generator
``

### Functions include:
- SineWave
- StepFunction
- UniformProcess
- GaussianProcess

Tests which generate images are available to get an idea of the function behaviour.

### Images:
![gaussian_process](https://github.com/MathisWellmann/timeseries_generator/img/gaussian_process.png)
![sinewave](https://github.com/MathisWellmann/timeseries_generator/img/sinewave.png)
![step_function](https://github.com/MathisWellmann/timeseries_generator/img/step_function.png)
![uniform_process](https://github.com/MathisWellmann/timeseries_generator/img/uniform_process.png)
