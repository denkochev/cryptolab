package main

import (
	"cryptolab/bigmath"
	"fmt"
)

func main() {
	// EXAMPLE OF USAGE
	numberA := bigmath.BigInt{}
	numberB := bigmath.BigInt{}

	hexA := "4f1df3530cb621949a20b5cb4c25532ec0d96ef880c91ec3"
	hexB := "4d602e11c8481fe86748041f"

	numberA.SetHex(hexA)
	numberB.SetHex(hexB)

	fmt.Println("hex A ", numberA.GetHex())
	// Output: 4f1df3530cb621949a20b5cb4c25532ec0d96ef880c91ec3
	fmt.Println("hex B ", numberB.GetHex())
	// Output: 4d602e11c8481fe86748041f

	// func that return blocks of uint64
	blocksA := numberA.GetBlocks()
	fmt.Println(blocksA)
	blocksB := numberB.GetBlocks()
	fmt.Println(blocksB)

	invertedA := bigmath.INV(numberA)
	invertedB := bigmath.INV(numberB)
	fmt.Println("INV a - ", invertedA)
	fmt.Println("INV b - ", invertedB)

	numbers_XOR := bigmath.XOR(numberA, numberB)
	fmt.Println("A xor B - ", numbers_XOR)

	numbers_OR := bigmath.OR(numberA, numberB)
	fmt.Println("A or B - ", numbers_OR)

	numbers_AND := bigmath.AND(numberA, numberB)
	// A and b without leading zeros
	numbers_AND_cleaned := bigmath.TrimLeadingZeros(bigmath.AND(numberA, numberB))
	fmt.Println("A and B - ", numbers_AND)
	fmt.Println("A and B cleaned - ", numbers_AND_cleaned)

	a_shiftR := bigmath.ShiftR(numberA, 4)
	fmt.Println("A >> 4 - ", a_shiftR)

	b_shiftL := bigmath.ShiftL(numberB, 3)
	fmt.Println("B << 3 - ", b_shiftL)

	numbers_ADD := bigmath.ADD(numberA, numberB)
	fmt.Println("A add B - ", numbers_ADD)

	numbers_SUB := bigmath.SUB(numberA, numberB)
	fmt.Println("A sub B - ", numbers_SUB)

	a_mod := bigmath.MOD(numberA, 5)
	fmt.Println("A mod 5 - ", a_mod)
}
