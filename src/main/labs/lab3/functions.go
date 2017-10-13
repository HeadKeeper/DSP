package lab3

import (
	"main/util"
	"math"
	"main/types"
)

const (
	HARMONICS_AMOUNT int = 111
)

var (
	AMPLITUDE_VALUES = []float64 {1, 5, 7, 8, 9, 10, 17}
	PHI_VALUES = []float64{math.Pi / 6, math.Pi / 4, math.Pi / 3, math.Pi / 2, 3 * math.Pi / 4, math.Pi}
)

func CreateHarmonics() []types.Harmonic {
	var harmonics []types.Harmonic
	for idx := 0; idx < HARMONICS_AMOUNT; idx++ {
		harmonics = append(harmonics, types.Harmonic {
			Amplitude: util.GetRandomValue(AMPLITUDE_VALUES),
			Phi: util.GetRandomValue(PHI_VALUES),
		})
	}
	return harmonics
}

func CalculateAmplitudesAndPhi(amount int, harmonicNumber float64) (float64, float64, float64, float64) {
	getAmplitudeCos := createAmplitudeFunction(float64(amount), math.Cos)
	getAmplitudeSin := createAmplitudeFunction(float64(amount), math.Sin)

	amplitudeC := getAmplitudeCos(int(harmonicNumber))
	amplitudeS := getAmplitudeSin(int(harmonicNumber))

	amplitude := math.Sqrt(
		math.Pow(amplitudeC, 2) +
		math.Pow(amplitudeS, 2),
	)

	phi := math.Atan(amplitudeS / amplitudeC)

	return amplitudeC, amplitudeS, amplitude, phi
}


func CalculateAmplitudesAndPhiForCreatedSignal(harmonicNumber int, amount float64, values []float64) (float64, float64, float64, float64) {
	getAmplitudeCos := createAmplitudeFunctionForCreatedSignal(amount, math.Cos, values)
	getAmplitudeSin := createAmplitudeFunctionForCreatedSignal(amount, math.Sin, values)

	amplitudeC := getAmplitudeCos(harmonicNumber)
	amplitudeS := getAmplitudeSin(harmonicNumber)

	amplitude := math.Sqrt(
		math.Pow(amplitudeC, 2) +
			math.Pow(amplitudeS, 2),
	)

	phi := math.Atan(amplitudeS / amplitudeC)

	return amplitudeC, amplitudeS, amplitude, phi
}


func createAmplitudeFunction(amount float64, fun func(x float64) float64) func(harmonicNumber int) float64 {
	createTestSignal := CreateTestSignalFunction()
	var values []float64
	for index := 0; index < int(amount); index ++ {
		values = append(values, createTestSignal(float64(index), amount))
	}
	return createAmplitudeFunctionForCreatedSignal(amount, fun, values)
}

func createAmplitudeFunctionForCreatedSignal(amount float64, fun func(x float64) float64, values []float64) func(harmonicNumber int) float64 {
	return func(harmonicNumber int) float64 {
		var sum float64
		for idx := 0; idx < int(amount); idx++ {
			sum += values[idx] * fun(2 * math.Pi * float64(idx) * float64(harmonicNumber) / amount)
		}

		return sum * 2 / amount
	}
}

func getRanges(amount float64) ([]float64, []float64, []float64) {
	var phasesRange []float64
	var amplitudesRange []float64
	var amountRange []float64

	harmonics := CreateHarmonics()
	for index := range harmonics {
		/*amplitudeC, amplitudeS, */_, _, amplitude, phi := CalculateAmplitudesAndPhi(index, amount)

		amountRange = append(amountRange, float64(index+1))
		phasesRange = append(phasesRange, phi)
		amplitudesRange = append(amplitudesRange, amplitude)
	}

	return amountRange, phasesRange, amplitudesRange
}

func getRangesForCreatedSignal(amount float64, values []float64, harmonics []types.Harmonic) ([]float64, []float64, []float64) {
	var phasesRange []float64
	var amplitudesRange []float64
	var amountRange []float64

	for index := range harmonics {
		_, _, amplitude, phi := CalculateAmplitudesAndPhiForCreatedSignal(index, amount, values)

		amountRange = append(amountRange, float64(index+1))
		phasesRange = append(phasesRange, phi)
		amplitudesRange = append(amplitudesRange, amplitude)
	}

	return amountRange, phasesRange, amplitudesRange
}

func CreateTestSignalFunction() func (index float64, amount float64) float64 {
	return func(index float64, amount float64) float64 {
		return 50 * math.Cos(2 * math.Pi * index / amount - math.Pi / 3)
	}
}

func CreatePolyharmonicSignalFunction() func(index float64, harmonics []types.Harmonic) float64 {
	return func(index float64, harmonics []types.Harmonic) float64 {
		var sum float64
		var amount = float64(len(harmonics))

		for j := 0; j < len(harmonics); j ++ {
			sum += harmonics[j].Amplitude * math.Cos(2 * math.Pi * float64(j) * index / amount - harmonics[j].Phi)
		}
		return sum
	}
}

func getSignalValues(amount float64, harmonics []types.Harmonic, createSignal func(x float64, harmonics []types.Harmonic) float64) ([]float64, []float64) {
	var values []float64
	var indexes []float64

	for index := 0.0; index < amount; index+=1 {
		values = append(values, createSignal(index, harmonics))
		indexes = append(indexes, index)
	}

	return values, indexes
}

func RestoreSignalByRanges(amplitudeRanges []float64, phaseRanges []float64, amount int) func(x float64) float64 {
	return func(x float64) float64 {
		var sum float64
		for j := 0; j < amount / 2; j++ {
			sum += amplitudeRanges[j] * math.Cos( 2 * math.Pi * float64(j) * x / float64(amount) - phaseRanges[j])
		}
		return sum
	}
}

func RestorePolyharmonicSignalByRanges(amplitudeRanges []float64, phaseRanges []float64, amount int) func(x float64) float64  {
	return func(x float64) float64 {
		var sum float64
		for j := 0; j < amount / 2; j++ {
			sum += amplitudeRanges[j] * math.Cos( 2 * math.Pi * float64(j) * x / float64(amount) - phaseRanges[j])
		}
		return sum
	}
}