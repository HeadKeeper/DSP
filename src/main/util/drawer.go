package util

import (
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
	"fmt"
	"main/types"
)

const OUT_PATH string = "out/"

func CreatePlot(axisXName string, axisYName string, plotName string, functions []types.FunctionData) {
	plotter,err := gnuplot.NewPlotter("", false, false)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer plotter.Close()

	plotter.SetStyle("lines")

	for _, functionData := range functions {
		plotter.PlotFunc(getValues(functionData.InitialN, functionData.EndN, functionData.Step), functionData.Function, functionData.Name)
	}

	plotter.SetYLabel(axisYName)
	plotter.SetXLabel(axisXName)

	plotter.CheckedCmd("set terminal png")
	plotter.CheckedCmd("set key bmargin left horizontal Right noreverse enhanced autotitles box linetype -1 linewidth 1.000")
	plotter.CheckedCmd("set output './out/" + plotName + ".png'")

	plotter.CheckedCmd("replot")
	plotter.CheckedCmd("q")

	return
}

func getValues(initialN float64, endN float64, step float64) []float64 {
	var n float64
	var values []float64
	for n = initialN; n < endN; n += step {
		values = append(values, n)
	}

	return values
}