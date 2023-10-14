package bigmath

import (
	"errors"
	"fmt"
)

type BigInt struct {
	value [4]uint64
}

func (b *BigInt) SetHex(hexValue string) {
	b.value = splitIntoBlocks(hexValue)
}

func (b *BigInt) GetHex() string {
	return blocksToHex(b.value)
}

func splitIntoBlocks(hexStr string) [4]uint64 {
	// Конвертація рядка до числового формату (uint64)

	var blocks [4]uint64
	length := len(hexStr)

	for i := 3; i >= 0; i-- {
		// Витягуємо 16 символів (що відповідає 64 бітам) з кінця рядка
		start := length - 16
		if start < 0 {
			start = 0
		}
		subStr := hexStr[start:length]

		// Конвертуємо ці символи до uint64
		num, _ := HexToInt(subStr)
		blocks[i] = num

		// Переміщуємося до наступного блоку
		length -= 16
		if length < 0 {
			break
		}
	}

	return blocks
}

func blocksToHex(blocks [4]uint64) string {
	var hexString string
	for _, block := range blocks {
		hexString += fmt.Sprintf("%016x", block)
	}
	return hexString
}

func HexToInt(hex string) (uint64, error) {
	if len(hex) > 16 {
		return 0, errors.New("overload 64-bit size")
	}

	var result uint64

	for _, char := range hex {
		result = result << 4 // Переміщення на 4 біти вліво, резервуємо одне число 16 розряду -> 0001 0000
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

func SingleHexToInt(symbol byte) (int, error) {
	switch {
	case '0' <= symbol && symbol <= '9':
		return int(symbol - '0'), nil
	case 'a' <= symbol && symbol <= 'f':
		return int(symbol - 'a' + 10), nil
	case 'A' <= symbol && symbol <= 'F':
		return int(symbol - 'A' + 10), nil
	default:
		return 0, fmt.Errorf("invalid hex character: %c", symbol)
	}
}
