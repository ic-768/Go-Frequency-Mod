package main

import "math"

func calculateAngle(t float64) float64 {
	return 2.0 * math.Pi * t
}

func sineGenerator(t float64) float64 {
	return math.Sin(calculateAngle(t) * frequency)
}

func AMGenerator(t float64) float64 {
	AMmodulator := (1 + math.Sin(calculateAngle(t)*AMModFreq)) / 2
	// Sweep AMMod upwards
	//AMModFreq *= 1.00001
	return (AMmodulator * AMModDepth) * sineGenerator(t)
}

func FMGenerator(t float64) float64 {
	angle := calculateAngle(t)
	FMmodulator := math.Sin(angle * FMModFreq)

	return math.Sin(angle*frequency + (FMmodulator * FMModDepth))
}

func AMFMGenerator(t float64) float64 {
	angle := calculateAngle(t)

	AMmodulator := (1 + math.Sin(angle*AMModFreq)) / 2
	FMmodulator := math.Sin(angle * FMModFreq)

	return (AMmodulator * AMModDepth) * math.Sin(angle*frequency+(FMmodulator*FMModDepth))
}
