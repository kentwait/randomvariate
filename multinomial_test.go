package randomvariate

import (
	"math"
	"math/rand"
	"testing"
)

func TestMultinomial(t *testing.T) {
	cases := []struct {
		name string
		n    int
		p    []float64
	}{
		{name: "n=1,plen=2,dist=uniform",
			n: 1,
			p: []float64{0.5, 0.5},
		},
		{name: "n=1,plen=2,dist=skew_left",
			n: 1,
			p: []float64{0.9, 0.1},
		},
		{name: "n=1,plen=2,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.9},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0},
		},
		{name: "n=1,plen=3,dist=uniform",
			n: 1,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=1,plen=3,dist=skew_left",
			n: 1,
			p: []float64{0.6, 0.3, 0.1},
		}, {name: "n=1,plen=3,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=1,plen=4,dist=uniform",
			n: 1,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=1,plen=10,dist=uniform",
			n: 1,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
		{name: "n=2,plen=2,dist=uniform",
			n: 2,
			p: []float64{0.5, 0.5},
		},
		{name: "n=2,plen=2,dist=skew_left",
			n: 2,
			p: []float64{0.9, 0.1},
		},
		{name: "n=2,plen=2,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.9},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0},
		},
		{name: "n=2,plen=3,dist=uniform",
			n: 2,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=2,plen=3,dist=skew_left",
			n: 2,
			p: []float64{0.6, 0.3, 0.1},
		},
		{name: "n=2,plen=3,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=2,plen=4,dist=uniform",
			n: 2,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=2,plen=10,dist=uniform",
			n: 2,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
	}
	iterations := 1000
	errSize := 0.05
	rand.Seed(0)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Simulate
			var cnt [][]int
			for i := 0; i < iterations; i++ {
				result := multinomial(tc.n, tc.p)
				cnt = append(cnt, result)
			}
			// Sum columns
			sum := make([]int, len(tc.p))
			for _, row := range cnt {
				for c, v := range row {
					sum[c] += v
				}
			}
			// Check frequency
			for i, v := range sum {
				freq := float64(v) / float64(iterations*tc.n)
				expected := tc.p[i]
				if expected+errSize < freq || expected-errSize > freq {
					t.Errorf("frequency (%f) is greater than expected (%f) +/- (%f)", freq, expected, errSize)
				}
			}
		})
	}
}

func TestMultinomialLog1p(t *testing.T) {
	cases := []struct {
		name string
		n    int
		p    []float64
	}{
		{name: "n=1,plen=2,dist=uniform",
			n: 1,
			p: []float64{0.5, 0.5},
		},
		{name: "n=1,plen=2,dist=skew_left",
			n: 1,
			p: []float64{0.9, 0.1},
		},
		{name: "n=1,plen=2,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.9},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0},
		},
		{name: "n=1,plen=3,dist=uniform",
			n: 1,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=1,plen=3,dist=skew_left",
			n: 1,
			p: []float64{0.6, 0.3, 0.1},
		}, {name: "n=1,plen=3,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=1,plen=4,dist=uniform",
			n: 1,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=1,plen=10,dist=uniform",
			n: 1,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
		{name: "n=2,plen=2,dist=uniform",
			n: 2,
			p: []float64{0.5, 0.5},
		},
		{name: "n=2,plen=2,dist=skew_left",
			n: 2,
			p: []float64{0.9, 0.1},
		},
		{name: "n=2,plen=2,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.9},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0},
		},
		{name: "n=2,plen=3,dist=uniform",
			n: 2,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=2,plen=3,dist=skew_left",
			n: 2,
			p: []float64{0.6, 0.3, 0.1},
		},
		{name: "n=2,plen=3,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=2,plen=4,dist=uniform",
			n: 2,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=2,plen=10,dist=uniform",
			n: 2,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
	}
	iterations := 1000
	errSize := 0.05
	rand.Seed(0)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert to log probabilities
			logP := make([]float64, len(tc.p))
			for i := range tc.p {
				logP[i] = math.Log1p(tc.p[i])
			}
			// Simulate
			var cnt [][]int
			for i := 0; i < iterations; i++ {
				result := multinomialLog1p(tc.n, logP)
				cnt = append(cnt, result)
			}
			// Sum columns
			sum := make([]int, len(tc.p))
			for _, row := range cnt {
				for c, v := range row {
					sum[c] += v
				}
			}
			// Check frequency
			for i, v := range sum {
				freq := float64(v) / float64(iterations*tc.n)
				expected := tc.p[i]
				if expected+errSize < freq || expected-errSize > freq {
					t.Errorf("frequency (%f) is greater than expected (%f) +/- (%f)", freq, expected, errSize)
				}
			}
		})
	}
}

func TestMultinomialLog(t *testing.T) {
	cases := []struct {
		name string
		n    int
		p    []float64
	}{
		{name: "n=1,plen=2,dist=uniform",
			n: 1,
			p: []float64{0.5, 0.5},
		},
		{name: "n=1,plen=2,dist=skew_left",
			n: 1,
			p: []float64{0.9, 0.1},
		},
		{name: "n=1,plen=2,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.9},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0},
		},
		{name: "n=1,plen=3,dist=uniform",
			n: 1,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=1,plen=3,dist=skew_left",
			n: 1,
			p: []float64{0.6, 0.3, 0.1},
		}, {name: "n=1,plen=3,dist=skew_right",
			n: 1,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=1,plen=4,dist=uniform",
			n: 1,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=1,plen=10,dist=uniform",
			n: 1,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
		{name: "n=2,plen=2,dist=uniform",
			n: 2,
			p: []float64{0.5, 0.5},
		},
		{name: "n=2,plen=2,dist=skew_left",
			n: 2,
			p: []float64{0.9, 0.1},
		},
		{name: "n=2,plen=2,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.9},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0},
		},
		{name: "n=2,plen=3,dist=uniform",
			n: 2,
			p: []float64{0.33, 0.33, 0.34},
		},
		{name: "n=2,plen=3,dist=skew_left",
			n: 2,
			p: []float64{0.6, 0.3, 0.1},
		},
		{name: "n=2,plen=3,dist=skew_right",
			n: 2,
			p: []float64{0.1, 0.3, 0.6},
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: []float64{0.0, 0.0, 1.0},
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: []float64{1.0, 0.0, 0.0},
		},
		{name: "n=2,plen=4,dist=uniform",
			n: 2,
			p: []float64{0.25, 0.25, 0.25, 0.25},
		},
		{name: "n=2,plen=10,dist=uniform",
			n: 2,
			p: []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		},
	}
	iterations := 1000
	errSize := 0.05
	rand.Seed(0)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert to log probabilities
			logP := make([]float64, len(tc.p))
			for i := range tc.p {
				if tc.p[i] == 0 {
					logP[i] = math.Inf(-1)
				} else {
					logP[i] = math.Log(tc.p[i])
				}
			}
			// Simulate
			var cnt [][]int
			for i := 0; i < iterations; i++ {
				result := multinomialLog(tc.n, logP)
				cnt = append(cnt, result)
			}
			// Sum columns
			sum := make([]int, len(tc.p))
			for _, row := range cnt {
				for c, v := range row {
					sum[c] += v
				}
			}
			// Check frequency
			for i, v := range sum {
				freq := float64(v) / float64(iterations*tc.n)
				expected := tc.p[i]
				if expected+errSize < freq || expected-errSize > freq {
					t.Errorf("frequency (%f) is greater than expected (%f) +/- (%f)", freq, expected, errSize)
				}
			}
		})
	}
}
