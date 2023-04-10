package main

import "fmt"

func main() {
	a := [10]float64{1.0, 6.0, 2.0, 8.0, 4.0, 9.0, 3.0, 7.0, 5.0, 10.0}

	var s []float64
	for _, v := range a {
		if v > 5 {
			s = append(s, v)
		}
	}

	fmt.Println("Novo Slice:", s)
}
