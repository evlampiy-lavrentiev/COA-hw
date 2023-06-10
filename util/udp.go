package util

import (
	"log"
	"net"
)

func MakeUDPConnector(IP string, port uint) *net.UDPConn {
	addr := net.UDPAddr{
		Port: int(port),
		IP:   net.ParseIP(IP),
	}
	UDPconn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	return UDPconn
}

func MakeDialConnector(IP string, port uint) *net.Conn {
	conn, err := net.Dial("udp", "native:1111")
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	return &conn
}

func MakeMulticastUDPConnector(IP string, port uint) *net.UDPConn {
	addr := net.UDPAddr{
		Port: int(port),
		IP:   net.ParseIP(IP),
	}
	UDPconn, err := net.ListenMulticastUDP("udp", nil, &addr)
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	return UDPconn
}
