package fuel

import "math"

// CounterUpper needs to know the total fuel requirement.
// To find it, individually calculate the fuel needed for the mass of
// each module (your puzzle input), then add together all the fuel values
func CounterUpper(masses []float64) float64 {
	var requiredFuel float64

	for _, mass := range masses {
		requiredFuel += moduelFuelRequierment(mass)
	}

	return requiredFuel
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass,
// divide by three, round down, and subtract 2.
func moduelFuelRequierment(mass float64) float64 {
	return math.Floor(mass/3) - 2
}
