package helper

import (
	"sort"
	"strings"
)

func SortSuggests(suggests []string) []string {
	sort.Slice(suggests, func(i int, j int) bool {
		wordsI := len(strings.Split(suggests[i], " "))
		wordsJ := len(strings.Split(suggests[j], " "))
		if wordsI != wordsJ {
			return wordsI < wordsJ
		}
		return len(suggests[i]) < len(suggests[j])
	})
	return suggests
}
