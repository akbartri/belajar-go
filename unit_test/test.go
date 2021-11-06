package unit_test

import (
	"testing"
)

func MultipleNumber(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 Equal 4")
	}
}
