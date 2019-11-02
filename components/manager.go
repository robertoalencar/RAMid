package components

import (
	"RAMid/util"
	"encoding/json"
	"io/ioutil"
)

type Manager struct{}

func (Manager) ObterVersaoComponente(identificador string) string {

	var componente string

	//Carrega o arquivo do JSON
	jsonComponentes, err := ioutil.ReadFile(util.URL_MANAGER_COMPONENTES)
	util.ChecaErro(err, "Falha ao carregar o descritor doscomponentes")

	//Cria o mapa que irá representar o JSON
	managerMap := make(map[string]string)

	//Converte o JSON no mapa
	json.Unmarshal([]byte(jsonComponentes), &managerMap)

	for key, value := range managerMap {
		if key == identificador {
			componente = value
		}
	}

	return componente
}
