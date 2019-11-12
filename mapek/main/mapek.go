package main

import (
	"RAMid/mapek"
	"fmt"
)

func main() {

	mapekExecutor := mapek.MAPEKExecutor{IdPlugin: "requestor", NovaVersao: "v5"}

	go mapekExecutor.Execute()

	fmt.Scanln()
}
