package main

import (
	"RAMid/aux"
	"RAMid/distribution/miop"
	"RAMid/plugins"
	"RAMid/util"
	"fmt"
	"plugin"
)

func Invoke(chParam chan interface{}) {

	fmt.Println("Nova versao do Requestor")

	param := <-chParam
	inv := param.(aux.Invocation)

	// create request packet
	reqHeader := miop.RequestHeader{Context: "Context", RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	manager := plugins.Manager{}

	marshallerInst, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_MARSHALLER))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcMarshall, err := marshallerInst.Lookup("Marshall")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Marshall := funcMarshall.(func(chan interface{}))

	chMarshaller := make(chan interface{})
	go Marshall(chMarshaller)

	// serialise request packet
	chMarshaller <- miopPacketRequest
	retornoMarshall := <-chMarshaller
	msgToClientBytes := retornoMarshall.([]byte)

	crhInst, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_CRH))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcSendReceive, err := crhInst.Lookup("SendReceive")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	SendReceive := funcSendReceive.(func(chan [3]interface{}))

	chRequestor := make(chan [3]interface{})
	go SendReceive(chRequestor)

	var parametros [3]interface{}
	parametros[0] = inv.Host
	parametros[1] = inv.Port
	parametros[2] = msgToClientBytes

	// send request packet and receive reply packet
	chRequestor <- parametros
	retornoSendReceive := <-chRequestor

	msgFromServerBytes := retornoSendReceive[2].([]byte)

	funcUnmarshall, err := marshallerInst.Lookup("Unmarshall")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Unmarshall := funcUnmarshall.(func(chan interface{}))

	chUnmarshall := make(chan interface{})
	go Unmarshall(chUnmarshall)

	chUnmarshall <- msgFromServerBytes
	retornoUnmarshall := <-chUnmarshall
	miopPacketReply := retornoUnmarshall.(miop.Packet)

	// extract result from reply packet
	r := miopPacketReply.Bd.RepBody.OperationResult

	chParam <- r
}
