package util

import (
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	//"github.com/veandco/go-sdl2/sdl"
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

	/*if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(1200, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer.SetDrawColor(1, 1, 1, 0)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 255, 0)

	var values []sdl.Point
	for idx := 0; idx < 1000; idx++ {
		values = append(values, sdl.Point{Y: 150 + 3 * int32(function(float64(idx))), X:3 * int32(idx)})
	}
	renderer.DrawLines(values)

	renderer.Present()

	sdl.Delay(3000)*/
}

func SavePlotImage(name string, plot *plot.Plot) error {
	return plot.Save(4*vg.Inch, 4*vg.Inch, OUT_PATH + name + ".png")
}