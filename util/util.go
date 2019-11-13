package util

import (
	"log"
)

const QTD_EXECUCOES_EXPERIMENTO = 5
const MIOP_REQUEST = 1
const NAMING_HOST = "localhost"
const NAMING_PORT = 1414

const URL_MANAGER_COMPONENTES = "/home/robertoalencar/go/src/RAMid/plugins/manager.json"
const URL_REPOSITORIO_COMPONENTES = "/home/robertoalencar/go/src/RAMid/plugins/"

const ID_COMPONENTE_REQUESTOR = "requestor"
const ID_COMPONENTE_CRH = "crh"
const ID_COMPONENTE_SRH = "srh"
const ID_COMPONENTE_MARSHALLER = "marshaller"

func ChecaErro(err error, msg string) {

	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
