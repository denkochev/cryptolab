This is a simple library for working with huge HEX numbers in GO/Golang. 

Example of usage API:
```
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
```

Output: 
```
hex A  4f1df3530cb621949a20b5cb4c25532ec0d96ef880c91ec3
hex B  4d602e11c8481fe86748041f
[79 29 243 83 12 182 33 148 154 32 181 203 76 37 83 46 192 217 110 248 128 201 30 195]
[77 96 46 17 200 72 31 232 103 72 4 31]
INV a -  b0e20cacf349de6b65df4a34b3daacd13f2691077f36e13c
INV b -  b29fd1ee37b7e01798b7fbe0
A xor B -  4f1df3530cb621949a20b5cb01457d3f08917110e7811adc
A or B -  4f1df3530cb621949a20b5cb4d657f3fc8d97ff8e7c91edf
A and B -  0000000000000000000000004c200200c0480ee800480403
A and B cleaned -  4c200200c0480ee800480403
A >> 4 -  4f1df3530cb621949a20b5cb4c25532ec0d96ef880c91ec03
B << 3 -  26b01708e4240ff433a4020f8
A add B -  4f1df3530cb621949a20b5cb9985814089218ee0e81122e2
A sub B -  4f1df3530cb621949a20b5cafec5251cf8914f1019811aa4
A mod 5 -  2
```

All functions tested with valid data.
Test cases in `bigmath/test_assets`