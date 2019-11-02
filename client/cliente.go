package main

import (
	"RAMid/components"
	"RAMid/util"
	"fmt"
	"plugin"
)

func Transmitir(n int) {

	ch := make(chan int)

	//Carrega o arquivo do componente
	versaoComponente := components.Manager{}.ObterVersaoComponente("componentA")
	componente, err := plugin.Open(util.URL_REPOSITORIO_COMPONENTES + versaoComponente)
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	//Indica qual a função que será executada de do componente
	funcao, err := componente.Lookup("Executar")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Executar := funcao.(func(chan int))
	go Executar(ch)
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
