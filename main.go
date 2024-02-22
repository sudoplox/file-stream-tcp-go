package main

import (
	"github.com/sudoplox/file-stream-tcp-go/nostream"
	"github.com/sudoplox/file-stream-tcp-go/stream"
	"net"
	"os"
	"time"
)

type FileServer interface {
	Start()
	ReadLoop(conn net.Conn)
	SendFile(size int) error
}

var FuncMap = map[string]FileServer{
	"stream":   &stream.StreamFileServer{},
	"nostream": &nostream.NoStreamFileServer{},
}

func main() {
	// read arguments from command line if -stream then call stream func if -no-stream then call noStream func
	// if no arguments are passed then call noStream func
	// if wrong arguments are passed then print usage
	server := FuncMap[os.Args[1]]
	go func() {
		time.Sleep(4 * time.Second)
		server.SendFile(16000)
	}()

	server.Start()
}
