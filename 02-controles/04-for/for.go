package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for i := 0; i <= 5; i += 2 {
		fmt.Print(i)
	}
}
