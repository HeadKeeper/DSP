package lab1

import "math"

const DEFAULT_N float64 = 512

type FormulaData struct {
	value      string
	parameters map[string]interface{}
}

func getFirstFormula(n float64, y float64) FormulaData {
	formula := "A*sin(((2*Pi*f*n)/N)*y)"
	parameters := make(map[string]interface{}, 5)
	parameters["A"] = 9
 	parameters["f"] = 4
 	parameters["n"] = n
 	parameters["y"] = y
	parameters["N"] = DEFAULT_N
	parameters["Pi"] = math.Pi

	return FormulaData{
		formula,
		parameters,
	}
}

func getSecondFormula() FormulaData {
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = 10
	parameters["y"] = 12

	return FormulaData{
		"",
		parameters,
	}
}

func getThirdFormula() FormulaData {
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = 10
	parameters["y"] = 12

	return FormulaData{
		"",
		parameters,
	}
}

func getFourthFormula() FormulaData {
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = 10
	parameters["y"] = 12

	return FormulaData{
		"",
		parameters,
	}
}
