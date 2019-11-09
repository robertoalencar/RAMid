package plugins

import (
	"RAMid/util"
	"encoding/json"
	"io/ioutil"
)

type Manager struct{}

func (Manager) ObterComponente(idComponent string) string {

	urlManagerComponentes := "/home/robertoalencar/go/src/RAMid/plugins/manager.json"
	urlRepositoriComponentes := "/home/robertoalencar/go/src/RAMid/plugins/"

	var versao string

	//Carrega o arquivo do JSON
	jsonComponentes, err := ioutil.ReadFile(urlManagerComponentes)
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

	return urlRepositoriComponentes + idComponent + "/" + versao + "/" + idComponent + ".so"
}
