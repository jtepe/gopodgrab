package cmd

import (
	"testing"
)

func TestHumanized(t *testing.T) {
	tests := map[string]struct {
		in       int64
		expected string
	}{
		"Zero":                {in: 0, expected: "0.00 B"},
		"Under one kilo":      {in: 872, expected: "872.00 B"},
		"Exactly one kilo":    {in: 1024, expected: "1.00 KB"},
		"A low kilo":          {in: 1672, expected: "1.63 KB"},
		"A few hundred kilo":  {in: 792141, expected: "773.58 KB"},
		"Just under one mega": {in: 1033575, expected: "1009.35 KB"},
		"Exactly one mega":    {in: 1024 * 1024, expected: "1.00 MB"},
		"A low mega":          {in: 1552872, expected: "1.48 MB"},
		"A few hundred mega":  {in: 471552872, expected: "449.71 MB"},
		"Just under one giga": {in: 1011552872, expected: "964.69 MB"},
		"Exactly one giga":    {in: 1024 * 1024 * 1024, expected: "1.00 GB"},
		"Over one giga":       {in: 3862248721, expected: "3.60 GB"},
	}

	for name, test := range tests {
		if res := humanized(test.in); res != test.expected {
			t.Errorf("%s: for %d got %s, but expected %s", name, test.in, res, test.expected)
		}
	}
}
