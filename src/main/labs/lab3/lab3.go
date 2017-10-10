package lab3

import (
	"main/util"
)

const (
	OUTPUT_NAME     string = "lab3_out_"
	OUTPUT_1_A_NAME string = OUTPUT_NAME + "1_a_range_"
	OUTPUT_1_B_NAME string = OUTPUT_NAME + "1_b_"
)

func PerformFirstOption() {
	amountRange, phasesRange, amplitudesRange := getRanges()
	showRanges(amountRange, phasesRange, amplitudesRange)
	showDifferenceOriginAndRestoredSignal(amountRange, phasesRange, amplitudesRange)
}

func getRanges() ([]float64, []float64, []float64) {
	var phasesRange []float64
	var amplitudesRange []float64
	var amountRange []float64

	harmonics := CreateHarmonicsForFirstOption()
	for index, harmonic := range harmonics {
		/*amplitudeC, amplitudeS, */_, _, amplitude, phi := CalculateAmplitudesAndPhi(harmonic, float64(index))

		amountRange = append(amountRange, float64(index+1))
		phasesRange = append(phasesRange, phi)
		amplitudesRange = append(amplitudesRange, amplitude)
	}

	return amountRange, phasesRange, amplitudesRange
}

func showRanges(amountRange []float64, phasesRange []float64, amplitudesRange []float64) {
	util.CreateXYPlotWithStyle("k", "Phases", amountRange, phasesRange, OUTPUT_1_A_NAME+ "phases", "impulses")
	util.CreateXYPlotWithStyle("k", "Amplitudes", amountRange, amplitudesRange, OUTPUT_1_A_NAME+ "amplitudes", "impulses")
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