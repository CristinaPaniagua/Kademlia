package Kademlia

import (
	"fmt"
	"net/rpc"

	"github.com/go-ping/ping"
)

type Network struct {
	sender   *Contact
	receiver *Contact
}

func (network *Network) Listen() {

	//TODO
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
	client, err := rpc.DialHTTP("tcp", network.receiver.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply ContactCandidates
	err = client.Call("Node.RPCFindNode", contact, &reply)
	if err != nil {
		fmt.Println("RCP Find Node:", err)
	}

	return reply
}

func (network *Network) SendFindDataMessage(hash string) (FindValueReply, bool, error) {
	client, err := rpc.DialHTTP("tcp", network.receiver.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply FindValueReply
	err = client.Call("Node.RPCFindValue", hash, &reply)
	if err != nil {
		fmt.Println("RCP Find Node:", err)
	}

	return reply, reply.found, err
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
