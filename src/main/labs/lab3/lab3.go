package lab3

import (
	"main/util"
)

const (
	OUTPUT_NAME     string = "lab3_out_"
	OUTPUT_1_A_NAME string = OUTPUT_NAME + "1_a_range_"
	OUTPUT_1_B_NAME string = OUTPUT_NAME + "1_b_"
	OUTPUT_2_A_NAME string = OUTPUT_NAME + "2_a_"
	OUTPUT_2_B_NAME string = OUTPUT_NAME + "2_b_"

	AMOUNT = util.SIGNAL_RATE * util.SOUND_LENGTH
)

func PerformFirstOption() {
	amountRange, phasesRange, amplitudesRange := getRanges(AMOUNT)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_1_A_NAME)
	showDifferenceOriginAndRestoredSignal(amountRange, phasesRange, amplitudesRange)
}

func showRanges(amountRange []float64, phasesRange []float64, amplitudesRange []float64, outputName string) {
	util.CreateXYPlotWithStyle("k", "Phases", amountRange, phasesRange, outputName + "phases", "impulses")
	util.CreateXYPlotWithStyle("k", "Amplitudes", amountRange, amplitudesRange, outputName + "amplitudes", "impulses")
}

func showDifferenceOriginAndRestoredSignal(amountRange []float64, phasesRange []float64, amplitudesRange []float64) {
	createTestSignal := CreateTestSignalFunction()
	restoreSignal := RestoreSignalByRanges(amplitudesRange, phasesRange, len(amountRange))

	var testSignalOrigin []float64
	var testSignalRestored []float64
	for index := 0; index < len(amountRange); index ++ {
		testSignalOrigin = append(testSignalOrigin, createTestSignal(float64(index), float64(len(amountRange))))
		testSignalRestored = append(testSignalRestored, restoreSignal(float64(index)))
	}

	util.CreateXYPlot("i", "x(i)", amountRange, testSignalOrigin, OUTPUT_1_B_NAME + "original_signal")
	util.CreateXYPlot("i", "x(i)", amountRange, testSignalRestored, OUTPUT_1_B_NAME + "restored_signal")
}

func PerformSecondOption() {
	harmonics := CreateHarmonics()
	createSignal := CreatePolyharmonicSignalFunction()

	values, _ := getSignalValues(AMOUNT, harmonics, createSignal)
	amountRange, phasesRange, amplitudesRange := getRangesForCreatedSignal(AMOUNT, values, harmonics)
	showRanges(amountRange, phasesRange, amplitudesRange, OUTPUT_2_A_NAME)

	restoreSignal := RestorePolyharmonicSignalByRanges(amplitudesRange, phasesRange, len(amountRange))

	var testSignalOrigin []float64
	var testSignalRestored []float64
	for index := 0; index < len(amountRange); index ++ {
		testSignalOrigin = append(testSignalOrigin, createSignal(float64(index), harmonics))
		testSignalRestored = append(testSignalRestored, restoreSignal(float64(index)))
	}

	util.CreateXYPlot("i", "x(i)", amountRange, testSignalOrigin, OUTPUT_2_B_NAME + "original_signal")
	util.CreateXYPlot("i", "x(i)", amountRange, testSignalRestored, OUTPUT_2_B_NAME + "restored_signal")
}