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

func Receive(ch chan [3]interface{}) {

	fmt.Println("Chegou no Receive do SRH v4")

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)

	fmt.Println("ServerHost:", serverHost)
	fmt.Println("ServerPort:", serverPort)

	// create listener
	ln, err = net.Listen("tcp", serverHost+":"+strconv.Itoa(serverPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// accept connections
	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
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

	// close connection
	conn.Close()
	ln.Close()
}
