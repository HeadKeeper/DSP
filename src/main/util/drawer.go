package util

import (
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const OUT_PATH string = "out/"
const DEFAULT_LINE_WIDTH float64 = 2

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

func AddFunctionOnPlotWithLegend(plot *plot.Plot, function func(n float64) float64, color color.RGBA, functionName string) {
	fun := plotter.NewFunction(function)
	fun.Color = color
	fun.Width = vg.Points(DEFAULT_LINE_WIDTH)
	plot.Legend.Add(functionName, fun)
	plot.Add(fun)
}

func SavePlotImage(name string, plot *plot.Plot) error {
	return plot.Save(4*vg.Inch, 4*vg.Inch, OUT_PATH + name + ".png")
}