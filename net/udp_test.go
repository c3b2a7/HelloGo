package net

import (
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	addr1 := &net.UDPAddr{IP: net.IPv4zero, Port: 9000}
	addr2 := &net.UDPAddr{IP: net.IPv4zero, Port: 9001}

	udpConn1, _ := net.ListenUDP("udp", addr1)
	udpConn2, _ := net.ListenUDP("udp", addr2)

	go readPacket(udpConn1)
	go readPacket(udpConn2)

	time.Sleep(1 * time.Second)

	var err error
	addr1 = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000}
	addr2 = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9001}
	_, err = udpConn1.WriteTo([]byte("ping to "+addr2.String()), addr2)
	if err != nil {
		log.Printf("err: %s", err)
	}
	_, err = udpConn2.WriteTo([]byte("ping to "+addr1.String()), addr1)
	if err != nil {
		log.Printf("err: %s", err)
	}

	b := make([]byte, 1)
	os.Stdin.Read(b)
}
