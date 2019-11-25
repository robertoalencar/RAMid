package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

var ln net.Listener
var conn net.Conn
var err error

var mapaConexoes map[string]net.Conn

func Receive(ch chan [3]interface{}) {

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)

	if mapaConexoes == nil {
		mapaConexoes = make(map[string]net.Conn)
	}

	conexao := serverHost + ":" + strconv.Itoa(serverPort)

	for key, value := range mapaConexoes {
		if key == conexao {
			conn = value
		}
	}

	if conn == nil {

		// create listener
		fmt.Println("Antes do net.Listen:", conexao)
		ln, err = net.Listen("tcp", conexao)
		fmt.Println("Depois do net.Listen:", conexao)
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}

		// accept connections
		fmt.Println("Antes do ln.Accept:", conexao)
		conn, err = ln.Accept()
		fmt.Println("Depois do ln.Accept:", conexao)
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}

		mapaConexoes[conexao] = conn
	}

	// receive message's size
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	dados[2] = msg
	ch <- dados
}

func Send(ch chan []byte) {

	msgToClient := <-ch

	// send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}
}