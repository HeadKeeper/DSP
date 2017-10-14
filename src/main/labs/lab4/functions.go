package lab4

import (
	"math"
	"main/util"
	"main/types"
	"math/cmplx"
)

const (
	B1 = 10000.0
	B2 = 1.0
	HARMONICS_AMOUNT = 20
)

var (
	EPSYLON_VALUES = []float64 { 1.0, 0.0 }
)

func CreateSignalFunction() func(x float64) float64 {
	return func(x float64) float64 {
		result := B1 * math.Sin(2 * math.Pi * x / util.BUFFER_SIZE)
		var sum float64
		for index := 50; index < 70; index++ {
			sum += math.Pow(-1, util.GetRandomValue(EPSYLON_VALUES)) * B2 * math.Sin(2 * math.Pi * x * float64(index) / util.BUFFER_SIZE)
		}

		return result + sum
	}
}

func CreateSignal(signalFunction func(x float64) float64) ([]float64, []float64) {
	var signal []float64
	var indexes []float64
	for idx := 0.0; idx < util.BUFFER_SIZE; idx += 1 {
		signal = append(signal, signalFunction(idx))
		indexes = append(indexes, idx)
	}
	return indexes, signal
}

func CreateHarmonics() []types.Harmonic {
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