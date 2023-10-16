package main

import (
	"cryptolab/bigmath"
	"fmt"
)

func main() {

	numberA := bigmath.BigInt{}
	//numberB := bigmath.BigInt{}
	// hexA := "d9eeffb8121dfac05c3512fd"
	// hexB := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	hexA := "00a05"
	//hexB := "146C0D"
	numberA.SetHex(hexA)
	//numberB.SetHex(hexB)

	fmt.Println(numberA.GetHex())

	//result_XOR := bigmath.XOR(numberA, numberB)
	//fmt.Println(result_XOR)

	// result_INV := bigmath.INV(numberB)
	// fmt.Println(result_INV == "ae409f7beb52a8d95c3e413f670884e4ab004d878072ad758b3e280219b8f15b")

	// result_INV = bigmath.INV(numberA)
	// fmt.Println(result_INV == "26110047ede2053fa3caed02")

	// numberC := bigmath.BigInt{}
	// numberC.SetHex("1eb93f2")
	// result_INV := bigmath.INV(numberC)
	// fmt.Println(result_INV == "e146c0d", "  actual -> ", result_INV, " needed -> e146c0d")

}
