package a

import (
	"RAMid/b"
	"fmt"
)

type A1 struct{}

func Transmitir(n int) {

	ch := make(chan int)
	componente := b.B1{}
	go componente.Executar(ch)
	ch <- n
}

func (A1) Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace A", n)
	Transmitir(n)
}
