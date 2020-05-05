package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func whois(domain, server string) {
	log.Printf("Asking %s server about %s domain\n", server, domain)
	// Open TCP connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:43", server))
	onErrPanic(err)
	defer conn.Close()
	// > Write message to opened connection. REQUEST
	fmt.Fprintf(conn, "%s\r\n", domain)
	// < Read message from opened connection. RESPONSE
	d, err := ioutil.ReadAll(conn)
	onErrPanic(err)
	fmt.Printf("%s", string(d))
}

func main() {
	who := "example.com"
	whoisServer := "com.whois-servers.net"
	if len(os.Args) > 1 {
		who = os.Args[1]
	}
	if t := os.Getenv("WHOIS_SERVER"); t != "" {
		whoisServer = t
	}
	whois(who, whoisServer)
}

func onErrPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
