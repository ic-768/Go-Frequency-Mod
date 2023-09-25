package main

import (
	"encoding/binary"
	"math"
	"os"
)

var (
	sampleRate float64 = 44100.0
	duration   float64 = 80.0
	frequency  float64 = 440.0
	FMModFreq  float64 = 5.0
	AMModFreq  float64 = 4.0
	FMModDepth float64 = 0.0
	AMModDepth float64 = 1.0
)

func main() {
	// Calculate number of samples
	numSamples := int(sampleRate * duration)

	// Create a new WAV file
	file, _ := os.Create("sine_wave.wav")

	// Write WAV file header
	writeWavHeader(file, numSamples, sampleRate)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		angle := 2.0 * math.Pi * t
		FMmodulator := math.Sin(angle * FMModFreq)
		AMmodulator := (1 + math.Sin(angle*AMModFreq)) / 2
		sample := (AMmodulator * AMModDepth) * math.Sin(angle*frequency+(FMmodulator*FMModDepth))
		sampleInt := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, sampleInt)
		// Sweep AMMod upwards
		//AMModFreq *= 1.00001
		// Sweep AMMod upwards
		FMModFreq *= 1.00001
	}
	file.Close()
}

func writeWavHeader(file *os.File, numSamples int, sampleRate float64) {
	// Constants
	numChannels := 1
	bitsPerSample := 16

	// Calculate sub-chunk sizes
	subChunk1Size := int32(16)                                           // Size of the PCM format subchunk
	subChunk2Size := int32(numSamples * numChannels * bitsPerSample / 8) // Size of the data subchunk

	fileSize := 36 + subChunk2Size

	// Write WAV file header
	file.Write([]byte("RIFF"))
	binary.Write(file, binary.LittleEndian, fileSize)
	file.Write([]byte("WAVEfmt "))
	binary.Write(file, binary.LittleEndian, subChunk1Size)
	binary.Write(file, binary.LittleEndian, int16(1))                                           // Audio format (PCM)
	binary.Write(file, binary.LittleEndian, int16(numChannels))                                 // Number of channels
	binary.Write(file, binary.LittleEndian, int32(int(sampleRate)))                             // Sample rate
	binary.Write(file, binary.LittleEndian, int32(int(sampleRate)*numChannels*bitsPerSample/8)) // Byte rate
	binary.Write(file, binary.LittleEndian, int16(numChannels*bitsPerSample/8))                 // Block align
	binary.Write(file, binary.LittleEndian, int16(bitsPerSample))                               // Bits per sample
	file.Write([]byte("data"))
	binary.Write(file, binary.LittleEndian, subChunk2Size)
}
