package lab3

import (
	"main/util"
	"math"
	"main/types"
	"math/cmplx"
)

const (
	HARMONICS_AMOUNT int = 10
)

var (
	AMPLITUDE_VALUES = []float64 {1, 5, 7, 8, 9, 10, 17}
	PHI_VALUES = []float64{math.Pi / 6, math.Pi / 4, math.Pi / 3, math.Pi / 2, 3 * math.Pi / 4, math.Pi}
)

func CreateTestSignalFunction() func (index float64) float64 {
	return func(index float64) float64 {
		return 50 * math.Cos(2 * math.Pi * index / util.BUFFER_SIZE - math.Pi / 3)
	}
}

func CreatePolyharmonicSignalFunction(harmonics []types.Harmonic) func(index float64) float64 {
	return func(index float64) float64 {
		var sum float64
		for j := 0; j < len(harmonics); j ++ {
			sum += harmonics[j].Amplitude * math.Cos(2 * math.Pi * float64(j) * index / util.BUFFER_SIZE - harmonics[j].Phi)
		}
		return sum
	}
}

func CreateSignal(signalFunction func(x float64) float64) []float64 {
	var signal []float64
	for idx := 0.0; idx < util.BUFFER_SIZE; idx += 1 {
		signal = append(signal, signalFunction(idx))
	}
	return signal
}

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

func Fourier(signal []float64) ([]float64, []float64, []float64) {
	length := len(signal)
	var beginPhases []float64
	var amplitudes []float64
	var indexes []float64

	for i := 0; i < length; i++ {
		amplitudeS := 0.0
		amplitudeC := 0.0
		for j := 0; j < length; j++ {
			w := 2 * math.Pi * float64(j) * float64(i) / float64(length)

			amplitudeS += signal[j] * math.Sin(w)
			amplitudeC += signal[j] * math.Cos(w)

		}
		amplitudeS *= 2 / float64(length)
		amplitudeC *= 2 / float64(length)

		beginPhases = append(beginPhases, math.Atan2(amplitudeS, amplitudeC))
		amplitudes = append(amplitudes, math.Hypot(amplitudeS, amplitudeC))
		indexes = append(indexes, float64(i))
	}

	return indexes, amplitudes, beginPhases
}

func FastFourierTransform(signal []float64, complexSignal []complex128, s int) []complex128 {
	//complexSignal := make([]complex128, len(signal))
	bufferSize := len(signal)

	FastFourierTransform(signal, complexSignal, 2 * s)
	FastFourierTransform(signal[s:], complexSignal, 2 * s)

	for k := 0; k < bufferSize / 2; k ++ {
		tf := cmplx.Rect(1, -2 * math.Pi * float64(k) / float64(bufferSize)) * complexSignal[k + bufferSize / 2]
		complexSignal[k], complexSignal[k + bufferSize / 2] = complexSignal[k] + tf, complexSignal[k] - tf
	}

	return complexSignal
}

func RestoreSignal(amplitudeRanges []float64, phaseRanges []float64) []float64 {
	amount := int(util.BUFFER_SIZE)
	var signal []float64

	for i := 0; i < amount; i ++ {
		var sum float64
		for j := 0; j < amount / 2; j++ {
			sum += amplitudeRanges[j] * math.Cos( 2 * math.Pi * float64(j) * float64(i) / float64(amount) - phaseRanges[j])
		}
		signal = append(signal, sum)
	}
	return signal
}

func RestorePolyharmonicSignal(amplitudeRanges []float64, phaseRanges []float64) []float64  {
	amount := int(util.BUFFER_SIZE)
	var signal []float64

	for i := 0; i < amount; i ++ {
		var sum float64
		for j := 1; j < amount / 2 - 1; j++ {
			sum += amplitudeRanges[j] * math.Cos( 2 * math.Pi * float64(j) * float64(i) / float64(amount) - phaseRanges[j])
		}
		sum += amplitudeRanges[0] / 2
		signal = append(signal, sum)
	}
	return signal
}