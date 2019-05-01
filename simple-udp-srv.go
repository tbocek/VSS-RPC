package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	fmt.Println("listening...")
	inet := &net.UDPAddr{net.IPv4zero, 7000, ""}
	udpConn, err := net.ListenUDP("udp", inet)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 4)
	n, b2, err := udpConn.ReadFromUDP(b)

	if err != nil {
		panic(err)
	}
	fmt.Printf("connecting... read: %d, addr: %v, data: %v, decoded: %v\n",
		n, b2, b[:n], binary.BigEndian.Uint32(b))
	//binary.LittleEndian.Uint32(b))
}
