package servidor

import (
	"fmt"
)

type Servidor struct{}

func (Servidor) Executar(ch chan int) {

	n := <-ch

	fmt.Println("Dado recebido pelo servidor:", n)
}
