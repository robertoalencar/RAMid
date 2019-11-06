package plugins

import (
	"RAMid/util"
	"encoding/json"
	"io/ioutil"
)

type Manager struct{}

const URL_MANAGER_COMPONENTES = "/home/robertoalencar/go/src/RAMid/plugins/manager.json"
const URL_REPOSITORIO_COMPONENTES = "/home/robertoalencar/go/src/RAMid/plugins/"

func (Manager) ObterComponente(idComponent string) string {

	var versao string

	//Carrega o arquivo do JSON
	jsonComponentes, err := ioutil.ReadFile(URL_MANAGER_COMPONENTES)
	util.ChecaErro(err, "Falha ao carregar o descritor doscomponentes")

	//Cria o mapa que irá representar o JSON
	managerMap := make(map[string]string)

	//Converte o JSON no mapa
	json.Unmarshal([]byte(jsonComponentes), &managerMap)

	for key, value := range managerMap {
		if key == idComponent {
			versao = value
		}
	}

	return URL_REPOSITORIO_COMPONENTES + idComponent + "/" + versao + "/" + idComponent + ".so"
}
