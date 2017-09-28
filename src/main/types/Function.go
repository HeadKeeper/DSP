package types

type FunctionData struct {
	Function func(n float64) float64
	Name string
	InitialN float64
	EndN float64
	Step float64
}
