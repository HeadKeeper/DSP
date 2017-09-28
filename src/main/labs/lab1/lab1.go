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

	FIRST_OPTION__PATH string = LAB_NAME + OUT_PATH + "_first_"
	SECOND_OPTION__PATH string = LAB_NAME + OUT_PATH + "_second"
	THIRD_OPTION__PATH string = LAB_NAME + OUT_PATH + "_third"

	THIRD_OPTION__AMOUNT_CYCLES = 5
)


func PerformFirstOption()  {
	makeFirstOptionFirstSub()
	makeFirstOptionSecondSub()
	makeFirstOptionThirdSub()
}

func makeFirstOptionSub(subName string, values []float64, createFunction func(x float64, soundLength float64) func(x float64) float64)  {
	var functionsData []types.FunctionData

	for index, element := range values {
		function := createFunction(element, float64(util.SOUND_LENGTH))
		functionsData = append(functionsData, types.FunctionData{
			Function: function,
			Name: strconv.Itoa(index),
			InitialN: 0,
			EndN: 10,
			Step: 0.001,
		})
		util.WriteWAV(FIRST_OPTION__PATH + subName + SOUND_NAME + "_" + strconv.Itoa(index + 1), util.SOUND_LENGTH, function)
	}

	util.CreatePlot("n", "f(n)", FIRST_OPTION__PATH + subName, functionsData)
}

func makeFirstOptionFirstSub() {
	makeFirstOptionSub("a", FIRST_OPTION__a__y__VALUES, createFirstFunctionA)
}

func makeFirstOptionSecondSub() {
	makeFirstOptionSub("b", FIRST_OPTION__b__f__VALUES, createFirstFunctionB)
}

func makeFirstOptionThirdSub() {
	makeFirstOptionSub("c", FIRST_OPTION__c__A__VALUES, createFirstFunctionC)
}


func PerformSecondOption()  {
	var value []types.FunctionData
	function := createSecondFunction(float64(util.SOUND_LENGTH))
	value = append(value, types.FunctionData {
		Function: function,
		Name: SECOND_OPTION__PATH,
		InitialN: 0,
		EndN: 100,
		Step: 0.01,
	})
	util.CreatePlot("n", "f(n)", SECOND_OPTION__PATH, value)

	util.WriteWAV(SECOND_OPTION__PATH + SOUND_NAME , util.SOUND_LENGTH, function)
}


func PerformThirdOption()  {
	var value []types.FunctionData
	function := createThirdFunction(THIRD_OPTION__INIT_HARMONIC, THIRD_OPTION__AMOUNT_CYCLES, float64(util.SOUND_LENGTH))
	value = append(value, types.FunctionData {
		Function: function,
		Name: THIRD_OPTION__PATH,
		InitialN: 0,
		EndN: 50,
		Step: 0.1,
	})
	util.CreatePlot("n", "f(n)",THIRD_OPTION__PATH, value)
	util.WriteWAV(THIRD_OPTION__PATH + SOUND_NAME, util.SOUND_LENGTH, function)

}
