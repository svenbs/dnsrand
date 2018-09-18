package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"strings"
)

func main() {
	var service string
	flag.StringVar(&service, "service", "", "servicename to resolve")
	flag.Parse()

	if service == "" {
		fmt.Printf("usage:\n\n")
		flag.PrintDefaults()
		return
	}

	_, addrs, err := net.LookupSRV("", "", service)
	if err != nil {
		log.Fatalf("could not resolve service %q: %v", service, err)
	}

	if len(addrs) <= 0 {
		log.Fatalf("no addresses found for %q", service)
	}

	r, err := rand.Int(rand.Reader, big.NewInt(int64(len(addrs))))
	if err != nil {
		log.Fatalf("could not generate randomness: %v", err)
	}

	// Print random srv name and trim trailing dot from output
	fmt.Print(strings.TrimSuffix(addrs[r.Int64()].Target, "."))
}
