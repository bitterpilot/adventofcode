package fuel

import "math"

// CounterUpper needs to know the total fuel requirement.
// To find it, individually calculate the fuel needed for the mass of
// each module (your puzzle input), then add together all the fuel values
func CounterUpper(masses []float64) float64 {
	var requiredFuel float64

	for _, mass := range masses {
		requiredFuel += moduleFuelRequirement(mass)
	}

	return requiredFuel
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass,
// divide by three, round down, and subtract 2.
func moduleFuelRequirement(mass float64) float64 {
	var fuel float64
	fuel += math.Floor(mass/3) - 2
	mass = fuel
	if mass > 0 {
		fuel += moduleFuelRequirement(mass)
	}
	if fuel < 0 {
		fuel = 0
	}

	return fuel
}
