package c

import (
	"RAMid/d"
	"fmt"
)

type C1 struct{}

func Transmitir(n int) {

	ch := make(chan int)
	componente := d.D1{}
	go componente.Executar(ch)
	ch <- n
}

func (C1) Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace C", n)
	Transmitir(n)
}
