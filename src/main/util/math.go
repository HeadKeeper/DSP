package util

import (
	"github.com/Knetic/govaluate"
	"fmt"
	"math"
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