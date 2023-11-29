package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("TINY TOXY PROXY")
	initProxy()
}

func initProxy() {
	listener, err := net.Listen("tcp", ":9797")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer listener.Close()
	//Accept a blocking call
	for {
		downstreamConn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error: ", err)
		}
		go func(downstreamConn net.Conn) {
			defer downstreamConn.Close()
			downstreamConn.Write([]byte("Hello from TINY TOXY\n"))
		}(downstreamConn)
	}

}
