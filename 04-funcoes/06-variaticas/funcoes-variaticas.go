/**
* Vários parametros
 */
package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println(media(5, 5.5, 10, 3, 2, 4, 10, 10, 10, 10))
}
