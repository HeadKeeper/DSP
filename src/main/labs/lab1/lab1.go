package lab1

import (
	"main/util"
	"strconv"
)

const (
	SOUND_LENGTH = 15
	AXIS_MIN_X = -10
	AXIS_MAX_X = 10000
	AXIS_MIN_Y = -30
	AXIS_MAX_Y = 30

	LAB_NAME = "lab1"
	INITIAL_PATH string = LAB_NAME + "/"
	PLOT_NAME string = "plot"
	SOUND_NAME string = "sound"

	FIRST_OPTION__PATH string = INITIAL_PATH + "1/"
	SECOND_OPTION__PATH string = INITIAL_PATH + "2/"
	THIRD_OPTION__PATH string = INITIAL_PATH + "3/"

	THIRD_OPTION__AMOUNT_CYCLES = 5
)


func PerformFirstOption()  {
	makeFirstOptionFirstSub()
	makeFirstOptionSecondSub()
	makeFirstOptionThirdSub()
}

func makeFirstOptionSub(subName string, values []float64, createFormula func(x float64, soundLength float64) func(x float64) float64)  {
	plot := util.CreatePlot("Harmonic", "n", "f(n)", AXIS_MIN_X, AXIS_MAX_X, -20, 20)

	for index, element := range values {
		formula := createFormula(element, float64(SOUND_LENGTH))

		util.AddFunctionOnPlotWithLegend(plot, formula, util.CreateRandomColor(), strconv.Itoa(index + 1))
		util.WriteWAV(FIRST_OPTION__PATH + subName + SOUND_NAME + "_" + strconv.Itoa(index + 1), SOUND_LENGTH, formula)
	}


	util.SavePlotImage(FIRST_OPTION__PATH + subName + PLOT_NAME, plot)
}

func makeFirstOptionFirstSub() {
	makeFirstOptionSub("a/", FIRST_OPTION__a__y__VALUES, createFirstFunctionA)
}

func makeFirstOptionSecondSub() {
	makeFirstOptionSub("b/", FIRST_OPTION__b__f__VALUES, createFirstFunctionB)
}

func makeFirstOptionThirdSub() {
	makeFirstOptionSub("c/", FIRST_OPTION__c__A__VALUES, createFirstFunctionC)
}


func PerformSecondOption()  {
	plot := util.CreatePlot("Polyharmonic", "n", "f(n)", AXIS_MIN_X,1000000, -40, 40)

	function := createSecondFunction(0, float64(SOUND_LENGTH))
	util.AddFunctionOnPlotWithLegend(plot, function, util.CreateRandomColor(), "")
	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME , SOUND_LENGTH, function)

	util.SavePlotImage(SECOND_OPTION__PATH + PLOT_NAME, plot)
}


func PerformThirdOption()  {
	plot := util.CreatePlot("Polyharmonic", "n", "f(n)", AXIS_MIN_X, 10, -500, 500)

	function := createThirdFunction(THIRD_OPTION__INIT_HARMONIC, THIRD_OPTION__AMOUNT_CYCLES, float64(SOUND_LENGTH))
	util.AddFunctionOnPlotWithLegend(plot, function, util.CreateRandomColor(), "Dynamic")
	util.WriteWAV(THIRD_OPTION__PATH + SOUND_NAME, SOUND_LENGTH / 5, function)

	util.SavePlotImage(THIRD_OPTION__PATH + PLOT_NAME, plot)
}
