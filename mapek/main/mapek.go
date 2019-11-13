package main

import (
	"RAMid/mapek/monitor"
	"fmt"
)

func main() {

	mapekMonitor := monitor.MAPEKMonitor{}
	go mapekMonitor.Monitor(true)

	fmt.Println("MAPE-K Running ...")
	fmt.Scanln()
}
