package randomvariate

func Binomial(n int, p float64) []int {
	return Multinomial(n, []float64{p, 1.0 - p})
}
