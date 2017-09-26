package util

import (
	"fmt"
	"os"

	"github.com/cryptix/wav"
	"math"
)

const (
	DEFAULT_CHANNELS = 1				// 1
	DEFAULT_BITS     = 16 				// Tone is lower when value is lower (8, 16, 32, 64 ...)
	DEFAULT_RATE 	 = 32768 			// 44100

	DEFAULT_OUT_PATH = "out/"
)

func WriteWAV(name string, soundLength int, function func(n float64) float64) {
	wavOut, err := os.Create(DEFAULT_OUT_PATH + name + ".wav")
	checkErr(err)
	defer wavOut.Close()

	meta := wav.File{
		Channels:        DEFAULT_CHANNELS,
		SampleRate:      DEFAULT_RATE,
		SignificantBits: DEFAULT_BITS,
	}

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	for n := 0; n < soundLength * DEFAULT_RATE; n++ {
		funRes := int32(
			math.Pow(2, DEFAULT_BITS-1) *
			function(float64(n)),
		)

		err := writer.WriteInt32(funRes)
		checkErr(err)
	}

	writer.Close()
	fmt.Println()
	fmt.Println("WAV file created successful")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}