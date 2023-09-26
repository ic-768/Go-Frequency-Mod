package main

// Create AM modulated sine wave samples
func AMGenerator(t float64, frequency float64, numHarmonics int) float64 {
	AMModulator := sineGenerator(t, AMModFreq, 0)
	AMModSignal := ((1 + AMModulator) / 2) * AMModDepth

	sample := AMModSignal * sineGenerator(t, frequency, numHarmonics)

	return sample
}

// Create FM modulated sine wave samples
func FMGenerator(t float64, frequency float64, numHarmonics int) float64 {
	FMModulator := sineGenerator(t, FMModFreq, 0)
	FMModSignal := FMModulator * FMModDepth

	sample := sineGenerator(t, frequency+FMModSignal, numHarmonics)

	return sample
}

// Create AM and FM modulated sine wave samples
func AMFMGenerator(t float64, frequency float64, numHarmonics int) float64 {
	AMModulator := sineGenerator(t, AMModFreq, 0)
	AMModSignal := ((1 + AMModulator) / 2) * AMModDepth
	FMModulator := sineGenerator(t, FMModFreq, 0)
	FMModSignal := FMModulator * FMModDepth

	sample := AMModSignal * sineGenerator(t, frequency+FMModSignal, numHarmonics)

	return sample
}

// Choose which sine wave sample generation function to use based on which modulation depths are non-zero
func chooseGenerator() func(float64, float64, int) float64 {
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
