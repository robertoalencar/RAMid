package main

import (
	"RAMid/components"
	"RAMid/util"
	"fmt"
	"plugin"
)

func Transmitir(n int) {

	ch := make(chan int)

	versaoComponente := components.Manager{}.ObterVersaoComponente("componentB")
	componente, err := plugin.Open(util.URL_REPOSITORIO_COMPONENTES + versaoComponente)
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Executar")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Executar := funcao.(func(chan int))
	go Executar(ch)
	ch <- n
}

func Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace A", n)
	Transmitir(n)
}
