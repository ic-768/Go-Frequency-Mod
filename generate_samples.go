package main

import "math"

func SineGenerator(t float64) float64 {
	angle := 2.0 * math.Pi * t

	return math.Sin(angle * frequency)
}

func AMGenerator(t float64) float64 {
	angle := 2.0 * math.Pi * t
	AMmodulator := (1 + math.Sin(angle*AMModFreq)) / 2

	// Sweep AMMod upwards
	AMModFreq *= 1.00001
	return (AMmodulator * AMModDepth) * SineGenerator(t)
}

func FMGenerator(t float64) float64 {
	angle := 2.0 * math.Pi * t
	FMmodulator := math.Sin(angle * FMModFreq)

	return math.Sin(angle*frequency + (FMmodulator * FMModDepth))
}

func AMFMGenerator(t float64) float64 {
	angle := 2.0 * math.Pi * t
	AMmodulator := (1 + math.Sin(angle*AMModFreq)) / 2
	FMmodulator := math.Sin(angle * FMModFreq)

	return (AMmodulator * AMModDepth) * math.Sin(angle*frequency+(FMmodulator*FMModDepth))
}
