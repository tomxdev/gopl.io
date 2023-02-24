package main

import (
	"fmt"
)

// squares devolve uma função que devolve
// o próximo quadrado perfeito sempre que ela é chamada

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

}
