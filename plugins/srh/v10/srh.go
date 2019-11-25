package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

var ln net.Listener
var conn net.Conn
var err error

func Receive(ch chan [3]interface{}) {

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)

	conexao := serverHost + ":" + strconv.Itoa(serverPort)

	// create listener
	for {
		ln, err = net.Listen("tcp", conexao)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("SRH 1 :: %s", err)
	}

	// accept connections
	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH 2 :: %s", err)
	}

	// receive message's size
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH 3 :: %s", err)
	}

	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH 4 :: %s", err)
	}

	dados[2] = msg
	ch <- dados
}

func Send(ch chan [2]interface{}) {

	dados := <-ch

	msgToClient := dados[0].([]byte)

	// send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("SRH 5 :: %s", err)
	}

	// send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("SRH 6 :: %s", err)
	}

	// close connection
	if conn != nil {
		conn.Close()
	}
	if ln != nil {
		ln.Close()
	}

	dados[1] = true
	ch <- dados
}
