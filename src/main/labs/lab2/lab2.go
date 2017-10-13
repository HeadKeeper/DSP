package lab2

import (
	"main/util"
	"main/types"
	"fmt"
	"strconv"
	"math"
)

const (
	LAB_NAME = "lab2"
	OUT_PATH string = "_out"
	SOUND_NAME string = "_sound"

	INITIAL_PATH string = LAB_NAME + OUT_PATH

	PHI float64 = math.Pi / 2
)

func PerformFirstOption()  {
	performOption("1", createFunction())
}

func PerformSecondOption()  {
	performOption("2", createSecondFunction(PHI))
}

func performOption(optionName string, functionCreator func(n float64) float64)  {
	var values []types.PlotData

	fmt.Println("Output for " + LAB_NAME + "_" + optionName)
	function := functionCreator
	var statistic []types.SecondLabStat
	for _, currentM := range createMArray() {
		currentFunctionData := types.PlotData{
			Function: function,
			Name: "M = " + strconv.Itoa(int(currentM)),
			InitialN: 0,
			EndN: currentM,
			Step: 0.1,
		}
		values = append(values, currentFunctionData)
		statistic = append(statistic, analyzeData(currentM, currentFunctionData))
	}
	drawStatisticData(optionName, statistic)
	drawPlotAndMakeSound(optionName, function, values)
}

func drawPlotAndMakeSound(optionName string, function func(n float64) float64, values []types.PlotData) {
	util.CreatePlotWithStyle("n", "f(n)", INITIAL_PATH + "_" + optionName, "points", values)
	util.WriteWAV(INITIAL_PATH + "_" + optionName + SOUND_NAME, util.SOUND_LENGTH, function)
}

func drawStatisticData(optionName string, statistic []types.SecondLabStat) {
	var arrayM []float64
	var arrayRMS1 []float64
	var arrayRMS2 []float64
	var arrayAmplitude []float64

	for _, currentData := range statistic {
		arrayM = append(arrayM, currentData.M)
		arrayRMS1 = append(arrayRMS1, currentData.RMSFirstInaccuracy)
		arrayRMS2 = append(arrayRMS2, currentData.RMSSecondInaccuracy)
		arrayAmplitude = append(arrayAmplitude, currentData.AmplitudeInaccuracy)
	}

	util.CreateXYPlot("M", "RMS1", arrayM, arrayRMS1, INITIAL_PATH + "_" + optionName + "_M-RMS1")
	util.CreateXYPlot("M", "RMS2", arrayM, arrayRMS2, INITIAL_PATH + "_" + optionName + "_M-RMS2")
	util.CreateXYPlot("M", "Amplitude", arrayM, arrayAmplitude, INITIAL_PATH + "_" + optionName + "_M-Amplitude")
}

func analyzeData(M float64, data types.PlotData) types.SecondLabStat {
	/*firstRMS*/ _ , firstRMSInaccuracy, _ /*secondRMS*/, secondRMSInaccuracy := CalculateRootMeanSquareValues(data)
	/*amplitude*/ _ , amplitudeInaccuracy := CalculateAmplitudeValue(data)
	/*fmt.Println(fmt.Sprintf(
		"M: %v\n\t" +
		"First RMS: %v (Inaccuracy: %v)\n\t" +
		"SecondRMS: %v (Inaccuracy: %v)\n\t" +
		"Amplitude: %v (Inaccuracy: %v)\n",
		M,
		firstRMS, firstRMSInaccuracy,
		secondRMS, secondRMSInaccuracy,
		amplitude, amplitudeInaccuracy,
	))*/
	return types.SecondLabStat{
		M                  : M,
		RMSFirstInaccuracy : firstRMSInaccuracy,
		RMSSecondInaccuracy: secondRMSInaccuracy,
		AmplitudeInaccuracy: amplitudeInaccuracy,
	}
}
