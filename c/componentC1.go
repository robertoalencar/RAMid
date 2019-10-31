package main

import (
	"fmt"
	"plugin"
)

func Transmitir(n int) {

	ch := make(chan int)

	componente, err := plugin.Open("/home/robertoalencar/go/src/RAMid/components/componentD1.so")

	if err != nil {
		fmt.Println(err)
		return
	}

	funcao, err := componente.Lookup("Executar")
	if err != nil {
		fmt.Println(err)
		return
	}

	Executar := funcao.(func(chan int))
	go Executar(ch)
	ch <- n
}

func Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace C", n)
	Transmitir(n)
}
