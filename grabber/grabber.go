package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
)

func scrapPage(ip, domain string) error {
	outFile := domain + ".html"
	// Open TCP connection to given address
	tcp, err := net.Dial("tcp", fmt.Sprintf("%s:80", ip))
	if err != nil {
		return err
	}
	defer tcp.Close()
	// Write HTTP Request headers
	// GET / HTTP/1.1
	// Host: example.com
	//
	//
	hdrs := fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\n\r\n", domain)
	fmt.Fprintf(tcp, hdrs)

	// Open file to store page
	outFileFullPath := path.Join(os.TempDir(), outFile)
	_ = os.Remove(outFileFullPath)
	f, err := os.OpenFile(outFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// A trick. We'll read RESPONSE from buffer until EOF symbol hit OR less than 256 bytes read. This is a DIRTY HACK.
	// Do not use in production.
	// Write to file by 256 bytes.
	tmp := make([]byte, 256)
	for {
		n, err := tcp.Read(tmp)
		if err == io.EOF {
			fmt.Println("got EOF, done reading from ")
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("got %d bytes, writing to %s\n", n, outFileFullPath)
		_, err = f.Write(tmp)
		if err != nil {
			return err
		}
		if n < 256 {
			break
		}
	}
	return nil
}

func main() {
	domainIP := "93.184.216.34"
	domainName := "example.com"
	if err := scrapPage(domainIP, domainName); err != nil {
		log.Fatal(err)
	}
}
