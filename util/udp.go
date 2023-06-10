package util

import (
	"log"
	"net"
	"strconv"
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

func MakeDialConnector(host string, port int) *net.Conn {
	conn, err := net.Dial("udp", host+":"+strconv.Itoa(port))
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
