package main

import (
	"os"
)

// Because samples are calculated concurrently (potentially writing to file out of order), we store the sample index too
type SampleBuffer struct {
	sample int16
	index  int
}

var (
	file, _       = os.Create("sine_wave.wav")
	duration      = 200.0
	sampleRate    = 44100.0
	numSamples    = int(sampleRate * duration)
	sampleChannel = make(chan SampleBuffer, numSamples)

	frequency    = 440.0
	FMModFreq    = 1.0
	AMModFreq    = 0.7
	FMModDepth   = 0.0
	AMModDepth   = 0.0
	numHarmonics = 30
)

func main() {
	// Write WAV file header
	writeWavHeader()
	writeSamples()

	file.Close()
}
