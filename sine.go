package main

import "math"

// Create a sine sample with (optional) harmonics
func sineGenerator(t float64, frequency float64, numHarmonics int) float64 {
	sample := math.Sin(2.0 * math.Pi * t * frequency)

	if numHarmonics > 0 {
		addHarmonics(t, &sample)
	}

	return sample
}

// Enrich a sample with harmonics
func addHarmonics(t float64, sample *float64) {
	generator := chooseGenerator()

	for i := 0; i < numHarmonics; i += 2 {
		// we want FM/AM to affect harmonics too
		*sample += generator(t, frequency*float64(i+1*2), 0)
	}

	// normalise because adding a bunch of harmonics is gonna peak the signal
	*sample /= (float64(numHarmonics + 1))
}
