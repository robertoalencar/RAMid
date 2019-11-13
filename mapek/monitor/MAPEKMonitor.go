package monitor

import (
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

		fmt.Println("qtdComponentesAntesAdaptacao: ", qtdComponentesAntesAdaptacao)
		fmt.Println("qtdAtualComponentes: ", qtdAtualComponentes)

		if qtdComponentesAntesAdaptacao != qtdAtualComponentes {
			//chama o Analyser
			fmt.Println("chama o Analyser")
		}

		time.Sleep(5 * time.Second)
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
