package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func processMessage(m string) string {
	return fmt.Sprintf("I got a '%s' message", strings.Trim(m, "\r\n"))
}

func main() {
	port := 8081
	if tPort := os.Getenv("port"); tPort != "" {
		dPort, err := strconv.ParseInt(tPort, 10, 64)
		onErrFail(err)
		port = int(dPort)
	}
	fmt.Printf("Launching server on port: %d \n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	onErrFail(err)
	conn, err := ln.Accept()
	onErrFail(err)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", message)
		conn.Write([]byte(processMessage(message) + "\n"))
	}
}

func onErrFail(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
