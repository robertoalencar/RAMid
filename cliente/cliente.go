package main

import (
	"RAMid/a"
	"fmt"
)

func Transmitir(n int) {

	ch := make(chan int)
	componente := a.A1{}
	go componente.Executar(ch)
	ch <- n
}

func Executar(n int) {

	fmt.Println("Enviando para o servidor", n)
	Transmitir(n)
}

func main() {

	go Executar(1)
	fmt.Scanln()
}
