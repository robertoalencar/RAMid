package monitor

import (
	"RAMid/mapek/analyser"
	"RAMid/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MAPEKMonitor struct{}

var qtdComponentesAntesAdaptacao int
var qtdAtualComponentes int

func (MAPEKMonitor) Monitor(ehPrimeiraExecucao bool) {

	for {

		if ehPrimeiraExecucao {

			// Registra a quantidade inicial de componentes
			filepath.Walk(util.URL_REPOSITORIO_COMPONENTES, contadorInicial)
			ehPrimeiraExecucao = false

		} else {

			qtdAtualComponentes = 0

			//Conta a quantidade atual de componentes e atualiza a variável qtdAtualComponentes
			filepath.Walk(util.URL_REPOSITORIO_COMPONENTES, contadorComponentes)
		}

		if qtdComponentesAntesAdaptacao != qtdAtualComponentes {

			//A linha abaixo é para medir o tempo de adaptação
			fmt.Println("Início adaptação:", time.Now())

			//Invoca o Analyser
			mapekAnalyse := analyser.MAPEKAnalyser{}
			go mapekAnalyse.Analyse()
			qtdComponentesAntesAdaptacao = qtdAtualComponentes
		}

		time.Sleep(1 * time.Second)
	}
}

func contadorInicial(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdComponentesAntesAdaptacao = qtdComponentesAntesAdaptacao + 1
		qtdAtualComponentes = qtdAtualComponentes + 1
	}

	return nil
}

func contadorComponentes(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdAtualComponentes = qtdAtualComponentes + 1
	}

	return nil
}
