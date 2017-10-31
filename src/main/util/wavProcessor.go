package util

import (
	"fmt"
	"os"

	"github.com/cryptix/wav"
	"main/types"
	"io"
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
	meta := wav.File{
		Channels:        DEFAULT_CHANNELS,
		SampleRate:      DEFAULT_RATE,
		SignificantBits: DEFAULT_BITS,
	}
	WriteWAVByMeta(name, soundLength, signal, meta)
}

func WriteWAVByMeta(name string, soundLength int, signal []float64, meta wav.File) {
	wavOut, err := os.Create(DEFAULT_OUT_PATH + name + ".wav")
	checkErr(err)
	defer wavOut.Close()

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	for n := 0; n < soundLength; n++ {
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

func ReadWAV(path string) ([]float64, wav.File){
	testInfo, err := os.Stat(path)
	checkErr(err)

	testWav, err := os.Open(path)
	checkErr(err)

	wavReader, err := wav.NewReader(testWav, testInfo.Size())
	checkErr(err)

	fmt.Println(wavReader)


	var i uint32
	var signal []float64
	for i = 0; i < wavReader.GetSampleCount(); i++{
		n, err := wavReader.ReadSample()
		if err == io.EOF {
			break
		}
		checkErr(err)

		signal = append(signal, float64(n))
	}
	inputMeta := wavReader.GetFile()
	meta := wav.File {
		Channels:        inputMeta.Channels,
		SampleRate:      inputMeta.SampleRate,
		SignificantBits: inputMeta.SignificantBits,
		Canonical: inputMeta.Canonical,
		AudioFormat: inputMeta.AudioFormat,
	}

	return signal, meta
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}