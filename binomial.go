package randomvariate

// Binomial draws n samples from a binomial probability distribution given
// by the probability p.
// Another way to imagine this is that the function counts the number of
// sucessful draws. If n = 1, the function draws once and the result can
// either be 1 or 0. The probability of successfully drawing is given by
// the slice p.
func Binomial(n int, p float64) int {
	return MultinomialI(n, []float64{p, 1.0 - p})[0]
}
