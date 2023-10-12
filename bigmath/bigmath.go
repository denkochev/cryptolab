package bigmath

import (
	"errors"
	"fmt"
)

type BigInt struct {
	hex   string
	value []uint64
}

func (b *BigInt) SetHex(hexValue string) {
	b.hex = hexValue

	for i := 0; i < len(hexValue); i += 16 {
		padding := i + 16
		if padding > len(hexValue) {
			padding = len(hexValue)
		}
		uintValue, err := HexToInt(hexValue[i:padding])
		if err != nil {
			fmt.Println("convertation was failed:", err)
			return
		}

		b.value = append(b.value, uint64(uintValue))
	}

	fmt.Println(b.value)
}

func (b *BigInt) GetHex() string {
	return b.hex
}

func HexToInt(hex string) (uint64, error) {
	if len(hex) > 16 {
		return 0, errors.New("overload 64-bit size")
	}
	fmt.Println(hex, len(hex))
	var result uint64

	for _, char := range hex {
		result = result << 4 // Переміщення на 4 біти вліво
		switch {
		case '0' <= char && char <= '9':
			result += uint64(char - '0')
		case 'a' <= char && char <= 'f':
			result += uint64(char - 'a' + 10)
		case 'A' <= char && char <= 'F':
			result += uint64(char - 'A' + 10)
		default:
			return 0, fmt.Errorf("invalid hex character: %c", char)
		}
	}

	return result, nil
}
