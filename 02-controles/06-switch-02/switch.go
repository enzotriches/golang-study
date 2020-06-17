package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Bom dia")
	case time.Hour() < 18:
		fmt.Println("Bom tarde")
	case time.Hour() < 24:
		fmt.Println("Bom noite")
	}
}
