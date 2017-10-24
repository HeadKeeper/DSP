package lab2

import (
	"main/util"
	"main/types"
	"math"
)

const K float64 = util.SIGNAL_RATE / 4

func getHarmonicFunction(phi float64) func(n float64) float64 {
	return util.GetHarmonicFunction(2000000000, 440, phi)
}

func createMArray() []float64 {
	var currentM float64
	var arrayM []float64
	for currentM = util.SIGNAL_RATE * 2; currentM >= K; currentM-=12 {
		arrayM = append(arrayM, currentM + 1)
	}

	return arrayM
}

func createFunction() func(n float64) float64 {
	return func(n float64) float64 {
		return getHarmonicFunction(0)(n)
	}
}

func createSecondFunction(phi float64) func(n float64) float64 {
	return func(n float64) float64 {
		return getHarmonicFunction(phi)(n)
	}
}

func CalculateRootMeanSquareValues(data types.PlotData) (float64, float64, float64, float64) {
	first := calculateRootMeanSquare(data, createFirstRootMeanSquareFunction())
	firstInaccuracy := calculateRMSInaccuracy(first)
	second := calculateRootMeanSquare(data, createSecondRootMeanSquareFunction())
	secondInaccuracy := calculateRMSInaccuracy(second)

	return first, firstInaccuracy, second, secondInaccuracy
}

func getFNValues(data types.PlotData) []float64 {
	var fnValues []float64
	for _, currentN := range util.GetValues(data.InitialN, data.EndN, data.Step) {
		fnValues = append(fnValues, data.Function(currentN))
	}

	return fnValues
}

func calculateRootMeanSquare(data types.PlotData, rootMeanSquareFunction func(values []float64) float64) float64 {
	return rootMeanSquareFunction(getFNValues(data))
}

func createFirstRootMeanSquareFunction() func(values []float64) float64 {
	return func(values []float64) float64 {
		var sum float64
		for _, value := range values {
			sum += value * value
		}
		return math.Sqrt(sum / float64(len(values) + 1))
	}
}

func createSecondRootMeanSquareFunction() func(values []float64) float64 {
	return func(values []float64) float64 {
		var sum float64
		for _, value := range values {
			sum += value
		}
		return math.Sqrt(createFirstRootMeanSquareFunction()(values) - math.Pow(sum / float64(len(values) + 1), 2))
	}
}


func CalculateAmplitudeValue(data types.PlotData) (float64, float64) {
	amplitude := math.Sqrt(
		math.Pow(createAmplitudeFunction(math.Sin)(getFNValues(data)), 2) +
			math.Pow(createAmplitudeFunction(math.Cos)(getFNValues(data)), 2),
	)
	inaccuracy := calculateAmplitudeInaccuracy(amplitude)

	return amplitude, inaccuracy
}

func createAmplitudeFunction(fun func(x float64) float64) func(values []float64) float64 {
	return func(values []float64) float64 {
		var sum float64
		for idx, value := range values {
			sum += value * fun(2 * math.Pi * float64(idx) / float64(len(values)))
		}

		return 2 / float64(len(values)) * sum
	}
}


func calculateRMSInaccuracy(RMS float64) float64 {
	return math.Abs(0.707 - RMS)
}

func calculateAmplitudeInaccuracy(amplitude float64) float64 {
	return math.Abs(1 - amplitude)
}