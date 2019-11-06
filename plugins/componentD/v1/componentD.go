package main

import (
	servidor "RAMid/server"
)

func Transmitir(n string) string {

	ch := make(chan string)
	componente := servidor.Servidor{}
	go componente.Executar(ch)
	ch <- n

	return <-ch
}

func Executar(ch chan string) {

	n := <-ch
	//fmt.Println("Trace D", n)
	retorno := Transmitir(n)
	ch <- retorno
}
