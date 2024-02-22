package nostream

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
)

type NoStreamFileServer struct {
}

func (nsfs *NoStreamFileServer) SendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	n, err := conn.Write(file)
	if err != nil {
		return err
	}
	fmt.Printf("sent %d bytes over the network\n", n)
	return nil
}

func (nsfs *NoStreamFileServer) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go nsfs.ReadLoop(conn)
	}
}

func (nsfs *NoStreamFileServer) ReadLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		file := buf[:n]
		fmt.Println(file)
		fmt.Printf("received %d bytes over the network\n", n)
	}
}
