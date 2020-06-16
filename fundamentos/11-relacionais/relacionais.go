package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==", "golang" == "golang")
	fmt.Println("!=", 3 != 2.9999)
	fmt.Println("<", 3 < 2.9999)
	fmt.Println(">", 3 > 2.9999)
	fmt.Println("<=", 3 <= 2.9999)
	fmt.Println("=>", 3 >= 2.9999)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Jelso"}
	p2 := Pessoa{"Jelso"}

	fmt.Println("Pessoas", p1 == p2)
}
