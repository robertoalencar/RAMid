package main

import (
	"RAMid/plugins"
	"RAMid/util"
	"plugin"
)

func Transmitir(n string) string {

	//Variável que indica o arquivo do próximo componente a ser executado
	idComponent := "componentB"

	//Carrega o arquivo do componente
	manager := plugins.Manager{}
	componente, err := plugin.Open(manager.ObterComponente(idComponent))
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
	//fmt.Println("Trace A", n)
	retorno := Transmitir(n)
	ch <- retorno
}
