package lab1

import (
	"main/util"
	"strconv"
	"main/types"
)

const (
	LAB_NAME = "lab1"
	OUT_PATH string ="_out"
	SOUND_NAME string = "_sound"

	FIRST_OPTION__PATH string = LAB_NAME + OUT_PATH + "_1_"
	SECOND_OPTION__PATH string = LAB_NAME + OUT_PATH + "_2"
	THIRD_OPTION__PATH string = LAB_NAME + OUT_PATH + "_3"

	THIRD_OPTION__AMOUNT_CYCLES = 5
)


func PerformFirstOption()  {
	makeFirstOptionFirstSub()
	makeFirstOptionSecondSub()
	makeFirstOptionThirdSub()
}

func makeFirstOptionSub(subName string, values []float64, createFunction func(x float64) func(x float64) float64)  {
	var functionsData []types.PlotData

	for index, element := range values {
		function := createFunction(element)
		functionsData = append(functionsData, types.PlotData{
			Function: function,
			Name: strconv.Itoa(index),
			InitialN: 0,
			EndN: 200,
			Step: 0.001,
		})
		util.WriteWAV(FIRST_OPTION__PATH + subName + SOUND_NAME + "_" + strconv.Itoa(index + 1), util.SOUND_LENGTH, function)
	}

	util.CreatePlot("n", "f(n)", FIRST_OPTION__PATH + subName, functionsData)
}

func makeFirstOptionFirstSub() {
	makeFirstOptionSub("a", PHI_VALUES, createFirstFunctionA)
}

func makeFirstOptionSecondSub() {
	makeFirstOptionSub("b", FREQUENCY_VALUES, createFirstFunctionB)
}

func makeFirstOptionThirdSub() {
	makeFirstOptionSub("c", AMPLITUDE_VALUES, createFirstFunctionC)
}


func PerformSecondOption()  {
	var plotData []types.PlotData
	function := createSecondFunction()
	plotData = append(plotData, types.PlotData{
		Function: function,
		Name: SECOND_OPTION__PATH,
		InitialN: 0,
		EndN: 2000,
		Step: 0.01,
	})
	util.CreatePlot("n", "f(n)", SECOND_OPTION__PATH, plotData)

	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME , util.SOUND_LENGTH, function)
}


func PerformThirdOption()  {
	var plotData []types.PlotData
	function := createThirdFunction(THIRD_OPTION__AMOUNT_CYCLES)
	plotData = append(plotData, types.PlotData{
		Function: function,
		Name: THIRD_OPTION__PATH,
		InitialN: 0,
		EndN: 1,
		Step: 0.005,
	})
	util.CreatePlot("n", "f(n)",THIRD_OPTION__PATH, plotData)

	util.WriteWAV(THIRD_OPTION__PATH + SOUND_NAME, util.SOUND_LENGTH, function)
}
