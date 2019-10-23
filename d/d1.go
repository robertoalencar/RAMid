package d

import (
	"RAMid/servidor"
	"fmt"
)

type D1 struct{}

func Transmitir(n int) {

	ch := make(chan int)
	componente := servidor.Servidor{}
	go componente.Executar(ch)
	ch <- n
}

func (D1) Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace D", n)
	Transmitir(n)
}
