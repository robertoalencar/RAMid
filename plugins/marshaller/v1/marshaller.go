package main

import (
	"RAMid/distribution/miop"
	"encoding/json"
	"log"
	"time"
)

func Marshall(ch chan interface{}) {

	param := <-ch
	msg := param.(miop.Packet)

	//testar a adaptação com GOB e com mensage pack

	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	time.Sleep(10 * time.Millisecond)

	ch <- r
}

func Unmarshall(ch chan interface{}) {

	param := <-ch
	msg := param.([]byte)

	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}

	ch <- r
}
