package util

import "sort"

const (
	LOW_PASS_FILTER = "LOW_PASS"
	HIGH_PASS_FILTER = "HIGH_PASS"
	BAND_PASS_FILTER = "BAND_PASS"

	MOVING_AVERAGE__K = 5
	MEDIAN_FILTER__N  = 9
)

func FilterValues(filterType string, amplitudes []float64, beginPhases []float64, borders []float64) ([]float64, []float64) {
	var filteredAmplitudes []float64
	var filteredBeginPhases []float64

	switch filterType {
	case LOW_PASS_FILTER:
		filter := lowPassFilter(borders[0])
		for _, v := range amplitudes {
			if filter(v) {
				filteredAmplitudes = append(filteredAmplitudes, v)
			}
		}
		for _, v := range beginPhases {
			if filter(v) {
				filteredBeginPhases = append(filteredBeginPhases, v)
			}
		}
		break
	case HIGH_PASS_FILTER:
		filter := highPassFilter(borders[0])
		for _, v := range amplitudes {
			if filter(v) {
				filteredAmplitudes = append(filteredAmplitudes, v)
			}
		}
		for _, v := range beginPhases {
			if filter(v) {
				filteredBeginPhases = append(filteredBeginPhases, v)
			}
		}
		break
	case BAND_PASS_FILTER:
		if len(borders) == 1 {
			break
		}
		for index := 0; index < len(borders) / 2 - 1; index ++ {
			filter := bandPassFilter(borders[index], borders[index + 1])
			for _, v := range amplitudes {
				if filter(v) {
					filteredAmplitudes = append(filteredAmplitudes, v)
				}
			}
			for _, v := range beginPhases {
				if filter(v) {
					filteredBeginPhases = append(filteredBeginPhases, v)
				}
			}
		}
		break
	}

	return filteredAmplitudes, filteredBeginPhases
}

func lowPassFilter(upperValue float64) func(value float64) bool {
	return func(value float64) bool {
		return value <= upperValue
	}
}

func highPassFilter(lowerValue float64) func(value float64) bool {
	return func(value float64) bool {
		return value >= lowerValue
	}
}

func bandPassFilter(lowerValue float64, upperValue float64) func(value float64) bool {
	return func(value float64) bool {
		return value >= lowerValue && value <= upperValue
	}
}

func FilterByMovingAverageAlgorithm(signal []float64) []float64 {
	var newValues []float64

	m := int((MOVING_AVERAGE__K - 1) / 2)

	for index := range signal {
		newValues = append(newValues, getXByMovingAverage(index, m, MOVING_AVERAGE__K, signal))
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

func FilterByParabolaAlgorithm(signal []float64) []float64 {
	var newValues []float64

	for index := range signal {
		newValues = append(newValues, getXByParabolaEleven(index, signal))
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

func FilterByMedianFilterAlgorithm(signal []float64) []float64 {
	var newValues []float64

	for index := range signal {
		newValues = append(newValues, getXByMedianFilter(index, signal))
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