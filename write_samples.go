package main

import (
	"encoding/binary"
	"math"
	"os"
)

// Has only FM (which can be 0)
func generateSamples(file *os.File) {
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		angle := 2.0 * math.Pi * t
		FMmodulator := math.Sin(angle * FMModFreq)
		sample := math.Sin(angle*frequency + (FMmodulator * FMModDepth))
		sampleInt := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, sampleInt)
		// Sweep FMMod upwards
		FMModFreq *= 1.00001
	}
}

// Has both AM and FM
func generateAmplitudeModulatedSamples(file *os.File) {
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		angle := 2.0 * math.Pi * t
		FMmodulator := math.Sin(angle * FMModFreq)
		AMmodulator := (1 + math.Sin(angle*AMModFreq)) / 2
		sample := (AMmodulator * AMModDepth) * math.Sin(angle*frequency+(FMmodulator*FMModDepth))
		sampleInt := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, sampleInt)
		// Sweep AMMod upwards
		AMModFreq *= 1.00001
	}
}
