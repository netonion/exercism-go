package space

const secondsInYear = 31557600

type Planet string

func Ratios(planet Planet) float64 {
	if planet == "Mercury" {
		return 0.2408467
	}
	if planet == "Venus" {
		return 0.61519726
	}
	if planet == "Mars" {
		return 1.8808158
	}
	if planet == "Jupiter" {
		return 11.862615
	}
	if planet == "Saturn" {
		return 29.447498
	}
	if planet == "Uranus" {
		return 84.016846
	}
	if planet == "Neptune" {
		return 164.79132
	}
	return 1.0
}

func Age(seconds float64, planet Planet) float64 {
	return float64(int(float64(seconds)/secondsInYear/Ratios(planet)*100+0.5)) / 100
}
