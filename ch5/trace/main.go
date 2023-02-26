package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // não se esqueça dos parênteses extras, ... muitas tarefas ...
	time.Sleep(10 * time.Second)      // simula uma operação demorada dormindo
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
