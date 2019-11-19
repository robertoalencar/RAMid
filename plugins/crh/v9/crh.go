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

	fmt.Println("Chegou no SendReceive do CRH v9")

	parametros := <-ch

	serverHost := parametros[0].(string)
	serverPort := parametros[1].(int)
	msgToServer := parametros[2].([]byte)

	// connect to server
	var conn net.Conn
	var err error

	if mapaConexoes == nil {
		fmt.Println("Cria o mapa de conexoes")
		mapaConexoes = make(map[string]net.Conn)
	}

	conexao := serverHost + ":" + strconv.Itoa(serverPort)

	for key, value := range mapaConexoes {
		if key == conexao {
			conn = value
		}
	}

	if conn == nil {
		fmt.Println("Cria uma conexao para", conexao)
		conn, err = net.Dial("tcp", conexao)
		mapaConexoes[conexao] = conn
	}

	fmt.Println("mapaConexoes", mapaConexoes)

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
