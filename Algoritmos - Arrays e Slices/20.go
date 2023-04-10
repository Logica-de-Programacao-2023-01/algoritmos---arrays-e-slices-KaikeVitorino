package main

import "fmt"

func main() {
	n := 5
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&array[i])
	}

	isSorted := true
	for i := 0; i < n-1; i++ {
		if array[i] > array[i+1] {
			isSorted = false
			break
		}
	}

	if isSorted {
		fmt.Println("O array está ordenado em ordem crescente")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente")
	}
}
