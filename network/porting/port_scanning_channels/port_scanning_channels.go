package main

import (
	"fmt"
	"net"
	"sort"
)


var ipToScan = "192.168.65.52"
var minPort = 1
var maxPort = 1024

func worker(ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", ipToScan, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := minPort; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			ports <- i
		}
	}()

	for i := minPort; i < maxPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("Host %s has open port: %d\n", ipToScan, port)
	}
}