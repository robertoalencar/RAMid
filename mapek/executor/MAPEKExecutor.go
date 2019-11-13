package executor

import (
	"RAMid/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MAPEKExecutor struct {
	IdPlugin   string
	NovaVersao string
}

func (executor MAPEKExecutor) Execute() {

	//Carrega o arquivo do JSON
	jsonComponentes, err := ioutil.ReadFile(util.URL_MANAGER_COMPONENTES)
	util.ChecaErro(err, "Falha ao carregar o descritor doscomponentes")

	//Cria o mapa que irá representar o JSON
	managerMapAnterior := make(map[string]string)
	managerMapNovo := make(map[string]string)

	//Converte o JSON no mapa
	json.Unmarshal([]byte(jsonComponentes), &managerMapAnterior)

	for key, value := range managerMapAnterior {

		versao := value

		if key == executor.IdPlugin {
			versao = executor.NovaVersao
		}

		managerMapNovo[key] = versao
	}

	//Converte o mapa atualizado no JSON
	arquivoAtualizado, err := json.Marshal(managerMapNovo)
	util.ChecaErro(err, "Falha ao converter o mapa no novo JSON")

	ioutil.WriteFile(util.URL_MANAGER_COMPONENTES, arquivoAtualizado, 0644)

	fmt.Println("Adaptação realizada com sucesso, componente", executor.IdPlugin, "atualizado para a versão", executor.NovaVersao)
}
