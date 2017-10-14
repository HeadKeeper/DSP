package lab3

import (
	"main/util"
)

const (
	OUTPUT_NAME     string = "lab3_"
	OUTPUT_1_A_NAME string = OUTPUT_NAME + "1_a_"
	OUTPUT_1_B_NAME string = OUTPUT_NAME + "1_b_"
	OUTPUT_2_A_NAME string = OUTPUT_NAME + "2_a_"
	OUTPUT_2_B_NAME string = OUTPUT_NAME + "2_b_"
)

func PerformFirstOption() {
	signal := CreateSignal(CreateTestSignalFunction())
	amountRange, amplitudesRange, phasesRange := Fourier(signal)

	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_1_A_NAME)

	restoredSignal := RestoreSignal(amplitudesRange, phasesRange)
	showDifferenceOriginAndRestoredSignal(signal, restoredSignal, amountRange, OUTPUT_1_B_NAME)
}

func PerformSecondOption() {
	harmonics := CreateHarmonics()
	signal := CreateSignal(CreatePolyharmonicSignalFunction(harmonics))
	amountRange, amplitudesRange, phasesRange := Fourier(signal)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_A_NAME)

	restoredSignal := RestorePolyharmonicSignal(amplitudesRange, phasesRange)
	showDifferenceOriginAndRestoredSignal(signal, restoredSignal, amountRange, OUTPUT_2_B_NAME)

	amplitudesRange, phasesRange = util.FilterValues(util.LOW_PASS_FILTER, amplitudesRange, phasesRange, []float64{11})
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_A_NAME + "filtered_by_low_pass_")

	amplitudesRange, phasesRange = util.FilterValues(util.HIGH_PASS_FILTER, amplitudesRange, phasesRange, []float64{6})
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_A_NAME + "filtered_by_high_pass_")

	amplitudesRange, phasesRange = util.FilterValues(util.BAND_PASS_FILTER, amplitudesRange, phasesRange, []float64{11, 6,  35, 45})
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_A_NAME + "filtered_by_band_pass_")

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
