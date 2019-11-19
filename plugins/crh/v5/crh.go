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

	fmt.Println("Chegou no SendReceive do CRH")

	parametros := <-ch

	//serverHost := parametros[0].(string)
	serverHost := "localhost"
	//serverPort := parametros[1].(int)
	serverPort := 12345
	msgToServer := parametros[2].([]byte)

	// connect to server
	var conn net.Conn
	var err error

	if mapaConexoes == nil {
		fmt.Println("mapaConexoes == nil")
		mapaConexoes = make(map[string]net.Conn)
	}

	conexao := serverHost + ":" + strconv.Itoa(serverPort)
	fmt.Println(conexao)

	for key, value := range mapaConexoes {
		if key == conexao {
			conn = value
			fmt.Println("encontrou a conexao")
		}
	}

	if conn == nil {
		fmt.Println("cria uma nova conexao")
		conn, err = net.Dial("tcp", conexao)
		mapaConexoes[conexao] = conn
	}

	fmt.Println("mapaConexoes:")
	fmt.Println(mapaConexoes)

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
