package main

import (
	"fmt"
)

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num int
	fmt.Print("Informe um número para ser adicionado ao primeiro e ao último elemento do array: ")
	fmt.Scanln(&num)

	array[0] += num
	array[len(array)-1] += num

	fmt.Println(array)
}
