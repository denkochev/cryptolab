package main

import (
	"cryptolab/bigmath"
	"fmt"
)

func main() {

	bigInt := bigmath.BigInt{}

	hexStr := "e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7"

	// fmt.Println(hexStr)
	bigInt.SetHex(hexStr)

	result := bigInt.GetHex()

	fmt.Println(result)
}
