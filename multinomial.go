package randomvariate

import (
	"math"
	"math/rand"
)

// Multinomial draws n samples from a probability distribution given by the
// set of probabilities p. Uses the inversion method which may be inefficient
// when the number of categories and number of samples are both large.
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

// MultinomialA draws n samples from a probability distribution given by the
// set of probabilities p. Uses the alias method. Faster when dealing with
// a larger number of categories and number of samples.
func MultinomialA(n int, p []float64) []int {
	// Setup uniform distribution
	K := len(p)
	q := make([]float64, K)
	J := make([]int, K)

	var smaller []int
	var larger []int
	for i, prob := range p {
		q[i] = float64(K) * prob
		if q[i] < 1.0 {
			smaller = append(smaller, i)
		} else {
			larger = append(larger, i)
		}
	}

	var small, large int
	for len(smaller) > 0 && len(larger) > 0 {
		small, smaller = smaller[len(smaller)-1], smaller[:len(smaller)-1]
		large, larger = larger[len(larger)-1], larger[:len(larger)-1]

		J[small] = large
		q[large] = float64(q[large] - (1.0 - q[small]))

		if q[large] < 1.0 {
			smaller = append(smaller, large)
		} else {
			larger = append(larger, large)
		}
	}

	// Draw sample
	result := make([]int, len(p))
	for i := 0; i < n; i++ {
		kk := rand.Intn(K)
		if rand.Float64() < float64(q[kk]) {
			result[kk]++
		} else {
			result[J[kk]]++
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
