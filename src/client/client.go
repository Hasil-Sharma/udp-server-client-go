package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", os.Args[1])
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", ":0")
	CheckError(err)

	OutConn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	f, err := os.Create("timestamp.dat")
	CheckError(err)

	defer f.Close()
	defer OutConn.Close()
	buffer := make([]byte, 1024)

	f.WriteString("t1\tt2\tt3\tt4\n")

	for {
		clientReqTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		_, err := OutConn.Write([]byte(clientReqTime))
		CheckError(err)

		n, _, err := OutConn.ReadFromUDP(buffer)
		CheckError(err)

		clientRecvTime := strconv.FormatInt(time.Now().UnixNano(), 10)

		tokens := strings.Split(string(buffer[:n]), " ")
		serverRecvTime, serverResTime := tokens[0], tokens[1]

		f.WriteString(clientReqTime + "\t" +
			serverRecvTime + "\t" +
			serverResTime + "\t" +
			clientRecvTime + "\n")

		fmt.Println("C Req Time\t", clientReqTime,
			"\tS Recv Time\t", serverRecvTime,
			"\tS Res Time\t", serverResTime,
			"\tC Recv Time\t", clientRecvTime)

		time.Sleep(time.Second * 10)
	}
}
