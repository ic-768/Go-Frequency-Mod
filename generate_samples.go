package main

import "math"

func calculateAngle(t float64) float64 {
	return 2.0 * math.Pi * t
}

// TODO pass frequency as a param - it'll be more reusable
func sineGenerator(t float64) float64 {
	return math.Sin(calculateAngle(t) * frequency)
}

// Create AM modulated sine wave samples
func AMGenerator(t float64) float64 {
	AMModAngle := calculateAngle(t) * AMModFreq
	AMModulator := (1 + math.Sin(AMModAngle)) / 2
	// Sweep AMMod upwards
	//AMModFreq *= 1.00001

	return (AMModulator * AMModDepth) * sineGenerator(t)
}

// Create FM modulated sine wave samples
func FMGenerator(t float64) float64 {
	angle := calculateAngle(t)
	FMModulator := math.Sin(angle * FMModFreq)
	FMModAngle := angle*frequency + (FMModulator * FMModDepth)

	return math.Sin(FMModAngle)
}

// Create AM and FM modulated sine wave samples
func AMFMGenerator(t float64) float64 {
	angle := calculateAngle(t)

	AMModulator := (1 + math.Sin(angle*AMModFreq)) / 2
	FMModulator := math.Sin(angle * FMModFreq)
	FMModAngle := angle*frequency + (FMModulator * FMModDepth)

	return (AMModulator * AMModDepth) * math.Sin(FMModAngle)
}

// Choose which sine wave sample generation function to use based on
// which modulation depths are non-zero
func chooseGenerator() func(float64) float64 {
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
func generateSample(sampleNum int, generate func(float64) float64) int16 {
	t := float64(sampleNum) / sampleRate
	sample := generate(t)
	sampleInt := int16(sample * 32767)

	return sampleInt
}
