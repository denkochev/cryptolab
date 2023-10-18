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
	var hex string = hexValue
	if len(hex)%2 != 0 {
		hex = "0" + hexValue
	}
	b.value = splitIntoBlocks(hex)
}

func (b *BigInt) GetHex() string {
	return trimLeadingZeros(blocksToHex(b.value))
}

func (b *BigInt) GetBlocks() []uint64 {
	return b.value
}

/*
ARITHMETIC OPERATIONS
*/
func ADD(a, b BigInt) string {
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

	var over uint64 = 0
	for i := length - 1; i >= 0; i-- {
		temp_sum := a.value[i] + b.value[i] + over
		blocks[i] = temp_sum & 0xFF
		over = temp_sum >> 8
	}
	if over != 0 {
		blocks = append([]uint64{over}, blocks...)
	}

	return trimLeadingZeros(blocksToHex(blocks))
}

func SUB(a, b BigInt) string {
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

	blocks := make([]int, length)

	var over int = 0
	for i := length - 1; i >= 0; i-- {
		temp_sum := int(a.value[i]) - int(b.value[i]) - over

		if temp_sum < 0 {
			temp_sum += (0xFF + 1)
			over = 1
		} else {
			over = 0
		}

		blocks[i] = temp_sum
	}
	if over != 0 {
		blocks = append([]int{over}, blocks...)
	}

	var uints []uint64
	for _, v := range blocks {
		uints = append(uints, uint64(v))
	}

	return trimLeadingZeros(blocksToHex(uints))
}

func MOD(a BigInt, modulus uint64) uint64 {
	var result uint64 = 0

	blocks := a.value
	for i := 0; i < len(blocks); i++ {
		result = ((result << 32) + blocks[i]) % modulus
	}
	return result
}

/*
BITWISE OPERATIONS
*/
func INV(a BigInt) string {
	blocks := a.value

	for i := 0; i < len(blocks); i++ {
		// маска на 8 бітів, для інвертування кожного блоку
		var mask uint64 = ((1 << 8) - 1)
		blocks[i] = mask ^ blocks[i]
	}
	return blocksToHex(blocks)
}

func XOR(a, b BigInt) string {
	var length int
	/*
		проста оптимізація для приведення
		двух блоків до однакового розміру
		шляхом заповнення нулями старших бітів
		значення з меншою кількістю блоків
		(саме це дасть очікуваний результат при бітових операціях з двума гекс числами)
	*/
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

func OR(a, b BigInt) string {
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
		blocks[i] = a.value[i] | b.value[i]
	}
	return trimLeadingZeros(blocksToHex(blocks))
}

func AND(a, b BigInt) string {
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
		blocks[i] = a.value[i] & b.value[i]
	}
	// not triming leading zeros!
	return blocksToHex(blocks)
}

func ShiftL(a BigInt, bitsShift int) string {
	blocks := a.value
	// змінна для переносу зайвих бітів
	var overflow uint64 = 0
	for i := len(blocks) - 1; i >= 0; i-- {
		shifted_value := (blocks[i] << bitsShift) + overflow
		overflow = shifted_value >> 8
		blocks[i] = shifted_value & 0xFF
	}

	for overflow > 0 {
		blocks = append([]uint64{overflow & 0xFF}, blocks...)
		overflow >>= 8
	}

	return trimLeadingZeros(blocksToHex(blocks))
}

func ShiftR(a BigInt, bitsShift int) string {
	blocks := a.value
	// змінна для переносу зайвих бітів
	var overflow uint64 = 0
	for i := 0; i < len(blocks); i++ {
		combined := (blocks[i] + (overflow << 8))
		blocks[i] = (combined >> bitsShift) & 0xFF
		overflow = combined & ((1 << bitsShift) - 1)

	}

	for overflow > 0 {
		blocks = append(blocks, overflow&0xFF)
		overflow >>= 8
	}

	for len(blocks) > 1 && blocks[0] == 0 {
		blocks = blocks[1:]
	}

	return trimLeadingZeros(blocksToHex(blocks))
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
	var amountOfBLocks int = (len(hexStr) / 2)

	blocks := make([]uint64, amountOfBLocks)

	for i := 0; i < amountOfBLocks; i++ {
		subStr := hexStr[i*2 : i*2+2]
		// Конвертуємо ці символи до uint64
		num, _ := HexToInt(subStr)
		blocks[i] = num
	}

	return blocks
}

func blocksToHex(blocks []uint64) string {
	var hexString string
	for _, val := range blocks {
		hexString += fmt.Sprintf("%02x", val)
	}
	return hexString
}

func trimLeadingZeros(s string) string {
	i := 0
	for ; i < len(s) && s[i] == '0'; i++ {
	}
	result := s[i:]
	if len(result) == 0 {
		return "0"
	}
	return result
}

func getBits(num uint64) int {
	if num == 0 {
		return 1 // 0 requires 1 bit
	}

	bitsUsed := 0
	for num > 0 {
		bitsUsed++
		num >>= 1
	}
	return bitsUsed
}
