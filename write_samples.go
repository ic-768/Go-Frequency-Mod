package main

import (
	"encoding/binary"
	"sync"
)

func writeSample(sample int16) {
	binary.Write(file, binary.LittleEndian, sample)
}

func generateSample(sampleNum int, generate func(float64, float64, int) float64, sampleChannel chan<- SampleBuffer) {
	t := float64(sampleNum) / sampleRate
	sample := generate(t, frequency, numHarmonics)
	sampleInt := int16(sample * 32767)

	sampleChannel <- SampleBuffer{sample: sampleInt, index: sampleNum}
}

func writeSamples() {
	sampleGenerator := chooseGenerator()

	// Create a slice to store the samples
	samples := make([]int16, numSamples)

	var wg sync.WaitGroup

	// concurrently create samples to be written
	for i := 0; i < numSamples; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			generateSample(i, sampleGenerator, sampleChannel)
			sampleData := <-sampleChannel
			samples[sampleData.index] = sampleData.sample
		}(i)
	}

	wg.Wait()

	// Write samples in order
	for i := 0; i < numSamples; i++ {
		writeSample(samples[i])
	}
}
