package logic

import (
	"testing"
)

func TestFindMatch(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		strings  []string
		expected string
	}{
		{
			name:     "testcase 1",
			n:        4,
			strings:  []string{"abcd", "acbd", "aaab", "acbd"},
			expected: "2 4",
		},
		{
			name:     "testcase 2",
			n:        5,
			strings:  []string{"pisang", "goreng", "enak", "sekali", "RASANYA"},
			expected: "false",
		},
		{
			name:     "testcase 3",
			n:        11,
			strings:  []string{"satu", "sate", "tujuh", "tusuk", "tujuh", "Sate", "Bonus", "tiga", "Puluh", "Tujuh", "Tusuk"},
			expected: "2 6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMatch(tt.n, tt.strings)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}
