package elasticsearch_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/elasticsearch"
)

func TestProductNameNgrams(t *testing.T) {
	tests := []struct {
		name           string
		inputString    string
		numWords       int
		expectedOutput []string
	}{
		{
			name:           "Empty string",
			inputString:    "",
			numWords:       1,
			expectedOutput: []string{""},
		},
		{
			name:           "Single word",
			inputString:    "Apple",
			numWords:       2,
			expectedOutput: []string{"Apple"},
		},
		{
			name:           "Two words",
			inputString:    "Apple macbook",
			numWords:       4,
			expectedOutput: []string{"Apple", "Apple macbook", "macbook"},
		},
		{
			name:        "Multiple words",
			inputString: "Apple macbook 14 m1 pro",
			numWords:    4,
			expectedOutput: []string{
				"Apple",
				"Apple macbook",
				"Apple macbook 14",
				"Apple macbook 14 m1",
				"macbook",
				"macbook 14",
				"macbook 14 m1",
				"14",
				"14 m1",
				"m1",
				"Apple macbook 14 m1 pro",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := elasticsearch.ProductNameNgrams(tc.inputString, tc.numWords)
			require.ElementsMatch(t, tc.expectedOutput, actualOutput)
		})
	}
}
