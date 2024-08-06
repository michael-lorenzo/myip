package main

import (
	"fmt"
	"log"

	"github.com/pion/stun"
)

func printIP(res stun.Event) {
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	var xorAddr stun.XORMappedAddress
	if err := xorAddr.GetFrom(res.Message); err != nil {
		log.Fatal(err)
	}
	fmt.Println(xorAddr.IP)
}

func main() {
	u, err := stun.ParseURI("stun:stun.l.google.com:19302")
	if err != nil {
		log.Fatal(err)
	}
	c, err := stun.DialURI(u, &stun.DialConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	message, err := stun.Build(stun.TransactionID, stun.BindingRequest)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Do(message, printIP); err != nil {
		log.Fatal(err)
	}
}
