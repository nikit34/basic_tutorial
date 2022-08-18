package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen = int32(320)
	promisc = true
	timeout = pcap.BlockForever
	// (ACK and FIN) or ACK or (ACK and PSH)
	filter = "tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[133] == 0x18"
	devFound = false
	results = make(map[string]int)
)

func capture(iface, target string) {
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(iface); err != nil {
		log.Panic(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Capturing packet")
	for packet := range source.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}

		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}

		srcHost := networkLayer.NetworkFlow().Src().String()
		srcPort := transportLayer.TransportFlow().Src().String()

		if srcHost != target {
			continue
		}

		results[srcPort] += 1
	}
}

func explode(portString string) ([]string, error) {
	ret := make([]string, 0)

	ports := strings.Split(portString, ",")
	for _, port := range ports {
		port := strings.TrimSpace(port)
		ret = append(ret, port)
	}
	return ret, nil
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("Usage: main.go <capture_iface> <target_ip> <port1,port2,port3>")
	}

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	iface := os.Args[1]
	for _, device := range devices {
		if device.Name == iface {
			devFound = true
			break
		}
	}

	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}

	ip := os.Args[2]
	go capture(iface, ip)
	time.Sleep(1 * time.Second)

	ports, err := explode(os.Args[3])
	if err != nil {
		log.Panicln(err)
	}

	for _, port := range ports {
		target := fmt.Sprintf("%s:%s", ip, string(port))
		fmt.Println("Trying: ", target)
		c, err := net.DialTimeout("tcp", target, 1000 * time.Millisecond)
		if err != nil {
			continue
		}
		c.Close()
	}
	time.Sleep(2 * time.Second)

	for port, confidence := range results {
		if confidence >= 1 {
			fmt.Printf("Port %s open (confidence: %d)\n", port, confidence)
		}
	}
}