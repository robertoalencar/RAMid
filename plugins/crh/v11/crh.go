package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

func SendReceive(ch chan [3]interface{}) {

	parametros := <-ch

	serverHost := parametros[0].(string)
	serverPort := parametros[1].(int)
	msgToServer := parametros[2].([]byte)

	// connect to server
	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", serverHost+":"+strconv.Itoa(serverPort))
		if err == nil {
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

	parametros[2] = msgFromServer

	ch <- parametros
}
