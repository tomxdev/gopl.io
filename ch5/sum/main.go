package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}
func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice
}
