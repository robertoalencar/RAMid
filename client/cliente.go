package main

import (
	"RAMid/components"
	"RAMid/util"
	"fmt"
	"os"
	"plugin"
	"strconv"
	"time"
)

func Transmitir(n int) string {

	//Obtem a versão do componente do manager.json
	versaoComponente := components.Manager{}.ObterVersaoComponente("componentA")

	//Carrega o arquivo do componente
	componente, err := plugin.Open(util.URL_REPOSITORIO_COMPONENTES + versaoComponente)
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	//Indica qual a função que será executada de do componente
	funcao, err := componente.Lookup("Executar")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Executar := funcao.(func(chan string))

	ch := make(chan string)
	go Executar(ch)
	ch <- strconv.Itoa(n)

	return <-ch
}

func Executar(n int) string {

	return Transmitir(n)
}

func main() {

	nomeArquivo := time.Now().Format("2006-01-02 15:04:05") + " - Avaliação de Desempenho.txt"
	arquivo, _ := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	for i := 0; i < 10; i++ {

		t1 := time.Now()

		fmt.Println("Envio: ", i)
		valRetorno := Executar(i)
		fmt.Println("Retorno: ", valRetorno)

		t2 := time.Now()
		x := t2.Sub(t1)
		arquivo.WriteString(strconv.FormatInt(x.Microseconds(), 10) + "\n")

		time.Sleep(10 * time.Second)
	}

	arquivo.Close()
}
