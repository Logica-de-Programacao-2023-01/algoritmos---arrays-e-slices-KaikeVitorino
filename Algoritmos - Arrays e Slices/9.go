package main

import (
	"fmt"
)

func main() {
	a := [6]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	var num float64
	fmt.Print("Digite o número a ser adicionado em todas as posições: ")
	fmt.Scan(&num)

	for i := range a {
		a[i] += num
	}

	fmt.Println("Array resultante:", a)
}
