package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

var mapaConexoes map[string]net.Conn

func SendReceive(ch chan [3]interface{}) {

	parametros := <-ch

	serverHost := parametros[0].(string)
	serverPort := parametros[1].(int)
	msgToServer := parametros[2].([]byte)

	var conn net.Conn
	var err error

	if mapaConexoes == nil {
		mapaConexoes = make(map[string]net.Conn)
	}

	conexao := serverHost + ":" + strconv.Itoa(serverPort)
	conexaoExistente := false

	for key, value := range mapaConexoes {
		if key == conexao {
			conn = value
			conexaoExistente = true
		}
	}

	if !conexaoExistente {

		fmt.Println("Antes do net.Dial:", conexao)
		// connect to server
		for {
			conn, err = net.Dial("tcp", conexao)
			if err == nil {
				break
			}
		}
		fmt.Println("Depois do net.Dial:", conexao)

		mapaConexoes[conexao] = conn
	}

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
