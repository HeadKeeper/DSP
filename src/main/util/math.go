package util

import (
	"math"
)

const (
	// w = 2 * Pi * f - angular speed of rotation
	// phase = (2 * Pi * f * x) / N) * soundLength + initialAngle
	SIGNAL_RATE float64 = 512
	SOUND_LENGTH = 15
	BUFFER_SIZE float64 = SIGNAL_RATE
)

func GetHarmonicFunction(amplitude float64, frequency float64, phi float64) func(n float64) float64 {
	return func(n float64) float64 {
		return amplitude * math.Sin(2 * math.Pi * frequency * n / BUFFER_SIZE + phi) }
}


func GetValues(initialN float64, endN float64, step float64) []float64 {
	var n float64
	var values []float64
	for n = initialN; n < endN; n += step {
		values = append(values, n)
	}

	return values
}