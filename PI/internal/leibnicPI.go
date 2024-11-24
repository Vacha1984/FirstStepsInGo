package pi

import "math"

func LeibnicSum(i int, c chan float64) {
	sum := math.Pow(-1, float64(i)) / (2.0*float64(i) + 1)
	if i%2 == 0 {
		c <- sum
	} else {
		c <- -sum
	}
}
