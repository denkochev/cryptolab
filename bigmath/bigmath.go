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

/*
BITWISE OPERATIONS
*/
func XOR(a, b BigInt) string {
	var length int
	/*
		проста оптимізація для приведення
		двух блоків до однакового розміру
		шляхом заповнення нулями старших бітів
		значення з меншою кількістю блоків
		(саме це дасть очікуваний результат при XOR)
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

func INV(a BigInt) string {
	fmt.Println(a.value)
	blocks := a.value
	/*
		Часто перший блок може фактично займати менше 64 бітів,
		якщо до такого блоку застосувати інверсію
		інвертованими також опиняться непотрібні лідуючі біти, які дорівнюють 0!
		для того, щоб цього не відбувалось перевіряємо фактичну кількість бітів першого блоку
	*/
	var iterator int
	actualBits := getBits(blocks[0])

	if actualBits <= 32 {
		// інвертую лише! дійсні біти
		var mask uint64 = ((1 << 64) - 1)
		blocks[0] = blocks[0] ^ mask
		iterator = 1
	} else {
		iterator = 0
	}

	for i := iterator; i < len(blocks); i++ {
		blocks[i] = ^blocks[i]
	}
	fmt.Println(blocks)
	var hexStr string
	for _, val := range blocks {
		hexStr += fmt.Sprintf("%02x", val)
	}

	return hexStr
	// return trimLeadingZeros(blocksToHex(blocks))
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

func intoBlocks(hexStr string) []uint64 {
	// Конвертація рядка до числового формату (uint64)

	binaryLength := len(hexStr) / 2
	binary := make([]uint64, binaryLength)
	for i := 0; i < binaryLength; i++ {
		h := hexStr[i*2 : i*2+2]
		msb, _ := SingleHexToInt(h[0])
		msb *= 16
		lsb, _ := SingleHexToInt(h[1])
		binary[i] = uint64(255 - (msb + lsb))
	}

	return binary
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
