package main

import (
	"RAMid/mapek/monitor"
	"fmt"
)

func main() {

	// mapekPlanner := planner.MAPEKPlanner{IdPlugin: "requestor", NovaVersao: "v1"}
	// go mapekPlanner.Planner()
	// fmt.Scanln()

	mapekMonitor := monitor.MAPEKMonitor{}
	go mapekMonitor.Monitor(true)
	fmt.Scanln()
}
