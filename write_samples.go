package main

import (
	"encoding/binary"
	"os"
)

func writeSamples(file *os.File) {
	sampleGenerator := SineGenerator

	if AMModDepth != 0.0 && FMModDepth != 0.0 {
		sampleGenerator = AMFMGenerator
	} else if AMModDepth != 0 {
		sampleGenerator = AMGenerator
	} else if FMModDepth != 0 {
		sampleGenerator = FMGenerator
	}
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		sample := sampleGenerator(t)
		sampleInt := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, sampleInt)
	}
}
