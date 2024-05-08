package elasticsearch

import (
	"slices"
	"strings"
)

func ProductNameNgrams(name string, n int) []string {
	words := strings.Split(name, " ")
	firstWords := words[:min(n, len(words))]

	var substrings []string
	for i := 0; i < len(firstWords); i++ {
		for j := i + 1; j <= len(firstWords); j++ {
			substrings = append(substrings, strings.Join(firstWords[i:j], " "))
		}
	}
	if slices.Contains(substrings, name) {
		return substrings
	}
	return append(substrings, name)
}
