package randomvariate

import (
	"math/rand"
	"testing"
)

func TestBinomial(t *testing.T) {
	cases := []struct {
		name string
		n    int
		p    float64
	}{
		{name: "n=1,plen=2,dist=uniform",
			n: 1,
			p: 0.5,
		},
		{name: "n=1,plen=2,dist=skew_left",
			n: 1,
			p: 0.9,
		},
		{name: "n=1,plen=2,dist=skew_right",
			n: 1,
			p: 0.1,
		},
		{name: "n=1,plen=2,dist=zero_left",
			n: 1,
			p: 0.0,
		},
		{name: "n=1,plen=2,dist=zero_right",
			n: 1,
			p: 1.0,
		},
		{name: "n=2,plen=2,dist=uniform",
			n: 2,
			p: 0.5,
		},
		{name: "n=2,plen=2,dist=skew_left",
			n: 2,
			p: 0.9,
		},
		{name: "n=2,plen=2,dist=skew_right",
			n: 2,
			p: 0.1,
		},
		{name: "n=2,plen=2,dist=zero_left",
			n: 2,
			p: 0.0,
		},
		{name: "n=2,plen=2,dist=zero_right",
			n: 2,
			p: 1.0,
		},
	}
	iterations := 1000
	errSize := 0.05
	rand.Seed(0)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Simulate
			var cnt []int
			for i := 0; i < iterations; i++ {
				result := Binomial(tc.n, tc.p)
				cnt = append(cnt, result)
			}
			// Sum columns
			var sum int
			for _, v := range cnt {
				sum += v
			}
			// Check frequency
			freq := float64(sum) / float64(iterations*tc.n)
			expected := tc.p
			if expected+errSize < freq || expected-errSize > freq {
				t.Errorf("frequency (%f) is greater than expected (%f) +/- (%f)", freq, expected, errSize)
			}
		})
	}
}
