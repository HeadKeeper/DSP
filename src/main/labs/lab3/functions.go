package lab3

import (
	"main/util"
	"math"
	"main/types"
)

const (
	HARMONICS_AMOUNT int = 110
)

var (
	AMPLITUDE_VALUES = []float64 {1, 5, 7, 8, 9, 10, 17}
	PHI_VALUES = []float64{math.Pi / 6, math.Pi / 4, math.Pi / 3, math.Pi / 2, 3 * math.Pi / 4, math.Pi}
)

func CreateHarmonicsForFirstOption() []types.Harmonic {
	var harmonics []types.Harmonic
	for idx := 0; idx < HARMONICS_AMOUNT; idx++ {
		harmonics = append(harmonics, types.Harmonic {
			Amplitude: util.GetRandomValue(AMPLITUDE_VALUES),
			Phi: util.GetRandomValue(PHI_VALUES),
		})
	}
	return harmonics
}

func CalculateAmplitudesAndPhi(harmonic types.Harmonic, harmonicNumber float64) (float64, float64, float64, float64) {
	getAmplitudeCos := createAmplitudeFunction(math.Cos)
	getAmplitudeSin := createAmplitudeFunction(math.Sin)

	amplitudeC := getAmplitudeCos(util.SIGNAL_RATE * util.SOUND_LENGTH, harmonicNumber)
	amplitudeS := getAmplitudeSin(util.SIGNAL_RATE * util.SOUND_LENGTH, harmonicNumber)

	amplitude := math.Sqrt(
		math.Pow(amplitudeC, 2) +
		math.Pow(amplitudeS, 2),
	)

	phi := math.Atan(amplitudeS / amplitudeC)

	return amplitudeC, amplitudeS, amplitude, phi
}

func createAmplitudeFunction(fun func(x float64) float64) func(amount float64, harmonicNumber float64) float64 {
	createTestSignal := CreateTestSignalFunction()
	return func(amount float64, harmonicNumber float64) float64 {
		var sum float64
		for idx := 0; idx < int(amount); idx++ {
			sum += createTestSignal(float64(idx), amount) * fun(2 * math.Pi * float64(idx) * harmonicNumber / amount)
		}

		return sum * 2 / amount
	}
}

func CreateTestSignalFunction() func (index float64, amount float64) float64 {
	return func(index float64, amount float64) float64 {
		return 50 * math.Cos(2 * math.Pi * index / amount - math.Pi / 3)
	}
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
