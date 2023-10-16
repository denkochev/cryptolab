package bigmath

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Record struct {
	HexNumber1 string
	HexNumber2 string
	XORResult  string
}

type Inv struct {
	BeforeHEX string
	AfterHEX  string
}

type Hex struct {
	hex string
}

func TestINV(t *testing.T) {
	tests, _ := readCSV_INV("./test_assets/hex_inversion_dataset.csv")

	for _, test := range tests {
		testA := BigInt{}

		testA.SetHex(test.BeforeHEX)

		result := INV(testA)

		if result != test.AfterHEX {
			t.Errorf("For input %s expected %s, but got %s", test.BeforeHEX, test.AfterHEX, result)
		}
	}
}

func TestXOR(t *testing.T) {
	tests, _ := readCSV("./test_assets/hex_xor_dataset.csv")

	for _, test := range tests {
		testA := BigInt{}
		testB := BigInt{}

		testA.SetHex(test.HexNumber1)
		testB.SetHex(test.HexNumber2)

		result := XOR(testA, testB)

		if result != test.XORResult {
			t.Errorf("For input %s and %s expected %s, but got %s", test.HexNumber1, test.HexNumber2, test.XORResult, result)
		}
	}
}

func TestGetSet(t *testing.T) {
	tests, _ := readCSV_HEX("./test_assets/hex_values.csv.csv")
	for _, test := range tests {

		testNumber := BigInt{}
		testNumber.SetHex(test.hex)
		result := testNumber.GetHex()

		if result != strings.ToLower(test.hex) {
			t.Errorf("For input %s expected %s, but got %s", test.hex, strings.ToLower(test.hex), result)
		}
	}
}

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

func readCSV(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Record
	for _, line := range lines[1:] { // Skip header
		records = append(records, Record{
			HexNumber1: line[0],
			HexNumber2: line[1],
			XORResult:  line[2],
		})
	}
	return records, nil
}

func readCSV_INV(filename string) ([]Inv, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Inv
	for _, line := range lines[1:] { // Skip header
		records = append(records, Inv{
			BeforeHEX: line[0],
			AfterHEX:  line[1],
		})
	}
	return records, nil
}

func readCSV_HEX(filename string) ([]Hex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Hex
	for _, line := range lines[1:] { // Skip header
		records = append(records, Hex{
			hex: line[0],
		})
	}
	return records, nil
}
