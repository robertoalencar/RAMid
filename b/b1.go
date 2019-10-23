package b

import (
	"RAMid/c"
	"fmt"
)

type B1 struct{}

func Transmitir(n int) {

	ch := make(chan int)
	componente := c.C1{}
	go componente.Executar(ch)
	ch <- n
}

func (B1) Executar(ch chan int) {

	n := <-ch

	fmt.Println("Trace B", n)
	Transmitir(n)
}
