package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("TINY TOXY PROXY")
	initProxy()
}

func initProxy() {
	//mockServer to test our proxy upstream/downstream connection
	go mockServer()
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
			upstreamConn, err := net.Dial("tcp", ":9898")
			if err != nil {
				log.Fatal("Error: ", err)
			}
			defer upstreamConn.Close()
			// downstreamConn.Write([]byte("Hello from TINY TOXY\n"))
			go io.Copy(upstreamConn, downstreamConn)
			io.Copy(downstreamConn, upstreamConn)

		}(downstreamConn)

	}

}
