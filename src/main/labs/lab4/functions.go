package lab4

import (
	"math"
	"main/util"
	"main/types"
	"sort"
)

const (
	B1 = 10000.0
	B2 = 1.0
	HARMONICS_AMOUNT = 20

	MOVING_AVERAGE__K = 5
	MEDIAN_FILTER__N  = 9
)

var (
	EPSYLON_VALUES = []float64 { 1.0, 0.0 }
)

func CreateSignalFunction(amount float64) func(x float64) float64 {
	return func(x float64) float64 {
		result := B1 * math.Sin(2 * math.Pi * x / amount)
		var sum float64
		for index := 50; index < 70; index++ {
			sum += math.Pow(-1, util.GetRandomValue(EPSYLON_VALUES)) * B2 * math.Sin(2 * math.Pi * x * float64(index) / amount)
		}

		return result + sum
	}
}

func getSignalValues(createSignal func(x float64) float64) ([]float64, []float64) {
	var values []float64
	var indexes []float64

	for index := 0.0; index < AMOUNT; index+=1 {
		values = append(values, createSignal(index))
		indexes = append(indexes, index)
	}

	return values, indexes
}

func CreateHarmonicsForSignal() []types.Harmonic {
	var harmonics []types.Harmonic
	initialHarmonic := types.Harmonic {
		Amplitude: B1,
	}

	harmonics = append(harmonics, initialHarmonic)

	for idx := 0; idx < HARMONICS_AMOUNT; idx++ {
		harmonics = append(harmonics, types.Harmonic {
			Amplitude: B2,
		})
	}
	return harmonics
}

func CalculateAmplitudesAndPhi(harmonicNumber int, amount float64) (float64, float64, float64, float64) {
	getAmplitudeCos := createAmplitudeFunction(amount, math.Cos)
	getAmplitudeSin := createAmplitudeFunction(amount, math.Sin)

	amplitudeC := getAmplitudeCos(harmonicNumber)
	amplitudeS := getAmplitudeSin(harmonicNumber)

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
	createTestSignal := CreateSignalFunction(amount)
	var values []float64
	for index := 0; index < int(amount); index ++ {
		values = append(values, createTestSignal(float64(index)))
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

	harmonics := CreateHarmonicsForSignal()
	for index := range harmonics {
		/*amplitudeC, amplitudeS, */_, _, amplitude, phi := CalculateAmplitudesAndPhi(index, amount)

		amountRange = append(amountRange, float64(index+1))
		phasesRange = append(phasesRange, phi)
		amplitudesRange = append(amplitudesRange, amplitude)
	}

	return amountRange, phasesRange, amplitudesRange
}

func getRangesForCreatedSignal(amount float64, values []float64) ([]float64, []float64, []float64) {
	var phasesRange []float64
	var amplitudesRange []float64
	var amountRange []float64

	harmonics := CreateHarmonicsForSignal()
	for index := range harmonics {
		_, _, amplitude, phi := CalculateAmplitudesAndPhiForCreatedSignal(index, amount, values)

		amountRange = append(amountRange, float64(index+1))
		phasesRange = append(phasesRange, phi)
		amplitudesRange = append(amplitudesRange, amplitude)
	}

	return amountRange, phasesRange, amplitudesRange
}

func filterByMovingAverageAlgorithm(values []float64) []float64 {
	var newValues []float64

	m := int((MOVING_AVERAGE__K - 1) / 2)

	for index := range values {
		newValues = append(newValues, getXByMovingAverage(index, m, MOVING_AVERAGE__K, values))
	}

	return newValues
}

func getXByMovingAverage(i int, m int, K int, values []float64) float64 {
	var sum float64
	startEdge := i - m
	endEdge := i + m

	if startEdge < 0 {
		startEdge = 0
	}
	if endEdge >= len(values) {
		endEdge = len(values)
	}

	for index := startEdge; index < endEdge; index++ {
		sum += values[index]
	}

	return sum / float64(K)
}

func filterByParabolaAlgorithm(values []float64) []float64 {
	var newValues []float64

	for index := range values {
		newValues = append(newValues, getXByParabolaEleven(index, values))
	}

	return newValues
}

func getXByParabolaEleven(i int, values []float64) float64 {
	var sum float64
	var sumCoeff float64

	length := len(values)

	if i - 5 > 0 {
		sum += 18 * values[i-5]
		sumCoeff += 18
	}
	if i - 4 > 0 {
		sum -= 45 * values[i-4]
		sumCoeff -= 45
	}
	if i - 3 > 0 {
		sum -= 10 * values[i-3]
		sumCoeff -= 10
	}
	if i - 2 > 0 {
		sum += 60 * values[i-2]
		sumCoeff += 60
	}
	if i - 1 > 0 {
		sum += 120 * values[i-1]
		sumCoeff += 120
	}
	sum += 143 * values[i]
	sumCoeff += 143
	if i + 1 < length {
		sum += 120 * values[i+1]
		sumCoeff += 120
	}
	if i + 2 < length {
		sum += 60 * values[i+2]
		sumCoeff += 60
	}
	if i + 3 < length {
		sum -= 10 * values[i+3]
		sumCoeff -= 10
	}
	if i + 4 < length {
		sum -= 45 * values[i+4]
		sumCoeff -= 45
	}
	if i + 5 < length {
		sum += 18 * values[i+5]
		sumCoeff += 18
	}

	return sum / sumCoeff
}

func filterByMedianFilterAlgorithm(values []float64) []float64 {
	var newValues []float64

	for index := range values {
		newValues = append(newValues, getXByMedianFilter(index, values))
	}

	return newValues
}

func getXByMedianFilter(i int, values []float64) float64 {
	n := int(MEDIAN_FILTER__N / 2)

	startEdge := i - n
	endEdge := i + n

	if startEdge < 0 {
		startEdge = 0
	}
	if endEdge >= len(values) {
		endEdge = len(values)
	}

	var currentWindow []float64
	for index := startEdge; index < endEdge; index ++ {
		currentWindow = append(currentWindow, values[index])
	}

	sort.Float64s(currentWindow)
	targetIndex := int(len(currentWindow) / 2)

	return currentWindow[targetIndex]
}
