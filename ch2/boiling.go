// Boiling exibe o ponto de ebulição da água.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point  = %gºF or %gºC\n", f, c)
	// Saída :
	// boiling point = 212ºF or 110ºC
}
