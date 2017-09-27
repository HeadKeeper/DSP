package lab1

import (
	"main/util"
	"strconv"
	"math"
)

const (
	SOUND_LENGTH = 5

	LAB_NAME = "lab1"
	INITIAL_PATH string = LAB_NAME + "/"
	PLOT_NAME string = "plot"
	SOUND_NAME string = "sound"

	FIRST_OPTION__PATH string = INITIAL_PATH + "1/"
	SECOND_OPTION__PATH string = INITIAL_PATH + "2/"
	THIRD_OPTION__PATH string = INITIAL_PATH + "3/"
)

func Start() {

}

func PerformFirstOption()  {
	makeFirstOptionFirstSub()
	makeFirstOptionSecondSub()
	makeFirstOptionThirdSub()
}

func makeFirstOptionSub(subName string, values []float64, createFormula func(x float64, soundLength float64) func(x float64) float64)  {
	plot := util.CreatePlot("Waves", "", "", -10, 10000, -20, 20)

	for index, element := range values {
		formula := createFormula(element, float64(SOUND_LENGTH))

		util.AddFunctionOnPlotWithLegend(plot, formula, util.CreateRandomColor(), strconv.Itoa(index + 1))
		util.WriteWAV(FIRST_OPTION__PATH + subName + SOUND_NAME + "_" + strconv.Itoa(index + 1), SOUND_LENGTH, formula)
	}


	util.SavePlotImage(FIRST_OPTION__PATH + subName + PLOT_NAME, plot)
}

func makeFirstOptionFirstSub() {
	makeFirstOptionSub("a/", FIRST_OPTION__a__y__VALUES, createFirstFormulaA)
}

func makeFirstOptionSecondSub() {
	makeFirstOptionSub("b/", FIRST_OPTION__b__f__VALUES, createFirstFormulaB)
}

func makeFirstOptionThirdSub() {
	makeFirstOptionSub("c/", FIRST_OPTION__c__A__VALUES, createFirstFormulaC)
}

func PerformSecondOption()  {
	plot := util.CreatePlot("Waves", "", "", -100, 1000000, -35, 35)

	formulaWithNilY := createSecondFormula(0, SOUND_LENGTH + 100)
	util.AddFunctionOnPlotWithLegend(plot, formulaWithNilY, util.CreateRandomColor(), "ay = 0")
	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME + "_ay_0", SOUND_LENGTH, formulaWithNilY)
	formulaWithPiDiv2Y := createSecondFormula(math.Pi / 2, SOUND_LENGTH + 100)
	util.AddFunctionOnPlotWithLegend(plot, formulaWithPiDiv2Y, util.CreateRandomColor(), "ay = Pi / 2")
	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME + "_ay_piDiv2", SOUND_LENGTH, formulaWithPiDiv2Y)
	formulaWith2PiY := createSecondFormula(math.Pi, SOUND_LENGTH + 100)
	util.AddFunctionOnPlotWithLegend(plot, formulaWith2PiY, util.CreateRandomColor(), "ay = Pi")
	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME + "_ay_Pi", SOUND_LENGTH, formulaWith2PiY)

	util.SavePlotImage(SECOND_OPTION__PATH + PLOT_NAME, plot)
}

func PerformThirdOption()  {

}
