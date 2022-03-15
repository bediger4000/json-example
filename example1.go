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

var exampleJSON = []byte(`[{"dom":"http://test.com","ip":"1.1.1.1"}, {"dom":"http://sub.domain.com", "ip":"1.2.3.4"}]`)

func main() {

	var data []AddressData

	err := json.Unmarshal(exampleJSON, &data)
	if err != nil {
		log.Fatal(err) // only call to log.Fatal, so we know where it is
	}

	for _, d := range data {
		fmt.Printf("Domain name: %q, IP address: %q\n", d.FQDN, d.IpAddr)
	}
}
