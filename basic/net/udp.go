package net

import (
	"fmt"
	"net"
)

func readPacket(udpConn *net.UDPConn) {
	for {
		buf := make([]byte, 1024)
		n, srcAddr, err := udpConn.ReadFrom(buf[:])
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("receive %s from <%s>\n", buf[:n], srcAddr)
	}
}
