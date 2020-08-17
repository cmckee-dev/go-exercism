// Package space contains nothing to do wih Space Pants or Peter Dinklage.
package space

// Planet is a string representing a planet in our Solar System.
type Planet string

// Age returns a float64 representing your age on other planets from our Solar
// System based on their orbital periods.
func Age(ageInSeconds float64, planet Planet) float64 {
	var planetRotations = map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return (ageInSeconds) / (planetRotations[planet] * float64(31557600))
}
