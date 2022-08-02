package main

import (
	"io"
	"log"
	"net"
)


var ipToScan = "192.168.65.52"

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", ipToScan + ":80")
	if err != nil {
		log.Fatalln("Unable to connect to " + ipToScan + " host")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}