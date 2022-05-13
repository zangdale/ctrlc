package main

import (
	"fmt"
	"net"
	"testing"
)

func TestBackup(t *testing.T) {
	backup("/Users/.ctrlc/", "/Users/.ctrlc/ctrlc.db")

}

func TestNet(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
