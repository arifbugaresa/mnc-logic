package logic

import (
	"fmt"
	"strings"
)

func FindMatch(n int, stringsList []string) string {
	// Convert all strings to lowercase for case-insensitive comparison
	for i := 0; i < n; i++ {
		stringsList[i] = strings.ToLower(stringsList[i])
	}

	// Loop to compare each pair of strings
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stringsList[i] == stringsList[j] {
				return fmt.Sprintf("%d %d", i+1, j+1)
			}
		}
	}

	return "false"
}
