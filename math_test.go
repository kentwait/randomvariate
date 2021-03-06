package randomvariate

import "testing"

func TestRound(t *testing.T) {
	cases := []struct {
		name     string
		value    float64
		e        int
		expected float64
		epsilon  float64
	}{
		// no change
		// positive values
		{name: "value=0.0,unit=1e-1,expected=0.0,epsilon=1e-10",
			value:    0.0,
			e:        -1,
			expected: 0.0,
			epsilon:  1e-10,
		},
		{name: "value=1e-1,unit=1e-1,expected=1e-1,epsilon=1e-10",
			value:    1e-1,
			e:        -1,
			expected: 1e-1,
			epsilon:  1e-10,
		},
		{name: "value=1e-10,unit=1e-10,expected=1e-10,epsilon=1e-20",
			value:    1e-10,
			e:        -10,
			expected: 1e-10,
			epsilon:  1e-20,
		},
		{name: "value=1e1,unit=1e1,expected=1e1,epsilon=1e-10",
			value:    1e1,
			e:        1,
			expected: 1e1,
			epsilon:  1e-10,
		},
		{name: "value=1e10,unit=1e10,expected=1e10,epsilon=1e-10",
			value:    1e10,
			e:        10,
			expected: 1e10,
			epsilon:  1e-10,
		},
		// no change
		// negative values
		{name: "value=-1e-1,unit=1e-1,expected=-1e-1,epsilon=1e-10",
			value:    -1e-1,
			e:        -1,
			expected: -1e-1,
			epsilon:  1e-10,
		},
		{name: "value=-1e-10,unit=1e-10,expected=-1e-10,epsilon=1e-20",
			value:    -1e-10,
			e:        -10,
			expected: -1e-10,
			epsilon:  1e-20,
		},
		{name: "value=-1e1,unit=1e1,expected=-1e1,epsilon=1e-10",
			value:    -1e1,
			e:        1,
			expected: -1e1,
			epsilon:  1e-10,
		},
		{name: "value=-1e10,unit=1e10,expected=-1e10,epsilon=1e-10",
			value:    -1e10,
			e:        10,
			expected: -1e10,
			epsilon:  1e-10,
		},
		// round down at different place values
		// positive values
		{name: "value=1.2e-1,unit=1e-1,expected=1e-1,epsilon=1e-10",
			value:    1.2e-1,
			e:        -1,
			expected: 1e-1,
			epsilon:  1e-10,
		},
		{name: "value=1.2e-10,unit=1e-10,expected=1e-10,epsilon=1e-20",
			value:    1.2e-10,
			e:        -10,
			expected: 1e-10,
			epsilon:  1e-20,
		},
		{name: "value=1.2e1,unit=1e1,expected=1e1,epsilon=1e-10",
			value:    1.2e1,
			e:        1,
			expected: 1e1,
			epsilon:  1e-10,
		},
		{name: "value=1.2e10,unit=1e10,expected=1e10,epsilon=1e-10",
			value:    1.2e10,
			e:        10,
			expected: 1e10,
			epsilon:  1e-10,
		},
		// round down at different place values
		// negative values
		{name: "value=-1.2e-1,unit=1e-1,expected=-1e-1,epsilon=1e-10",
			value:    -1.2e-1,
			e:        -1,
			expected: -1e-1,
			epsilon:  1e-10,
		},
		{name: "value=-1.2e-10,unit=1e-10,expected=-1e-10,epsilon=1e-20",
			value:    -1.2e-10,
			e:        -10,
			expected: -1e-10,
			epsilon:  1e-20,
		},
		{name: "value=-1.2e1,unit=1e1,expected=-1e1,epsilon=1e-10",
			value:    -1.2e1,
			e:        1,
			expected: -1e1,
			epsilon:  1e-10,
		},
		{name: "value=-1.2e10,unit=1e10,expected=-1e10,epsilon=1e-20",
			value:    -1.2e10,
			e:        10,
			expected: -1e10,
			epsilon:  1e-20,
		},
		// round up at different place values
		// positive values
		{name: "value=0.8e-1,unit=1e-1,expected=1e-1,epsilon=1e-10",
			value:    0.8e-1,
			e:        -1,
			expected: 1e-1,
			epsilon:  1e-10,
		},
		{name: "value=0.8e-10,unit=1e-10,expected=1e-10,epsilon=1e-20",
			value:    0.8e-10,
			e:        -10,
			expected: 1e-10,
			epsilon:  1e-20,
		},
		{name: "value=0.8e1,unit=1e1,expected=1e1,epsilon=1e-10",
			value:    0.8e1,
			e:        1,
			expected: 1e1,
			epsilon:  1e-10,
		},
		{name: "value=0.8e10,unit=1e10,expected=1e10,epsilon=1e-20",
			value:    0.8e10,
			e:        10,
			expected: 1e10,
			epsilon:  1e-20,
		},
		// round up at different place values
		// negative values
		{name: "value=-0.8e-1,unit=1e-1,expected=-1e-1,epsilon=1e-10",
			value:    -0.8e-1,
			e:        -1,
			expected: -1e-1,
			epsilon:  1e-10,
		},
		{name: "value=-0.8e-10,unit=1e-10,expected=-1e-10,epsilon=1e-20",
			value:    -0.8e-10,
			e:        -10,
			expected: -1e-10,
			epsilon:  1e-20,
		},
		{name: "value=-0.8e1,unit=1e1,expected=-1e1,epsilon=1e-10",
			value:    -0.8e1,
			e:        1,
			expected: -1e1,
			epsilon:  1e-10,
		},
		{name: "value=-0.8e10,unit=1e10,expected=-1e10,epsilon=1e-20",
			value:    -0.8e10,
			e:        10,
			expected: -1e10,
			epsilon:  1e-20,
		},
		// 0.5 value at different place values
		{name: "value=5e-1,unit=1e0,expected=1e0,epsilon=1e-10",
			value:    5e-1,
			e:        0,
			expected: 1e0,
			epsilon:  1e-10,
		},
		{name: "value=5e-10,unit=1e-9,expected=1e-9,epsilon=1e-20",
			value:    5e-10,
			e:        -9,
			expected: 1e-9,
			epsilon:  1e-20,
		},
		{name: "value=5e0,unit=1e1,expected=1e1,epsilon=1e-10",
			value:    5e0,
			e:        1,
			expected: 1e1,
			epsilon:  1e-10,
		},
		{name: "value=5e9,unit=1e10,expected=1e10,epsilon=1e-20",
			value:    5e9,
			e:        10,
			expected: 1e10,
			epsilon:  1e-20,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := Round(tc.value, tc.e)
			if v > tc.expected+tc.epsilon || v < tc.expected-tc.epsilon {
				t.Errorf("expected value is %e, instead got %e", tc.expected, v)
			}
		})
	}
}
