package main

import (
	"os"
)

var (
	sampleRate float64 = 44100.0
	duration   float64 = 80.0
	frequency  float64 = 440.0
	FMModFreq  float64 = 1.0
	AMModFreq  float64 = 4.0
	FMModDepth float64 = 0.2
	AMModDepth float64 = 0.0
	numSamples int     = int(sampleRate * duration)
)

func main() {
	// Create a new WAV file
	file, _ := os.Create("sine_wave.wav")

	// Write WAV file header
	writeWavHeader(file, numSamples, sampleRate)

	if AMModDepth != 0.0 {
		generateAmplitudeModulatedSamples(file)
	} else {
		generateSamples(file)
	}

	file.Close()
}
