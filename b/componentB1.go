package main

import (
	"RAMid/util"
	"fmt"
	"plugin"
)

func Transmitir(n int) {

	ch := make(chan int)

	componente, err := plugin.Open(util.REPOSITORIO_COMPONENTES + "componentC1.so")
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Executar")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Executar := funcao.(func(chan int))
	go Executar(ch)
	ch <- n
}

func Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace B", n)
	Transmitir(n)
}
