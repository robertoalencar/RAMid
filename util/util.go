package util

import (
	"log"
)

const URL_MANAGER_COMPONENTES = "/home/robertoalencar/go/src/RAMid/components/manager.json"
const URL_REPOSITORIO_COMPONENTES = "/home/robertoalencar/go/src/RAMid/plugins/"
const QTD_EXECUCOES_EXPERIMENTO = 5

func ChecaErro(err error, msg string) {

	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
