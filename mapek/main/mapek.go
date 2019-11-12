package main

import (
	"RAMid/mapek/planner"
	"fmt"
)

func main() {

	mapekPlanner := planner.MAPEKPlanner{IdPlugin: "requestor", NovaVersao: "v1"}
	go mapekPlanner.Planner()
	fmt.Scanln()
}
