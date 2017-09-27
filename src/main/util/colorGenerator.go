package util

import (
	"image/color"
	"math/rand"
	"golang.org/x/image/colornames"
)

func CreateRandomColor() color.RGBA {
	var colors []color.RGBA
	for _, value := range colornames.Map {
		colors = append(colors, value)
	}
	return colors[random(0, 146)]
}

func createRandomUint8() uint8 {
	return uint8(random(0, 255))
}

func random(min, max int) int {
	rand.Seed(rand.Int63())
	return rand.Intn(max - min) + min
}