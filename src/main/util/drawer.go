package util

import (
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const OUT_PATH string = "out/"
const DEFAULT_LINE_WIDTH float64 = 2

/*
	Example of functions :
			// A quadratic function x^2
			quad := plotter.NewFunction(func(x float64) float64 { return x * x })
			quad.Color = color.RGBA{B: 255, A: 255}

			// An exponential function 2^x
			exp := plotter.NewFunction(func(x float64) float64 { return math.Pow(2, x) })
			exp.Width = vg.Points(2)
			exp.Color = color.RGBA{G: 255, A: 255}

			// The sine function, shifted and scaled
			// to be nicely visible on the plot.
			sin := plotter.NewFunction(func(x float64) float64 { return 10*math.Sin(x) + 50 })
			sin.Width = vg.Points(4)
			sin.Color = color.RGBA{R: 255, A: 255}
*/

func CreatePlot(name string, axisXName string, axisYName string, axisXMin float64, axisXMax float64, axisYMin float64, axisYMax float64) *plot.Plot {
	graph, err := plot.New()
	if err != nil {
		panic(err)
	}
	graph.Title.Text = name
	graph.X.Label.Text = axisXName
	graph.Y.Label.Text = axisYName

	graph.X.Min = axisXMin
	graph.X.Max = axisXMax
	graph.Y.Min = axisYMin
	graph.Y.Max = axisYMax

	return graph
}

func AddFunctionOnPlot(plot *plot.Plot, function func(n float64) float64, color color.RGBA) {
	fun := plotter.NewFunction(function)
	fun.Color = color
	fun.Width = vg.Points(DEFAULT_LINE_WIDTH)
	plot.Add(fun)
}

func SavePlotImage(name string, plot *plot.Plot) error {
	return plot.Save(4*vg.Inch, 4*vg.Inch, OUT_PATH + name + ".png")
}