package main

import (
	"os"
)

var (
	sampleRate = 44100.0
	duration   = 80.0
	frequency  = 440.0
	FMModFreq  = 2.0
	AMModFreq  = 0.2
	FMModDepth = 10.0
	AMModDepth = 1.0
	numSamples = int(sampleRate * duration)
	file, _    = os.Create("sine_wave.wav")
)

func main() {
	// Write WAV file header
	writeWavHeader(file)
	writeSamples(file)

	file.Close()
}
