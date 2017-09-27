package lab1

import (
	"math"
	"main/types"
)

const (
	// w = 2 * Pi * f - angular speed of rotation
	DEFAULT_N          float64 = 512 //

	FIRST_OPTION__a__f float64 = 4   		   // frequency
	FIRST_OPTION__a__A float64 = 9   		   // amplitude of fluctuations

	FIRST_OPTION__b__A float64 = 7
	FIRST_OPTION__b__y float64 = math.Pi / 6   // initial angular

	FIRST_OPTION__c__f float64 = 7
	FIRST_OPTION__c__y float64 = math.Pi / 6
)

var (
	FIRST_OPTION__a__y__VALUES = []float64 {
		math.Pi / 3,
		3 * math.Pi / 4,
		2 * math.Pi,
		math.Pi,
		math.Pi / 6,
	}
	FIRST_OPTION__b__f__VALUES = []float64 {
		4,
		8,
		2,
		1,
		9,
	}
	FIRST_OPTION__c__A__VALUES = []float64 {
		4,
		5,
		3,
		1,
		7,
	}

	SECOND_OPTION_HARMONICS    = []types.Harmonic {
		{
			A:       7,
			F_small: 1,
			Y_small: math.Pi,
		},
		{
			A:       7,
			F_small: 2,
			Y_small: math.Pi / 4,
		},
		{
			A:       7,
			F_small: 3,
			Y_small: 0,
		},
		{
			A:       7,
			F_small: 4,
			Y_small: 3 * math.Pi / 4,
		},
		{
			A:       7,
			F_small: 5,
			Y_small: math.Pi / 2,
		},
	}
)


func createFirstFormulaA(y float64, soundLength float64) func(x float64) float64 {
	return func(x float64) float64 {
		return FIRST_OPTION__a__A * math.Sin( ((2 * math.Pi * FIRST_OPTION__a__f * x) / DEFAULT_N) * soundLength + y ) }
}

func createFirstFormulaB(f float64, soundLength float64) func(x float64) float64 {
	return func(x float64) float64 {
		return FIRST_OPTION__b__A * math.Sin( ((2 * math.Pi * f * x) / DEFAULT_N) * soundLength + FIRST_OPTION__b__y ) }
}

func createFirstFormulaC(A float64, soundLength float64) func(x float64) float64 {
	return func(x float64) float64 {
		return A * math.Sin( ((2 * math.Pi * FIRST_OPTION__c__f * x) / DEFAULT_N) * soundLength + FIRST_OPTION__c__y )}
}


func createSecondFormula(y float64, soundLength float64) func(x float64) float64 {
	return func(x float64) float64 {
		var result float64 = 0
		for _, element := range SECOND_OPTION_HARMONICS {
			result += element.A * math.Sin( ((2 * math.Pi * element.F_small * x) / DEFAULT_N) * soundLength + element.Y_small + y)
		}
		return result
	}
}
