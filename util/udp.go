package util

import (
	"log"
	"net"
	"strconv"
)

func MakeDialConnector(host string, port int) *net.Conn {
	conn, err := net.Dial("udp", host+":"+strconv.Itoa(port))
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	return &conn
}

func MakeMulticastUDPConnector(mutlicastAddress string) *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", mutlicastAddress)
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	UDPconn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("Error creating UDP connector %v\n", err)
	}
	return UDPconn
}
