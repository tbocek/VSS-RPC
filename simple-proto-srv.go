package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/tbocek/VSS-RPC/go-gen2"
	"net"
)

func main() {
	fmt.Println("Listening...")

	tcpConn, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	conn, err := tcpConn.Accept() //do this in a go routine

	if err != nil {
		panic(err)
	}

	binary := make([]byte, 100)
	n, _ := conn.Read(binary)

	m := &pb.AMessage{}
	if err := proto.Unmarshal(binary[:n], m); err != nil {
		panic(err)
	}

	fmt.Printf("read: %v\n", m)
}
