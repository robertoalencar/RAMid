package analyser

import (
	"RAMid/mapek/planner"
	"RAMid/plugins"
	"RAMid/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MAPEKAnalyser struct{}

var versaoAtualRequestor int64
var versaoAtualMarshaller int64
var versaoAtualCrh int64
var versaoAtualSrh int64

func (MAPEKAnalyser) Analyse() {

	registraVersoesAtuaisComponentes()

	//Analisa adaptações do Requestor
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"requestor", analisaVersoesRequestor)

	//Analisa adaptações do Marshaller
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"marshaller", analisaVersoesMarshaller)

	//Analisa adaptações do Crh
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"crh", analisaVersoesCrh)

	//Analisa adaptações do Srh
	filepath.Walk(util.URL_REPOSITORIO_COMPONENTES+"srh", analisaVersoesSrh)
}

func registraVersoesAtuaisComponentes() {

	manager := plugins.Manager{}

	versaoAtualRequestor = manager.ObterVersaoAtualComponente(util.ID_COMPONENTE_REQUESTOR)
	versaoAtualMarshaller = manager.ObterVersaoAtualComponente(util.ID_COMPONENTE_MARSHALLER)
	versaoAtualCrh = manager.ObterVersaoAtualComponente(util.ID_COMPONENTE_CRH)
	versaoAtualSrh = manager.ObterVersaoAtualComponente(util.ID_COMPONENTE_SRH)
}

func analisaVersoesRequestor(path string, f os.FileInfo, err error) error {

	arquivo := strings.ToLower(f.Name())

	if strings.HasPrefix(arquivo, "v") {

		versao := strings.TrimPrefix(arquivo, "v")
		numVersao, err := strconv.ParseInt(versao, 10, 64)
		util.ChecaErro(err, "Falha ao tentar obter o número da versão do diretório: "+f.Name())

		if numVersao > versaoAtualRequestor {

			mapekPlanner := planner.MAPEKPlanner{IdPlugin: "requestor", NovaVersao: arquivo}
			go mapekPlanner.Planner()
		}
	}

	return nil
}

func analisaVersoesMarshaller(path string, f os.FileInfo, err error) error {

	arquivo := strings.ToLower(f.Name())

	if strings.HasPrefix(arquivo, "v") {

		versao := strings.TrimPrefix(arquivo, "v")
		numVersao, err := strconv.ParseInt(versao, 10, 64)
		util.ChecaErro(err, "Falha ao tentar obter o número da versão do diretório: "+f.Name())

		if numVersao > versaoAtualMarshaller {

			mapekPlanner := planner.MAPEKPlanner{IdPlugin: "marshaller", NovaVersao: arquivo}
			go mapekPlanner.Planner()
		}
	}

	return nil
}

func analisaVersoesCrh(path string, f os.FileInfo, err error) error {

	arquivo := strings.ToLower(f.Name())

	if strings.HasPrefix(arquivo, "v") {

		versao := strings.TrimPrefix(arquivo, "v")
		numVersao, err := strconv.ParseInt(versao, 10, 64)
		util.ChecaErro(err, "Falha ao tentar obter o número da versão do diretório: "+f.Name())

		if numVersao > versaoAtualCrh {

			mapekPlanner := planner.MAPEKPlanner{IdPlugin: "crh", NovaVersao: arquivo}
			go mapekPlanner.Planner()
		}
	}

	return nil
}

func analisaVersoesSrh(path string, f os.FileInfo, err error) error {

	arquivo := strings.ToLower(f.Name())

	if strings.HasPrefix(arquivo, "v") {

		versao := strings.TrimPrefix(arquivo, "v")
		numVersao, err := strconv.ParseInt(versao, 10, 64)
		util.ChecaErro(err, "Falha ao tentar obter o número da versão do diretório: "+f.Name())

		if numVersao > versaoAtualSrh {

			mapekPlanner := planner.MAPEKPlanner{IdPlugin: "srh", NovaVersao: arquivo}
			go mapekPlanner.Planner()
		}
	}

	return nil
}
