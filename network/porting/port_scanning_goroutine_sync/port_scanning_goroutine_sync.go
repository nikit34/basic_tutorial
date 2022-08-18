package main

import (
	"fmt"
	"sync"
)


var minPort = 1
var maxPort = 1024

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := minPort; i <= maxPort; i++ {
		wg.Add(1)
		ports <- i
	}
	
	wg.Wait()
	close(ports)
}