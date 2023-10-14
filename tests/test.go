package main

import (
	"cryptolab/bigmath"
	"fmt"
	"strconv"
)

func main() {

	// bigInt := bigmath.BigInt{}
	// bigInt.SetHex("e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7")

	// ... Так само для b, c, d, e, f

	//fmt.Println(bigInt.GetHex()) // виведе: e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7

	myHex := "e035c6cfa42609b9"

	myMethod, _ := bigmath.HexToInt(myHex)
	fmt.Println(myMethod)

	num, _ := strconv.ParseUint(myHex, 16, 64)
	fmt.Println(num)
}
