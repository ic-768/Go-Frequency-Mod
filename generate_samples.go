package main

import "math"

func sineGenerator(t float64, frequency float64) float64 {
	return math.Sin(2.0 * math.Pi * t * frequency)
}

// Create AM modulated sine wave samples
func AMGenerator(t float64, frequency float64) float64 {
	AMModulator := sineGenerator(t, AMModFreq)
	AMModSignal := ((1 + AMModulator) / 2) * AMModDepth
	sample := (sineGenerator(t, frequency))

	for i := 0; i < numHarmonics; i += 2 {
		sample += (sineGenerator(t, frequency*float64(i+1*2)))
	}

	if numHarmonics > 0 {
		sample /= (float64(numHarmonics + 1))
	}

	sample *= AMModSignal

	return sample
}

// Create FM modulated sine wave samples
func FMGenerator(t float64, frequency float64) float64 {
	FMModulator := sineGenerator(t, FMModFreq)
	FMModSignal := FMModulator * FMModDepth
	sample := sineGenerator(t, frequency+FMModSignal)

	return sample
}

// Create AM and FM modulated sine wave samples
func AMFMGenerator(t float64, frequency float64) float64 {
	AMModulator := sineGenerator(t, AMModFreq)
	AMModSignal := ((1 + AMModulator) / 2) * AMModDepth
	FMModulator := sineGenerator(t, FMModFreq)
	FMModSignal := FMModulator * FMModDepth
	sample := (AMModSignal) * sineGenerator(t, frequency+FMModSignal)

	return sample
}

// Choose which sine wave sample generation function to use based on
// which modulation depths are non-zero
func chooseGenerator() func(float64, float64) float64 {
	if AMModDepth != 0.0 && FMModDepth != 0.0 {
		return AMFMGenerator
	} else if AMModDepth != 0 {
		return AMGenerator
	} else if FMModDepth != 0 {
		return FMGenerator
	} else {
		return sineGenerator
	}
}

// Generate a single sample based based on a provided sample generation function
func generateSample(sampleNum int, generate func(float64, float64) float64) int16 {
	t := float64(sampleNum) / sampleRate
	sample := generate(t, frequency)
	sampleInt := int16(sample * 32767)

	return sampleInt
}
