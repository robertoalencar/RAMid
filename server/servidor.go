package servidor

type Servidor struct{}

func (Servidor) Executar(ch chan string) {

	n := <-ch
	ch <- "A" + n
}
