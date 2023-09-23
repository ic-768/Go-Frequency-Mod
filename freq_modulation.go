package main

import (
	"encoding/binary"
	"math"
	"os"
)

func main() {
	// Constants
	sampleRate := 44100.0
	duration := 3.0
	frequency := 440.0
	modFrequency := 5.0
	modDepth := 3.0

	// Calculate number of samples
	numSamples := int(sampleRate * duration)

	// Create a new WAV file
	file, _ := os.Create("sine_wave.wav")

	// Write WAV file header
	writeWavHeader(file, numSamples, sampleRate)

	// Create modulator, generate and write audio samples
	for i := 0; i < numSamples; i++ {
		t := float64(i) / sampleRate
		modulator := math.Sin(2.0 * math.Pi * t * modFrequency)
		sample := math.Sin(2.0*math.Pi*frequency*t + (modulator * modDepth))
		sampleInt := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, sampleInt)
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
