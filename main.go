package main

import "log"

func main() {
	log.Println(Soma(10, 10))
}

func Soma(a, b int) (soma int) {
	return a + b
}
