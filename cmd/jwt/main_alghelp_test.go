package main

import (
	"strings"
	"testing"
	"github.com/golang-jwt/jwt/v5"
	"sort"
)







func TestalgHelp(t *testing.T) {

	expectedAlgs := jwt.GetAlgorithms()
	sort.Strings(expectedAlgs)

	t.Run("Normal Operation", func(t *testing.T) {
		result := algHelp()

		for _, alg := range expectedAlgs {
			if !strings.Contains(result, alg) {
				t.Errorf("expected %v to include %q, but it did not", result, alg)
			}
		}
		t.Log("Normal operation checked")
	})

	t.Run("Formatting", func(t *testing.T) {
		result := algHelp()

		identifiers := strings.Split(result, ",\n")
		for i, identifier := range identifiers {
			if i%7 == 0 && i > 0 && len(identifier) > 0 {
				t.Errorf("expected a comma and a newline after the 7th identifier, but got %q", identifier)
			}
		}
		t.Log("Formatting checked")
	})

	t.Run("Sorting", func(t *testing.T) {
		result := algHelp()

		identifiers := strings.Split(result, ", ")
		if !sort.StringsAreSorted(identifiers) {
			t.Errorf("expected identifiers to be sorted, but they were not: %v", identifiers)
		}
		t.Log("Sorting checked")
	})
}
