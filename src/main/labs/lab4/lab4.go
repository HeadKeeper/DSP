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
)

func PerformFirstOption() {
	indexes, signal, meta := ReadSignal("test.wav")

	util.CreateXYPlot("i", "x(i)", indexes, signal, OUTPUT_1_NAME + "signal")
	util.WriteWAVByMeta(OUTPUT_1_NAME + "signal_sound", util.SOUND_LENGTH, signal, meta)
}

func PerformSecondOption() {
	_, signal, _ := ReadSignal("test.wav")
	amountRange, amplitudesRange, phasesRange := Fourier(signal)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_NAME)
}

func PerformThirdOption() {
	indexes, signal, meta := ReadSignal("test.wav")

	newSignalByMovingAverage := util.FilterByMovingAverageAlgorithm(signal)
	newSignalByParabola := util.FilterByParabolaAlgorithm(signal)
	newSignalByMedianFilter := util.FilterByMedianFilterAlgorithm(signal)

	util.CreateXYPlot("i", "x(i)", indexes, signal, OUTPUT_3_NAME + "original")
	util.CreateXYPlot("i", "x(i)", indexes, newSignalByMovingAverage, OUTPUT_3_NAME + "moving_average")
	util.CreateXYPlot("i", "x(i)", indexes, newSignalByParabola, OUTPUT_3_NAME + "parabola_eleven")
	util.CreateXYPlot("i", "x(i)", indexes, newSignalByMedianFilter, OUTPUT_3_NAME + "median_filter")

	util.WriteWAVByMeta(OUTPUT_3_NAME + "original_sound", util.SOUND_LENGTH, signal, meta)
	util.WriteWAVByMeta(OUTPUT_3_NAME + "moving_average", util.SOUND_LENGTH, newSignalByMovingAverage, meta)
	util.WriteWAVByMeta(OUTPUT_3_NAME + "parabola_eleven", util.SOUND_LENGTH, newSignalByParabola, meta)
	util.WriteWAVByMeta(OUTPUT_3_NAME + "median_filter", util.SOUND_LENGTH, newSignalByMedianFilter, meta)
}

func PerformFourthOption() {
	_, signal, meta := ReadSignal("test.wav")
	util.WriteWAVByMeta(OUTPUT_4_NAME + "signal_sound", util.SOUND_LENGTH, signal, meta)
	/*
	newSignalByMovingAverage := util.FilterByMovingAverageAlgorithm(signal)
	amountRange, amplitudesRange, phasesRange := Fourier(newSignalByMovingAverage)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_4_NAME + "moving_average_")
	util.WriteWAVByMeta(OUTPUT_4_NAME + "moving_average_signal_sound", util.SOUND_LENGTH, newSignalByMovingAverage, meta)

	newSignalByParabola := util.FilterByParabolaAlgorithm(signal)
	amountRange, amplitudesRange, phasesRange = Fourier(newSignalByParabola)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_4_NAME + "parabola_")
	util.WriteWAVByMeta(OUTPUT_4_NAME + "parabola_signal_sound", util.SOUND_LENGTH, newSignalByParabola, meta)

	newSignalByMedianFilter := util.FilterByMedianFilterAlgorithm(signal)
	amountRange, amplitudesRange, phasesRange = Fourier(newSignalByMedianFilter)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_4_NAME + "median_")
	util.WriteWAVByMeta(OUTPUT_4_NAME + "median_signal_sound", util.SOUND_LENGTH, newSignalByMedianFilter, meta)
	*/
}

func showRanges(amountRange []float64, phasesRange []float64, amplitudesRange []float64, outputName string) {
	util.CreateXYPlotWithStyle(
		"k",
		"Phases",
		amountRange,
		phasesRange,
		outputName + "phases",
		"impulses",
	)
	util.CreateXYPlotWithStyle(
		"k",
		"Amplitudes",
		amountRange,
		amplitudesRange,
		outputName + "amplitudes",
		"impulses",
	)
}

func showDifferenceOriginAndRestoredSignal(originSignal []float64, restoredSignal []float64, amountRange []float64, plotName string) {
	util.CreateXYPlot("i", "x(i)", amountRange, originSignal, plotName + "original_signal")
	util.CreateXYPlot("i", "x(i)", amountRange, restoredSignal, plotName + "restored_signal")
}