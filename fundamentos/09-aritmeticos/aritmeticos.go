package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 4, 2
	fmt.Println("Soma =>", a+b)

	fmt.Println("Subt =>", a-b)

	fmt.Println("Divi =>", a/b)

	fmt.Println("Mult =>", a*b)

	fmt.Println("Modulo =>", a%b)

	//bitwise
	fmt.Println("E  =>", a&b)   // 11 & 10 = 10
	fmt.Println("OU  =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor  =>", a^b) // 11 ^ 10 = 1

	c, d := 3.2, 2.0

	fmt.Println("Maior=", math.Max(c, d))

	fmt.Println("Menor=", math.Min(c, d))

	fmt.Println("Potencia=", math.Pow(c, d))

}
