package util

import (
	"log"
)

const QTD_EXECUCOES_EXPERIMENTO = 5

func ChecaErro(err error, msg string) {

	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
