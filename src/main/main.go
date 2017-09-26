package main

import (
	"main/util"
	"image/color"
	"math"
)

const DEFAULT_PLOT_NAME string = "Functions"

/* Variant 6 */

func main() {

	//lab1.PerformFirstOption()

	plot := util.CreatePlot("Waves", "", "", -10000, 10000, -10000, 10000)
	util.AddFunctionOnPlot(plot, func(x float64) float64 { return 9 * math.Sin((2*math.Pi*2*x)/44100*math.Pi) }, color.RGBA{R: 255, A: 255})
	util.SavePlotImage("/lab1/waves", plot)

	util.WriteWAV("lab1/waves", 15, func(x float64) float64 { return 9 * math.Sin((2*math.Pi*2*x)/44100*math.Pi) })
}