package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClockAngle(t *testing.T) {
	tests := []struct {
		name   string
		hour   int
		minute int

		hasError bool
		result   int
	}{
		{
			name:   "happy path angle hour major",
			hour:   9,
			minute: 20,
			result: 150,
		},
		{
			name:   "happy path angle minute major",
			hour:   1,
			minute: 30,
			result: 150,
		},
		{
			name:   "happy path 00:01hrs",
			hour:   12,
			minute: 1,
			result: 6,
		},
		{
			name:     "invalid hour",
			hour:     90,
			minute:   20,
			hasError: true,
		},
		{
			name:     "invalid minute",
			hour:     5,
			minute:   120,
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, err := clockAngle(test.hour, test.minute)

			if (err != nil) != test.hasError {
				t.Fatal(err)
			}
			if test.hasError {
				return
			}

			assert.Equal(t, test.result, r)
		})
	}
}
