package main

import (
	servidor "RAMid/server"
	"fmt"
)

func Transmitir(n int) {

	ch := make(chan int)
	componente := servidor.Servidor{}
	go componente.Executar(ch)
	ch <- n
}

func Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace D", n)
	Transmitir(n)
}
