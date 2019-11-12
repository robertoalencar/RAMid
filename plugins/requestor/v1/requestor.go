package main

import (
	"RAMid/aux"
	"RAMid/distribution/marshaller"
	"RAMid/distribution/miop"
	"RAMid/plugins"
	"RAMid/util"
	"plugin"
)

func Invoke(chParam chan interface{}) {

	param := <-chParam
	inv := param.(aux.Invocation)

	marshallerInst := marshaller.Marshaller{}

	// create request packet
	reqHeader := miop.RequestHeader{Context: "Context", RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	// serialise request packet
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)

	manager := plugins.Manager{}
	componente, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_CRH))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("SendReceive")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	SendReceive := funcao.(func(chan [3]interface{}))

	ch := make(chan [3]interface{})
	go SendReceive(ch)

	var parametros [3]interface{}
	parametros[0] = inv.Host
	parametros[1] = inv.Port
	parametros[2] = msgToClientBytes

	// send request packet and receive reply packet
	ch <- parametros
	retorno := <-ch

	msgFromServerBytes := retorno[2].([]byte)
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	// extract result from reply packet
	r := miopPacketReply.Bd.RepBody.OperationResult

	chParam <- r
}
