package main

import (
	"cryptolab/bigmath"
	"fmt"
)

type BigInt struct {
	hex   string
	value []uint64
}

func (b *BigInt) SetHex(hexValue string) (uint64, error) {
	var result uint64

	for _, char := range hexValue {
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

func main() {
	// b := &BigInt{}
	// val, err := b.SetHex("e035c6cfa42609b9")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Converted value:", val)

	bigInt := bigmath.BigInt{}
	bigInt.SetHex("e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7")
	//fmt.Println(bigInt.GetHex()) // виведе: e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7
}
