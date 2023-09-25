package main

import "math"

func calculateAngle(t float64) float64 {
	return 2.0 * math.Pi * t
}

func sineGenerator(t float64) float64 {
	return math.Sin(calculateAngle(t) * frequency)
}

func AMGenerator(t float64) float64 {
	AMModulator := (1 + math.Sin(calculateAngle(t)*AMModFreq)) / 2
	// Sweep AMMod upwards
	//AMModFreq *= 1.00001

	return (AMModulator * AMModDepth) * sineGenerator(t)
}

func FMGenerator(t float64) float64 {
	angle := calculateAngle(t)
	FMModulator := math.Sin(angle * FMModFreq)

	return math.Sin(angle*frequency + (FMModulator * FMModDepth))
}

func AMFMGenerator(t float64) float64 {
	angle := calculateAngle(t)

	AMModulator := (1 + math.Sin(angle*AMModFreq)) / 2
	FMModulator := math.Sin(angle * FMModFreq)

	return (AMModulator * AMModDepth) * math.Sin(angle*frequency+(FMModulator*FMModDepth))
}
