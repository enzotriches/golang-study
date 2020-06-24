package main

import "fmt"

func main() {
	arr := [3]int{4, 2, 1}
	var slice []int
	//append(slice, values...)
	slice = append(slice, 1, 2, 3)
	fmt.Println(arr, slice)

	slice2 := make([]int, 2)
	//copy(slice,slice) apenas slices por argumento
	copy(slice2, slice)

	fmt.Println(arr, slice, slice2)
}
