package lab1

import (
	"main/util"
	"strconv"
	"main/types"
)

const (
	LAB_NAME = "lab1_"
	SOUND_NAME string = "_sound"

	FIRST_OPTION__PATH string = LAB_NAME + "1_"
	SECOND_OPTION__PATH string = LAB_NAME + "2"
	THIRD_OPTION__PATH string = LAB_NAME + "3"

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
		functionData := types.PlotData{
			Function: function,
			Name: strconv.Itoa(index),
			InitialN: 0,
			EndN: 50000,
			Step: 1,
		}
		functionsData = append(functionsData, functionData)
		util.WriteWAV(FIRST_OPTION__PATH + subName + SOUND_NAME + "_" + strconv.Itoa(index + 1), util.SOUND_LENGTH, functionData)
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
	functionData := types.PlotData{
		Function: function,
		Name: SECOND_OPTION__PATH,
		InitialN: 0,
		EndN: 50000,
		Step: 1,
	}
	plotData = append(plotData, functionData)
	util.CreatePlot("n", "f(n)", SECOND_OPTION__PATH, plotData)

	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME , util.SOUND_LENGTH, functionData)
}


func PerformThirdOption()  {
	var plotData []types.PlotData
	function := createThirdFunction(THIRD_OPTION__AMOUNT_CYCLES)
	functionData := types.PlotData{
		Function: function,
		Name: THIRD_OPTION__PATH,
		InitialN: 0,
		EndN: 100000,
		Step: 5,
	}
	plotData = append(plotData, functionData)
	util.CreatePlot("n", "f(n)",THIRD_OPTION__PATH, plotData)

	util.WriteWAV(THIRD_OPTION__PATH + SOUND_NAME, util.SOUND_LENGTH, functionData)
}
