package main

import (
	"net"
	"runtime"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [4]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		if n > 0 {
			_, err := conn.Write([]byte("Pong"))
			if err != nil {
				return
			}
		}
	}
}
func main() {
	runtime.GOMAXPROCS(2)

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ":8080")
	ln, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, _ := ln.Accept()
		go handleClient(conn)
	}
}
