package main

import (
	"log"

	"github.com/kenzo0107/ipinfo"
)

func main() {
	ip := "172.217.26.3"

	ipInfo := ipinfo.Get(ip)
	log.Printf("%#v", ipInfo)
}
