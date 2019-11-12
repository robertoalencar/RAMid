package main

import (
	"RAMid/distribution/miop"
	"encoding/json"
	"log"
)

func Marshall(ch chan interface{}) {

	param := <-ch
	msg := param.(miop.Packet)

	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

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
