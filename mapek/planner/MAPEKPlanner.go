package planner

import "RAMid/mapek/executor"

type MAPEKPlanner struct {
	IdPlugin   string
	NovaVersao string
}

func (planner MAPEKPlanner) Planner() {

	mapekExecutor := executor.MAPEKExecutor{IdPlugin: planner.IdPlugin, NovaVersao: planner.NovaVersao}
	go mapekExecutor.Execute()
}
