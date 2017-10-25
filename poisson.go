package randomvariate

import (
	"math"
	"math/rand"
)

func Poisson(lambda float64) int {
	L := math.Exp(-1 * lambda)
	k := 0
	p := 1.
	for p > L {
		k++
		p *= rand.Float64()
	}
	return int(k - 1)
}

func PoissonXL(lambda float64) int {
	c := 0.767 - (3.36 / lambda)
	beta := math.Pi / math.Sqrt(3.0*lambda)
	alpha := beta * lambda
	k := math.Log(c) - lambda - math.Log(beta)

	for {
		u := rand.Float64()
		x := (alpha - math.Log((1.0-u)/u)) / beta
		n := math.Floor(x + 0.5)
		if n < 0 {
			continue
		}

		v := rand.Float64()
		y := alpha - (beta * x)
		lhs := y + math.Log(v/math.Pow((1.0+math.Exp(y)), 2))

		logNFactorial, _ := math.Lgamma(n + 1)
		rhs := k + n*math.Log(lambda) - logNFactorial
		if lhs <= rhs {
			return int(n)
		}
	}
}
