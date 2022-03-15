package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type AddressData struct {
	FQDN   string `json:"dom"`
	IpAddr string `json:"ip"`
}

var exampleJSON = []byte(`[{"dom":"http://test.com","ip":"1.1.1.1"}, {"dom":"http://sub.domain.com", "ip":"1.2.3.4"}, {"dom":"http://sub.domain.com", "ip":"5.6.7.8"}]`)

func main() {

	var data []AddressData

	err := json.Unmarshal(exampleJSON, &data)
	if err != nil {
		log.Fatal(err) // only call to log.Fatal, so we know where it is
	}

	addresses := make(map[string][]string)

	for _, d := range data {
		addresses[d.FQDN] = append(addresses[d.FQDN], d.IpAddr)
	}

	for fqdn, addrs := range addresses {
		fmt.Printf("%q: %v\n", fqdn, addrs)
	}
}
