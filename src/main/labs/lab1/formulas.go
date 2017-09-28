package lab1

import (
	"math"
	"main/types"
)

const (
	// w = 2 * Pi * f - angular speed of rotation
	// phase = (2 * Pi * f * x) / N) * soundLength + initialAngle

	SIGNAL_RATE		    float64 = 512 			// signal rate

	FIRST_OPTION__a__f  float64 = 4   		    // frequency
	FIRST_OPTION__a__A  float64 = 9   		    // amplitude of fluctuations

	FIRST_OPTION__b__A  float64 = 7
	FIRST_OPTION__b__y  float64 = math.Pi / 6   // initial angle

	FIRST_OPTION__c__f  float64 = 7
	FIRST_OPTION__c__y  float64 = math.Pi / 6

	THIRD_OPTION__INC_MAX_VALUE float64 = 0.20
)

var (
	FIRST_OPTION__a__y__VALUES  = []float64 {
		math.Pi / 3,
		3 * math.Pi / 4,
		2 * math.Pi,
		math.Pi,
		math.Pi / 6,
	}
	FIRST_OPTION__b__f__VALUES  = []float64 {
		4,
		8,
		2,
		1,
		9,
	}
	FIRST_OPTION__c__A__VALUES  = []float64 {
		4,
		5,
		3,
		1,
		7,
	}

	SECOND_OPTION_HARMONICS     = []types.Harmonic {
		{
			Amplitude: 7,
			Frequency: 1,
			Phi:       math.Pi,
		},
		{
			Amplitude: 7,
			Frequency: 2,
			Phi:       math.Pi / 4,
		},
		{
			Amplitude: 7,
			Frequency: 3,
			Phi:       0,
		},
		{
			Amplitude: 7,
			Frequency: 4,
			Phi:       3 * math.Pi / 4,
		},
		{
			Amplitude: 7,
			Frequency: 5,
			Phi:       math.Pi / 2,
		},
	}

	THIRD_OPTION__INIT_HARMONIC = types.Harmonic{
		Amplitude: 7,
		Frequency: 4,
		Phi:       math.Pi / 6,
	}
)

func getHarmonicFunction(amplitude float64, frequency float64, soundLength float64, phi float64) func(n float64) float64 {
	return func(n float64) float64 {
		return amplitude * math.Sin(2 * math.Pi * frequency * n / SIGNAL_RATE * soundLength + phi) }
}


func createFirstFunctionA(phi float64, soundLength float64) func(n float64) float64 {
	return getHarmonicFunction(FIRST_OPTION__a__A, FIRST_OPTION__a__f, soundLength, phi)
}

func createFirstFunctionB(frequency float64, soundLength float64) func(x float64) float64 {
	return getHarmonicFunction(FIRST_OPTION__b__A, frequency, soundLength, FIRST_OPTION__b__y)
}

func createFirstFunctionC(amplitude float64, soundLength float64) func(x float64) float64 {
	return getHarmonicFunction(amplitude, FIRST_OPTION__c__f, soundLength, FIRST_OPTION__c__y)
}


func createSecondFunction(additionalPhi float64, soundLength float64) func(n float64) float64 {
	return func(n float64) float64 {
		var result float64 = 0
		for _, harmonic := range SECOND_OPTION_HARMONICS {
			result += harmonic.Amplitude * math.Sin(2 * math.Pi * harmonic.Frequency * n / SIGNAL_RATE * soundLength + harmonic.Phi + additionalPhi)
		}
		return result
	}
}


func createThirdFunction(harmonic types.Harmonic, cyclesCount int, soundLength float64) func(x float64) float64 {
	var coefficient float64 = float64(THIRD_OPTION__INC_MAX_VALUE) / float64(cyclesCount)

	return func(n float64) float64 {
		var value float64
		if n == -10 {
			harmonic = THIRD_OPTION__INIT_HARMONIC
		}

		for idx := 0; idx < cyclesCount; idx++ {
			value += harmonic.Amplitude * math.Sin(2 * math.Pi * harmonic.Frequency * n / SIGNAL_RATE * soundLength + harmonic.Phi)
		}

		harmonic.Phi = harmonic.Phi * (1.0 + coefficient)
		harmonic.Amplitude = harmonic.Amplitude * (1.0 + coefficient)
		harmonic.Frequency = harmonic.Frequency * (1.0 + coefficient)
		return value
	}
}