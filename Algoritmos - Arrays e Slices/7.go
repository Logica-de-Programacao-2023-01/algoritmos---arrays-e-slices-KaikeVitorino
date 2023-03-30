package main

import "fmt"

func main() {
	slice := make([]int, 5)
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	Presente := false
	for _, valor := range slice {
		if valor == num {
			Presente = true
			break
		}
	}
	if !Presente {
		slice = append(slice, num)
	}
	fmt.Println(slice)
}
