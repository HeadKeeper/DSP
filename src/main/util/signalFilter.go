package util

const (
	LOW_PASS_FILTER = "LOW_PASS"
	HIGH_PASS_FILTER = "HIGH_PASS"
	BAND_PASS_FILTER = "BAND_PASS"
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

