package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func ping(times int, lockChan chan bool) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ":8080")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	for i := 0; i < int(times); i++ {
		conn.Write([]byte("Ping"))
		var buff [4]byte
		conn.Read(buff[0:])
	}
	lockChan <- true
	conn.Close()
}
func main() {
	runtime.GOMAXPROCS(2)

	totalPings := 1000000
	concurrentConnections := 100
	pingPerConnection := totalPings / concurrentConnections
	actualTotalPing := pingPerConnection * concurrentConnections

	lockChan := make(chan bool, concurrentConnections)

	start := time.Now()
	for i := 0; i < concurrentConnections; i++ {
		go ping(pingPerConnection, lockChan)
	}
	for i := 0; i < int(concurrentConnections); i++ {
		<-lockChan
	}

	elapsed := 1000000 * time.Since(start).Seconds()
	fmt.Println(elapsed / float64(actualTotalPing))
}
