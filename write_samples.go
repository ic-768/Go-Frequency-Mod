package main

import (
	"encoding/binary"
	"os"
)

func writeSample(sample int16) {
	binary.Write(file, binary.LittleEndian, sample)
}

// Generate a single sample based based on a provided sample generation function
func generateSample(sampleNum int, generate func(float64, float64, int) float64) int16 {
	t := float64(sampleNum) / sampleRate
	sample := generate(t, frequency, numHarmonics)
	sampleInt := int16(sample * 32767)

	return sampleInt
}

func writeSamples(file *os.File) {
	sampleGenerator := chooseGenerator()

	for i := 0; i < numSamples; i++ {
		sample := generateSample(i, sampleGenerator)
		writeSample(sample)
	}
}
