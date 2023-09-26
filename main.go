package main

import (
	"os"
)

var (
	file, _    = os.Create("sine_wave.wav")
	duration   = 8.0
	sampleRate = 44100.0
	numSamples = int(sampleRate * duration)

	frequency    = 440.0
	FMModFreq    = 1.0
	AMModFreq    = 0.7
	FMModDepth   = 10.0
	AMModDepth   = 1.0
	numHarmonics = 10
)

func main() {
	// Write WAV file header
	writeWavHeader(file)
	writeSamples(file)

	file.Close()
}
