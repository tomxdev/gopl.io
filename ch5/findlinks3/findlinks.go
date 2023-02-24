package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// breadthFirst chama f para cada item em worklist
// Todo item devolvido por f é adicionado à worklist
// f é chamada no máximo uma vez para cada item
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Faz crawling da web em largura (breadth-first),
	// começando com os argumentos da linha de comando
	breadthFirst(crawl, os.Args[1:])
}
