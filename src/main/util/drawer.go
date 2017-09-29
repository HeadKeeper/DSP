package util

import (
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
	"fmt"
	"main/types"
)

func CreatePlot(axisXName string, axisYName string, plotName string, functions []types.FunctionData)  {
	CreatePlotWithStyle(axisXName, axisYName, plotName, "lines", functions)
}

func CreatePlotWithStyle(axisXName string, axisYName string, plotName string, styleName string, functions []types.FunctionData) {
	plotter,err := gnuplot.NewPlotter("", false, false)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer plotter.Close()

	plotter.SetStyle(styleName)

	for _, functionData := range functions {
		plotter.PlotFunc(GetValues(functionData.InitialN, functionData.EndN, functionData.Step), functionData.Function, functionData.Name)
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

func CreateXYPlot(axisXName string, axisYName string, arrayX []float64, arrayY []float64, plotName string)  {
	plotter,err := gnuplot.NewPlotter("", false, false)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer plotter.Close()

	plotter.SetStyle("lines")

	plotter.PlotXY(arrayX, arrayY, plotName)

	plotter.SetYLabel(axisYName)
	plotter.SetXLabel(axisXName)

	plotter.CheckedCmd("set terminal png")
	plotter.CheckedCmd("set key bmargin left horizontal Right noreverse enhanced autotitles box linetype -1 linewidth 1.000")
	plotter.CheckedCmd("set output './out/" + plotName + ".png'")

	plotter.CheckedCmd("replot")
	plotter.CheckedCmd("q")

	return
}