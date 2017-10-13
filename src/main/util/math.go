package util

import (
	"github.com/Knetic/govaluate"
	"fmt"
	"math"
)

const (
	// w = 2 * Pi * f - angular speed of rotation
	// phase = (2 * Pi * f * x) / N) * soundLength + initialAngle
	SIGNAL_RATE float64 = 512
	SOUND_LENGTH = 15
)

var functions = map[string]govaluate.ExpressionFunction {
	"log": func(args ...interface{}) (interface{}, error) {
		value := math.Log(args[0].(float64))
		return (float64)(value), nil
	},
	"sin": func(arguments ...interface{}) (interface{}, error) {
		value := math.Sin(arguments[0].(float64))
		return (float64)(value), nil
	},
}

func EvaluateFormula(formula string, values map[string]interface{}) interface{} {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(formula, functions)
	result, err := expression.Evaluate(values)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return result
}

func GetHarmonicFunction(amplitude float64, frequency float64, phi float64) func(n float64) float64 {
	return func(n float64) float64 {
		return amplitude * math.Sin(2 * math.Pi * frequency * n / SIGNAL_RATE + phi) }
}


func GetValues(initialN float64, endN float64, step float64) []float64 {
	var n float64
	var values []float64
	for n = initialN; n < endN; n += step {
		values = append(values, n)
	}

	return values
}