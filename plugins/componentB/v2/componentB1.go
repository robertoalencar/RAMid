package main

import (
	"RAMid/components"
	"RAMid/util"
	"fmt"
	"plugin"
)

func Transmitir(n string) string {

	versaoComponente := components.Manager{}.ObterVersaoComponente("componentC")

	componente, err := plugin.Open(util.URL_REPOSITORIO_COMPONENTES + versaoComponente)
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Executar")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Executar := funcao.(func(chan string))

	ch := make(chan string)
	go Executar(ch)
	ch <- n

	return <-ch
}

func Executar(ch chan string) {

	n := <-ch
	fmt.Println("Trace B", n)
	retorno := Transmitir(n)
	ch <- retorno
}
