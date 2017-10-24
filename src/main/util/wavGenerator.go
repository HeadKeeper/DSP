package util

import (
	"fmt"
	"os"

	"github.com/cryptix/wav"
	"main/types"
)

const (
	DEFAULT_CHANNELS = 1				// 1
	DEFAULT_BITS     = 32 				// Tone is lower when value is lower (8, 16, 32, 64 ...)
	DEFAULT_RATE 	 = 44100 			// 44100 32768

	DEFAULT_OUT_PATH = "out/"
)

func WriteWAV(name string, soundLength int, functionData types.PlotData) {
	var signal []float64
	for n := functionData.InitialN; n < functionData.EndN; n+= functionData.Step {
		signal = append(signal, functionData.Function(float64(n)))
	}
	WriteWAVForSignal(name, soundLength, signal)
}

func WriteWAVForSignal(name string, soundLength int, signal []float64) {
	wavOut, err := os.Create(DEFAULT_OUT_PATH + name + ".wav")
	checkErr(err)
	defer wavOut.Close()

	meta := wav.File{
		Channels:        DEFAULT_CHANNELS,
		SampleRate:      DEFAULT_RATE,
		SignificantBits: DEFAULT_BITS,
		//NumberOfSamples: uint32(len(signal) * SOUND_LENGTH),
	}

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	for n := 0; n < soundLength / 2; n++ {
		for idx := range signal {
			funRes := int32(
				signal[idx],
			)
			writer.WriteInt32(funRes)
		}
		checkErr(err)
	}

	writer.Close()
	fmt.Println()
	fmt.Println("WAV file '" + name + "' created successful")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}