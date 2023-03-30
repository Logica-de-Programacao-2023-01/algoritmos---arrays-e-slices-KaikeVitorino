package main

import "fmt"

func main() {
	numero := [4]float64{10.1, 20.2, 30.3, 40.4}
	soma := 0.0
	for _, num := range numero {
		soma += num
	}
	fmt.Println("Resultado da soma: ", soma)
}
