package lab1

import (
	"math"
	"main/types"
	"main/util"
)

const (
	FIRST_OPTION__a__f  float64 = 4   		    // frequency
	FIRST_OPTION__a__A  float64 = 9   		    // amplitude of fluctuations

	FIRST_OPTION__b__A   float64 = 7
	FIRST_OPTION__b__Phi float64 = math.Pi / 6 // initial angle

	FIRST_OPTION__c__f   float64 = 7
	FIRST_OPTION__c__Phi float64 = math.Pi / 6

	THIRD_OPTION__INC_MAX_VALUE float64 = 0.20
)

var (
	PHI_VALUES = []float64 {
		math.Pi / 3,
		3 * math.Pi / 4,
		2 * math.Pi,
		math.Pi,
		math.Pi / 6,
	}
	FREQUENCY_VALUES = []float64 {
		4,
		8,
		2,
		1,
		9,
	}
	AMPLITUDE_VALUES = []float64 {
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

func createFirstFunctionA(phi float64) func(n float64) float64 {
	return util.GetHarmonicFunction(FIRST_OPTION__a__A, FIRST_OPTION__a__f, phi)
}

func createFirstFunctionB(frequency float64) func(x float64) float64 {
	return util.GetHarmonicFunction(FIRST_OPTION__b__A, frequency, FIRST_OPTION__b__Phi)
}

func createFirstFunctionC(amplitude float64) func(x float64) float64 {
	return util.GetHarmonicFunction(amplitude, FIRST_OPTION__c__f, FIRST_OPTION__c__Phi)
}


func createSecondFunction() func(n float64) float64 {
	return func(n float64) float64 {
		var result float64 = 0
		for _, harmonic := range SECOND_OPTION_HARMONICS {
			result += util.GetHarmonicFunction(harmonic.Amplitude, harmonic.Frequency, harmonic.Phi)(n)
		}
		return result
	}
}


func createThirdFunction(cyclesCount int) func(x float64) float64 {
	var coefficient float64 = float64(THIRD_OPTION__INC_MAX_VALUE) / float64(cyclesCount)
	harmonic := THIRD_OPTION__INIT_HARMONIC
	return func(n float64) float64 {
		var value float64

		for idx := 0; idx < cyclesCount; idx++ {
			value += util.GetHarmonicFunction(harmonic.Amplitude, harmonic.Frequency, harmonic.Phi)(n)
		}

		harmonic.Phi *= 1.0 + coefficient
		harmonic.Amplitude *= 1.0 + coefficient
		harmonic.Frequency *= 1.0 + coefficient
		return value
	}
}