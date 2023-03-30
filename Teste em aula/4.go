package main

import "fmt"

func main() {
	numeros := [6]float64{2, 4, 8, 16, 32, 64}
	var soma float64 = 0.0
	for _, num := range numeros {
		soma += num
	}
	media := soma / float64(len(numeros))
	fmt.Printf("Media vai ser: %.2f\n", media)

}
