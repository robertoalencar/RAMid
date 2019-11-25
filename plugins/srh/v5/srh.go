package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

var err error

func Receive(ch chan [4]interface{}) {

	dados := <-ch

	serverHost := dados[0].(string)
	serverPort := dados[1].(int)

	conexao := serverHost + ":" + strconv.Itoa(serverPort)

	var ln net.Listener
	var conn net.Conn

	// create listener
	for {
		ln, err = net.Listen("tcp", conexao)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	fmt.Println("Antes do Accept:", conexao)
	// accept connections
	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	fmt.Println("Depois do Accept:", conexao)

	// receive message's size
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	//fmt.Println("SRH :: Read :: 2")

	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	//fmt.Printf("SRH :: Read :: 3:%v\n", ch)

	dados[2] = msg
	dados[3] = conn
	ch <- dados
}

func Send(ch chan [2]interface{}) {

	dados := <-ch

	msgToClient := dados[0].([]byte)
	conn := dados[1].(net.Conn)

	// send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("CRH teste 1:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("CRH teste 2:: %s", err)
	}

	// close connection
	//conn.Close()
	//ln.Close()
}
