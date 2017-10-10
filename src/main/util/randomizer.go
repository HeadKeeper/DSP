package util

import (
	"image/color"
	"math/rand"
	"golang.org/x/image/colornames"
	"time"
)

func CreateRandomColor() color.RGBA {
	var colors []color.RGBA
	for _, value := range colornames.Map {
		colors = append(colors, value)
	}
	return colors[random(0, 146)]
}

func GetRandomValue(values []float64) float64 {
	index := random(0, len(values) - 1)
	return values[index]
}

func createRandomUint8() uint8 {
	return uint8(random(0, 255))
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max - min) + min
}
