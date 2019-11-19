package invoker

import (
	"RAMid/distribution/clientproxy"
	"RAMid/distribution/miop"
	"RAMid/plugins"
	"RAMid/services/naming"
	"RAMid/util"
	"plugin"
)

type NamingInvoker struct{}

func (NamingInvoker) Invoke() {

	manager := plugins.Manager{}

	srhInst, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_SRH))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcReceive, err := srhInst.Lookup("Receive")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Receive := funcReceive.(func(chan [3]interface{}))

	funcSend, err := srhInst.Lookup("Send")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Send := funcSend.(func(chan []byte))

	namingImpl := naming.NamingService{}
	miopPacketReply := miop.Packet{}
	replyParams := make([]interface{}, 1)

	marshallerInst, err := plugin.Open(manager.ObterComponente(util.ID_COMPONENTE_MARSHALLER))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcUnmarshall, err := marshallerInst.Lookup("Unmarshall")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Unmarshall := funcUnmarshall.(func(chan interface{}))

	// control loop
	for {
		// receive request packet
		chReceiveSrh := make(chan [3]interface{})
		go Receive(chReceiveSrh)

		var parametros [3]interface{}
		parametros[0] = util.NAMING_HOST
		parametros[1] = util.NAMING_PORT

		// send request packet and receive reply packet
		chReceiveSrh <- parametros
		retornoReceive := <-chReceiveSrh

		rcvMsgBytes := retornoReceive[2].([]byte)

		// unmarshall request packet
		chUnmarshall := make(chan interface{})
		go Unmarshall(chUnmarshall)

		chUnmarshall <- rcvMsgBytes
		retornoUnmarshall := <-chUnmarshall
		miopPacketRequest := retornoUnmarshall.(miop.Packet)

		// extract operation name
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// demux request
		switch operation {
		case "Register":
			_p1 := miopPacketRequest.Bd.ReqBody.Body[0].(string)
			_map := miopPacketRequest.Bd.ReqBody.Body[1].(map[string]interface{})
			_proxyTemp := _map["Proxy"].(map[string]interface{})
			_p2 := clientproxy.ClientProxy{TypeName: _proxyTemp["TypeName"].(string), Host: _proxyTemp["Host"].(string), Port: int(_proxyTemp["Port"].(float64)), Id: int(_proxyTemp["Id"].(float64))}

			// dispatch request
			replyParams[0] = namingImpl.Register(_p1, _p2)
		case "Lookup":
			_p1 := miopPacketRequest.Bd.ReqBody.Body[0].(string)

			// dispatch request
			replyParams[0] = namingImpl.Lookup(_p1)
		}

		// assembly reply packet
		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replyParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// marshall reply packet
		funcMarshall, err := marshallerInst.Lookup("Marshall")
		util.ChecaErro(err, "Falha ao carregar a função do componente")

		Marshall := funcMarshall.(func(chan interface{}))

		chMarshaller := make(chan interface{})
		go Marshall(chMarshaller)

		// serialise request packet
		chMarshaller <- miopPacketReply
		retornoMarshall := <-chMarshaller
		msgToClientBytes := retornoMarshall.([]byte)

		// send reply packet
		chSendSrh := make(chan []byte)
		go Send(chSendSrh)
		chSendSrh <- msgToClientBytes
	}
}
