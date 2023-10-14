package main

import (
	"cryptolab/bigmath"
	"fmt"
)

func main() {

	numberA := bigmath.BigInt{}
	numberB := bigmath.BigInt{}
	hexA := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	hexB := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	numberA.SetHex(hexA)
	numberB.SetHex(hexB)

	result := bigmath.XOR(numberA, numberB)

	fmt.Println(result)
	fmt.Println(result == "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8")
}
