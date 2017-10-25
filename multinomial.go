package randomvariate

import (
	"math"
	"math/rand"
)

// Multinomial draws n samples from a probability distribution given by the
// set of probabilities p.
// Another way to imagine this is that the function draws from a population
// whose members belong to distinct categories. The probability of drawing a
// sample from a category is given by the slice p. The function returns a
// slice of integers where the position represents a particular category and
// the value indicates the number of samples drawn for that particular
// category.
func Multinomial(n int, p []float64) []int {
	result := make([]int, len(p))
	cumP := make([]float64, len(p))
	lastIdx := len(p) - 1

	// Create a cummulative distribution of p
	cumP[0] = p[0]
	for i := 1; i < len(p); i++ {
		cumP[i] = cumP[i-1] + p[i]
	}
	for i := 0; i < n; i++ {
		// Generate pseudorandom number
		x := rand.Float64()
		for j := 0; j < len(cumP); j++ {
			if x < cumP[j] {
				result[j]++
				break
			} else if x > cumP[lastIdx] && x <= 1.0 {
				result[lastIdx]++
				break
			}
		}
	}
	return result
}

// MultinomialLog1p draws n samples from a log-probability distribution given
// by the set of probabilities p. Note that the log probabilities are actually
// log(1+p) where p is from 0 to 1. This prevents solves the problem of
// computing log probability of 0.
func MultinomialLog1p(n int, logP []float64) []int {
	// Transform log probabilities into decimal
	p := make([]float64, len(logP))
	for i, logProb := range logP {
		p[i] = math.Expm1(logProb)
	}
	return Multinomial(n, p)
}

// MultinomialLog draws n samples from a log-probability distribution given
// by the set of probabilities p. Note that the log probabilities are in the
// format log(p) where p is from 0 to 1. If p = 0, the log-probability should
// be encoded as negative infinity.
func MultinomialLog(n int, logP []float64) []int {
	// Transform log probabilities into decimal
	p := make([]float64, len(logP))
	for i, logProb := range logP {
		if math.IsInf(logProb, -1) {
			p[i] = 0
		} else if logProb == 0 {
			p[i] = 1
		} else {
			p[i] = math.Exp(logProb)
		}
	}
	return Multinomial(n, p)
}
