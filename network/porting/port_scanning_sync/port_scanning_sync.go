package main

import (
	"fmt"
	"net"
	"sync"
	"strconv"
)


var ipToScan = "192.168.65.52"
var minPort = 1
var maxPort = 1024

func main() {
	var wg sync.WaitGroup
	for port := minPort; port < maxPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf(ipToScan + ":" + strconv.Itoa(port))
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Host %s has open port: %d\n", ipToScan, port)
		}(port)
	}
	wg.Wait()
}