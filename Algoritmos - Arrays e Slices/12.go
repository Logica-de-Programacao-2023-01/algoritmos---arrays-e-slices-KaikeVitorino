package main

import (
	"fmt"
)

func main() {
	array := [5]int{2, 3, 6, 8, 9}

	var slice []int
	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}

	fmt.Println(slice)
}
