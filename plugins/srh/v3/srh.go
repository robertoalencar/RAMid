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

	fmt.Println("chegou no receive")

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)

	fmt.Println("Host:", serverHost)
	fmt.Println("Port:", serverPort)

	if mapaConexoes == nil {
		fmt.Println("mapaConexoes == nil")
		mapaConexoes = make(map[string]net.Conn)
	}

	conexao := serverHost + ":" + strconv.Itoa(serverPort)

	for key, value := range mapaConexoes {
		if key == conexao {
			conn = value
		}
	}

	if conn == nil {

		fmt.Println("Cria uma nova conexao no SRH")

		// create listener
		ln, err = net.Listen("tcp", serverHost+":"+strconv.Itoa(serverPort))
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}

		fmt.Println("Listen", ln)

		// accept connections
		conn, err = ln.Accept()
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}

		fmt.Println("Conn", conn)

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
