package main

import (
	"os"
)

var (
	sampleRate = 44100.0
	duration   = 80.0
	frequency  = 440.0
	FMModFreq  = 2.0
	AMModFreq  = 1.0
	FMModDepth = 0.0
	AMModDepth = 0.8
	numSamples = int(sampleRate * duration)
	file, _    = os.Create("sine_wave.wav")
)

func main() {
	// Write WAV file header
	writeWavHeader(file)
	writeSamples(file)

	file.Close()
}
