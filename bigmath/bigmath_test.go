package bigmath

import (
	"testing"
)

func TestBigIntPow(t *testing.T) {
	var tests = []struct {
		val, result uint64
		pow         int
	}{
		{2, 2, 1},
		{2, 4, 2},
		{2, 8, 3},
		{2, 16, 4},
		{3, 3, 1},
		{3, 9, 2},
		{3, 27, 3},
		// можна додати більше тестових випадків за бажанням
	}

	for _, tt := range tests {
		res := BigIntPow(tt.val, tt.pow)
		if res != tt.result {
			t.Errorf("For value=%d and power=%d, expected %d, but got %d", tt.val, tt.pow, tt.result, res)
		}
	}
}

func TestHexByteToInt(t *testing.T) {
	tests := []struct {
		input    byte
		expected int
		hasError bool
	}{
		{'0', 0, false},
		{'1', 1, false},
		{'9', 9, false},
		{'A', 10, false},
		{'a', 10, false},
		{'F', 15, false},
		{'f', 15, false},
		{'G', 0, true}, // недійсний символ
		{'z', 0, true}, // недійсний символ
	}

	for _, tt := range tests {
		result, err := SingleHexToInt(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("Expected error for input %c, but got none", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for input %c, but got %s", tt.input, err.Error())
			}
			if result != tt.expected {
				t.Errorf("For input %c, expected %d, but got %d", tt.input, tt.expected, result)
			}
		}
	}
}
