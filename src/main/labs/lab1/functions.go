package lab1

import (
	"math"
	"main/types"
	"main/util"
)

const (
	FIRST_OPTION__a__f  float64 = 440   		    // frequency
	FIRST_OPTION__a__A  float64 = 2000000000   		// amplitude of fluctuations

	FIRST_OPTION__b__A   float64 = 2000000000
	FIRST_OPTION__b__Phi float64 = math.Pi / 6 // initial angle

	FIRST_OPTION__c__f   float64 = 440
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
		440,
		880,
		220,
		1760,
		1000,
	}
	AMPLITUDE_VALUES = []float64 {
		2000000000,
		1500000000,
		2100000000,
		1000000000,
		1750000000,
	}

	SECOND_OPTION_HARMONICS     = []types.Harmonic {
		{
			Amplitude: 2000000000,
			Frequency: 440,
			Phi:       math.Pi,
		},
		{
			Amplitude: 2000000000,
			Frequency: 220,
			Phi:       math.Pi / 4,
		},
		{
			Amplitude: 2000000000,
			Frequency: 880,
			Phi:       0,
		},
		{
			Amplitude: 2000000000,
			Frequency: 1760,
			Phi:       3 * math.Pi / 4,
		},
		{
			Amplitude: 2000000000,
			Frequency: 1000,
			Phi:       math.Pi / 2,
		},
	}

	THIRD_OPTION__INIT_HARMONIC = types.Harmonic{
		Amplitude: 2000000000,
		Frequency: 220,
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

		//for idx := 0; idx < cyclesCount; idx++ {
			value += util.GetHarmonicFunction(harmonic.Amplitude, harmonic.Frequency, harmonic.Phi)(n)
		//}

		harmonic.Phi *= 1.0 + coefficient / 100
		harmonic.Amplitude *= 1.0 + coefficient / 100
		harmonic.Frequency *= 1.0 + coefficient / 100
		return value
	}
}