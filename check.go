package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination,
			conn.LocalAddr(), conn.RemoteAddr())
	}

	// Close the connection after checking
	if conn != nil {
		conn.Close()
	}

	return status
}
