package main

import (
	"encoding/binary"
	"os"
)

func writeSamples(file *os.File) {
	sampleGenerator := chooseGenerator()

	for i := 0; i < numSamples; i++ {
		sample := generateSample(i, sampleGenerator)
		writeSample(sample)
	}
}

func writeSample(sample int16) {
	binary.Write(file, binary.LittleEndian, sample)
}
