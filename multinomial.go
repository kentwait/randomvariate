package randomvariate

import "math/rand"

// multinomial draws n samples from a population whose members belong to
// distinct categories. The probability of drawing a sample from a category
// is given by the slice p. The function returns a slice of integers where the
// position represents a particular category and the value indicates the
// number of samples drawn for that particular category.
func multinomial(n int, p []float64) []int {
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
