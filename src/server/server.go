package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	portNumber := ":" + os.Args[1]
	ServerAddr, err := net.ResolveUDPAddr("udp", portNumber)
	CheckError(err)

	// Opening port for packets
	InConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	defer InConn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := InConn.ReadFromUDP(buffer)
		CheckError(err)

		reqRecvTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		fmt.Println("Received", string(buffer[0:n]), " from ", addr)
		CheckError(err)

		reqResTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		_, err = InConn.WriteToUDP([]byte(reqRecvTime + " " + reqResTime), addr)
		CheckError(err)
	}
}
