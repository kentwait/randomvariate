package randomvariate

import (
	"math/rand"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestPoisson(t *testing.T) {
	cases := []struct {
		name   string
		lambda float64
	}{
		{name: "exp=-2",
			lambda: 1e-2,
		},
		{name: "exp=-1",
			lambda: 1e-1,
		},
		{name: "exp=0",
			lambda: 1e0,
		},
		{name: "exp=1",
			lambda: 1e1,
		},
		{name: "exp=2",
			lambda: 1e2,
		},
	}
	iterations := 100000
	errSize := 0.1
	rand.Seed(0)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Simulate
			var cnt []float64
			for i := 0; i < iterations; i++ {
				result := Poisson(tc.lambda)
				cnt = append(cnt, float64(result))
			}
			// Sum columns
			var sum int
			for _, v := range cnt {
				sum += int(v)
			}
			// Check mean and variance
			mean := float64(sum) / float64(iterations)
			variance, _ := stats.Variance(cnt)
			expected := tc.lambda
			err := expected * errSize
			if expected+err <= mean || expected-err >= mean {
				t.Errorf("frequency mean (%f) is greater than expected (%f) +/- (%f)", mean, expected, err)
			}
			if expected+err <= variance || expected-err >= variance {
				t.Errorf("frequency variance (%f) is greater than expected (%f) +/- (%f)", variance, expected, err)
			}
		})
	}
}
