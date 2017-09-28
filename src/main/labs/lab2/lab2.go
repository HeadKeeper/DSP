package lab2

import "main/util"

const (
	AXIS_MIN_X = -10

	LAB_NAME = "lab2"
	INITIAL_PATH string = LAB_NAME + "/"
	PLOT_NAME string = "plot"
)

func PerformSecondOption()  {
	plot := util.CreatePlot("Polyharmonic", "n", "f(n)",
		AXIS_MIN_X, util.SIGNAL_RATE * util.SOUND_LENGTH, -40, 40)

	for _, currentM := range createMArray() {
		function := createFunction(float64(util.SOUND_LENGTH))
		util.AddFunctionOnPlot(plot, function, util.CreateRandomColor(), "")
	}

	util.SavePlotImage(INITIAL_PATH + PLOT_NAME, plot)
}

