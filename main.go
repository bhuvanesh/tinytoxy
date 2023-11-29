package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
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

		upstreamLink := &Link{
			ch:      make(chan []byte),
			latency: time.Second,
		}
		downstreamLink := &Link{
			ch:      make(chan []byte),
			latency: time.Second,
		}

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
			go io.Copy(upstreamLink, downstreamConn)
			go io.Copy(upstreamConn, upstreamLink)
			go io.Copy(downstreamLink, upstreamConn)
			io.Copy(downstreamConn, downstreamLink)
			// go io.Copy(upstreamConn, downstreamConn)
			// io.Copy(downstreamConn, upstreamConn)

		}(downstreamConn)

	}

}

type Link struct {
	ch      chan []byte
	latency time.Duration
}

// Read reads data from the connection.
// Read can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetReadDeadline.
func (l *Link) Read(b []byte) (int, error) {
	// b = <- l.ch
	data := <-l.ch
	time.Sleep(l.latency)
	copy(b, data)
	return len(b), nil
}

// Write writes data to the connection.
// Write can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetWriteDeadline.
func (l *Link) Write(b []byte) (int, error) {
	l.ch <- b
	return len(b), nil
}
