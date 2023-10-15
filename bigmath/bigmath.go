package bigmath

import (
	"errors"
	"fmt"
)

/*
API FOR LIBRARY
*/
type BigInt struct {
	value []uint64
}

func (b *BigInt) SetHex(hexValue string) {
	b.value = splitIntoBlocks(hexValue)
}

func (b *BigInt) GetHex() string {
	return trimLeadingZeros(blocksToHex(b.value))
}

/*
BITWISE OPERATIONS
*/
func XOR(a, b BigInt) string {
	var length int
	if len(a.value) == len(b.value) {
		length = len(a.value)
	} else if len(a.value) > len(b.value) {
		length = len(a.value)
		blocksNeeded := len(a.value) - len(b.value)
		b.value = append(make([]uint64, blocksNeeded), b.value...)
	} else {
		length = len(b.value)
		blocksNeeded := len(b.value) - len(a.value)
		a.value = append(make([]uint64, blocksNeeded), a.value...)
	}

	blocks := make([]uint64, length)

	for i := length - 1; i >= 0; i-- {
		blocks[i] = a.value[i] ^ b.value[i]
	}

	return trimLeadingZeros(blocksToHex(blocks))
}

func INV(a BigInt) string {
	fmt.Println(a.value)
	blocks := a.value

	for i := 0; i < len(a.value); i++ {
		blocks[i] = ^blocks[i]
	}
	return blocksToHex(blocks)
}

/*
PARSING FUNCTIONS
*/
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

/*
HELPERS FUNCTIONS
*/
func splitIntoBlocks(hexStr string) []uint64 {
	// Конвертація рядка до числового формату (uint64)

	var amountOfBLocks int = (len(hexStr) / 16)
	if len(hexStr)%16 > 0 {
		amountOfBLocks += 1
	}

	blocks := make([]uint64, amountOfBLocks)
	length := len(hexStr)

	for i := amountOfBLocks - 1; i >= 0; i-- {
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

func blocksToHex(blocks []uint64) string {
	var hexString string
	for _, block := range blocks {
		hexString += fmt.Sprintf("%016x", block)
	}
	return hexString
}

func trimLeadingZeros(s string) string {
	i := 0
	for ; i < len(s) && s[i] == '0'; i++ {
	}
	return s[i:]
}
