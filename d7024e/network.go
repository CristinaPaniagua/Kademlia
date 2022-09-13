package d7024e

import (
	"fmt"
	"net"
	"strconv"

	"github.com/go-ping/ping"
)

type Network struct {
	sender   *Contact
	receiver *Contact
}

func Listen(ip string, port int) {
	address := ip + ":" + strconv.Itoa(port)
	// Set up a listener
	ln, err := net.Listen("udp", address)

	// check if server was successfully created
	if err != nil {
		fmt.Println("The following error occured", err)
	} else {
		fmt.Println("The listener object has been created:", ln)
	}
}

func (network *Network) SendPingMessage(contact *Contact) {
	pinger, err := ping.NewPinger(contact.Address)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}

func (network *Network) SendFindContactMessage(contact *Contact) ContactCandidates {
	return ContactCandidates{}
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
