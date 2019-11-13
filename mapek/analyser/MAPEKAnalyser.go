package analyser

import (
	"RAMid/util"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MAPEKAnalyser struct{}

var qtdArquivosRequestor int
var qtdArquivosMarshaller int
var qtdArquivosCrh int
var qtdArquivosSrh int

func (MAPEKAnalyser) Analyse(ehPrimeiraExecucao bool) {

	for {

		if ehPrimeiraExecucao {

			registraQuantidadeComponentesRequestor()
			ehPrimeiraExecucao = false

		} else {

			// monitoraVersoesRequestor()
			// monitoraVersoesMarshaller()
			// monitoraVersoesCrh()
			// monitoraVersoesSrh()
		}

		time.Sleep(5 * time.Second)
	}
}

func registraQuantidadeComponentesRequestor() {

	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"requestor", contadorComponentesRequestor)
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"marshaller", contadorComponentesMarshaller)
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"crh", contadorComponentesCrh)
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"srh", contadorComponentesSrh)
}

func contadorComponentesRequestor(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdArquivosRequestor = qtdArquivosRequestor + 1
	}

	return nil
}

func contadorComponentesMarshaller(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdArquivosMarshaller = qtdArquivosMarshaller + 1
	}

	return nil
}

func contadorComponentesCrh(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdArquivosCrh = qtdArquivosCrh + 1
	}

	return nil
}

func contadorComponentesSrh(path string, arquivo os.FileInfo, err error) error {

	if strings.HasSuffix(arquivo.Name(), ".so") {
		qtdArquivosSrh = qtdArquivosSrh + 1
	}

	return nil
}
