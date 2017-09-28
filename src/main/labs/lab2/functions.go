package lab2

import (
	"main/util"
)

const K float64 = util.SIGNAL_RATE / 4

func getHarmonicFunction(soundLength float64) func(n float64) float64 {
	return util.GetHarmonicFunction(1, 1, soundLength, 0)
}

func createMArray() []float64 {
	var currentM float64
	var arrayM []float64
	for currentM = K; currentM < util.SIGNAL_RATE * 2; currentM+=12 {
		arrayM = append(arrayM, currentM - 1)
	}

	return arrayM
}

func createFunction(soundLength float64) func(n float64) float64 {
	return func(n float64) float64 {
		return getHarmonicFunction(soundLength)(n)
	}
}
