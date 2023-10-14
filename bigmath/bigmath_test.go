package bigmath

import (
	"errors"
	"fmt"
	"testing"
)

func TestHexToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
		err      error
	}{
		{"0", 0, nil},
		{"a", 10, nil},
		{"A", 10, nil},
		{"f", 15, nil},
		{"F", 15, nil},
		{"10", 16, nil},
		{"1a", 26, nil},
		{"1A", 26, nil},
		{"ABCDEF", 11259375, nil},
		{"abcdef", 11259375, nil},
		{"FFFFFFFFFFFFFFFF", 18446744073709551615, nil},
		{"G", 0, fmt.Errorf("invalid hex character: G")},
		{"10000000000000000", 0, errors.New("overload 64-bit size")},
		{"00", 0, nil},
		{"000a", 10, nil},
		{"00A", 10, nil},
		{"1aBc", 6844, nil},
		{"1AbC", 6844, nil},
		{"", 0, nil}, // empty string should return 0
		{"1g", 0, fmt.Errorf("invalid hex character: g")},
		{"1H", 0, fmt.Errorf("invalid hex character: H")},
		{"1a3fG", 0, fmt.Errorf("invalid hex character: G")},
		{"FFFFFFFFFFFFFFF", 1152921504606846975, nil},  // Just below the 16-character limit
		{"1234567890abcdef", 1311768467294899695, nil}, // 16 characters, all valid
		{"1a3f4Z", 0, fmt.Errorf("invalid hex character: Z")},
		{"ZZZZ", 0, fmt.Errorf("invalid hex character: Z")},
		{"0000000q00000000", 0, fmt.Errorf("invalid hex character: q")},
		{"FF", 255, nil},              // Maximum 1-byte value
		{"FFFF", 65535, nil},          // Maximum 2-byte value
		{"FFFFFF", 16777215, nil},     // Maximum 3-byte value
		{"FFFFFFFF", 4294967295, nil}, // Maximum 4-byte value
		{"1ABCDEF1", 448585457, nil},
		{"000000000000000Z", 0, fmt.Errorf("invalid hex character: Z")},
		{"FFFFFFFFFFFFFFFE", 18446744073709551614, nil}, // 1 less than the maximum 64-bit value
		{"1234567z", 0, fmt.Errorf("invalid hex character: z")},
		{"abcdEFABabcdEFAB", 12379814471312207787, nil}, // 16 characters, all valid

	}

	for _, test := range tests {
		result, err := HexToInt(test.input)
		if result != test.expected || !equalError(err, test.err) {
			t.Errorf("For input %s expected %d and error %v, but got %d and error %v", test.input, test.expected, test.err, result, err)
		}
	}
}

func equalError(a, b error) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Error() == b.Error()
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
