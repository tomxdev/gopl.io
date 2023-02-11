// Exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for key, arg := range os.Args[1:] {
		s += sep + arg
		concatenated := fmt.Sprintf("%d %s", key, arg)
		fmt.Println(concatenated)
		sep = " "
	}
	fmt.Println(s)
}
