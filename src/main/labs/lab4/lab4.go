package lab4

import (
	"main/util"
)

const (
	OUTPUT_NAME     string = "lab4_out_"
	OUTPUT_1_NAME string = OUTPUT_NAME + "1_"
	OUTPUT_2_NAME string = OUTPUT_NAME + "2_"
	OUTPUT_3_NAME string = OUTPUT_NAME + "3_"
	OUTPUT_4_NAME string = OUTPUT_NAME + "4_"

	AMOUNT = util.SIGNAL_RATE * util.SOUND_LENGTH
)

func PerformFirstOption() {
	createSignal := CreateSignalFunction(AMOUNT)
	values, indexes := getSignalValues(createSignal)

	util.CreateXYPlot("i", "x(i)", indexes, values, OUTPUT_1_NAME + "signal")
	/*
	util.WriteWAV(OUTPUT_1_NAME + "signal_sound", util.SOUND_LENGTH, createSignal)
	*/
}

func PerformSecondOption() {
	amountRange, phasesRange, amplitudesRange := getRanges(AMOUNT)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_NAME)
}

func showRanges(amountRange []float64, phasesRange []float64, amplitudesRange []float64, outputName string) {
	util.CreateXYPlotWithStyle("k", "Phases", amountRange, phasesRange, outputName + "phases", "impulses")
	util.CreateXYPlotWithStyle("k", "Amplitudes", amountRange, amplitudesRange, outputName + "amplitudes", "impulses")
}

func PerformThirdOption() {
	createSignal := CreateSignalFunction(AMOUNT)
	values, indexes := getSignalValues(createSignal)

	newValuesByMovingAverage := filterByMovingAverageAlgorithm(values)
	newValuesByParabola := filterByParabolaAlgorithm(values)
	newValuesByMedianFilter := filterByMedianFilterAlgorithm(values)

	util.CreateXYPlot("i", "x(i)", indexes, values, OUTPUT_3_NAME + "original")
	util.CreateXYPlot("i", "x(i)", indexes, newValuesByMovingAverage, OUTPUT_3_NAME + "moving_average")
	util.CreateXYPlot("i", "x(i)", indexes, newValuesByParabola, OUTPUT_3_NAME + "parabola_eleven")
	util.CreateXYPlot("i", "x(i)", indexes, newValuesByMedianFilter, OUTPUT_3_NAME + "median_filter")
	/*
	util.WriteWAV(OUTPUT_3_NAME + "original_sound", util.SOUND_LENGTH, createSignal)
	util.WriteWAV(OUTPUT_3_NAME + "moving_average", util.SOUND_LENGTH, createSignal)
	util.WriteWAV(OUTPUT_3_NAME + "parabola_eleven", util.SOUND_LENGTH, createSignal)
	util.WriteWAV(OUTPUT_3_NAME + "median_filter", util.SOUND_LENGTH, createSignal)
	*/
}

func PerformFourthOption() {
	createSignal := CreateSignalFunction(AMOUNT)
	values, _ := getSignalValues(createSignal)

	newValuesByMovingAverage := filterByMovingAverageAlgorithm(values)
	newValuesByParabola := filterByParabolaAlgorithm(values)
	newValuesByMedianFilter := filterByMedianFilterAlgorithm(values)

	showRangesForCreatedSignal(newValuesByMovingAverage, "moving_average_")
	showRangesForCreatedSignal(newValuesByParabola, "parabola_")
	showRangesForCreatedSignal(newValuesByMedianFilter, "median_filter_")
}

func showRangesForCreatedSignal(values []float64, filterName string)  {
	amountRange, phasesRange, amplitudesRange := getRangesForCreatedSignal(AMOUNT, values)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_4_NAME + filterName)
}