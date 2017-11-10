package randomvariate

import (
	"contagion/utils"
	"math"
	"math/rand"
	"sort"
)

// Multinomial draws n samples from a probability distribution given by the
// set of probabilities p. Uses the alias method.
func Multinomial(n int, p []float64) []int {
	result := make([]int, len(p))

	// Create an discrete uniform distribution by combining categories
	// together using the alias method
	limit := 1.0 / float64(len(p))
	s := NewSlice(p...)
	sort.Sort(s)

	aliasIdx := make(map[int][]int)
	aliasValues := make(map[int][]float64)
	b := 0
	a := len(p) - 1
	for b < a {
		gap := limit - s.Float64Slice[b]
		// subtract from highest to fill the lowest
		aliasIdx[b] = []int{s.idx[b], s.idx[a]}
		//aliasValues[b] = []float64{s.Float64Slice[b], limit}
		aliasValues[b] = []float64{s.Float64Slice[b] / limit, 1.0}
		s.Float64Slice[a] -= gap
		s.Float64Slice[b] = limit
		aliasIdx[a] = []int{s.idx[a]}
		//aliasValues[a] = []float64{Round(s.Float64Slice[a], -8)}
		aliasValues[a] = []float64{utils.Round(s.Float64Slice[a]/limit, -8)}

		// check if highest is still above limit
		// if less than limit, fill to limit with next highest
		if utils.Round(s.Float64Slice[a], -8) == utils.Round(limit, -8) {
			a--
		} else if s.Float64Slice[a] < limit {
			gap := limit - s.Float64Slice[a]
			aliasIdx[a] = append(aliasIdx[a], s.idx[a-1])
			//aliasValues[a] = append(aliasValues[a], limit)
			aliasValues[a] = append(aliasValues[a], 1.0)
			// go to next highest
			a--
			s.Float64Slice[a] -= gap
			s.Float64Slice[a+1] = limit
		}
		b++
	}
	for i := 0; i < n; i++ {
		// Generate pseudo-random integer
		r := rand.Intn(len(p))
		if len(aliasValues[r]) > 1 {
			// Generate pseudorandom number
			if x := rand.Float64(); x <= aliasValues[r][0] {
				result[aliasIdx[r][0]]++
			} else {
				result[aliasIdx[r][1]]++
			}
		} else {
			result[aliasIdx[r][0]]++
		}
	}
	return result
}

// MultinomialI draws n samples from a probability distribution given by the
// set of probabilities p. Uses the inversion method which may be inefficient
// when the number of categories is large.
func MultinomialI(n int, p []float64) []int {
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

type Slice struct {
	sort.Float64Slice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(n ...float64) *Slice {
	// Copy
	nn := make([]float64, len(n))
	copy(nn, n)
	s := &Slice{Float64Slice: sort.Float64Slice(nn), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}
