package crh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

func Transmitir(ch chan [3]interface{}) {

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)
	msgToServer := dados[2].([]byte)

	retorno := SendReceive(serverHost, serverPort, msgToServer)

	dados[2] = retorno

	ch <- dados
}

func SendReceive(serverHost string, serverPort int, msgToServer []byte) []byte {

	// connect to server
	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", serverHost+":"+strconv.Itoa(serverPort))
		if err == nil {
			//log.Fatalf("CRH:: %s", err)
			break
		}
	}

	defer conn.Close()

	// send message's size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToServer))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	conn.Write(sizeMsgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// receive message's size
	sizeMsgFromServer := make([]byte, 4)
	_, err = conn.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	// receive reply
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = conn.Read(msgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msgFromServer
}
