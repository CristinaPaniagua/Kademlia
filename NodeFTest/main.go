package main

import (
	"fmt"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	//NODE F
	rt := Kademlia.NewRoutingTable(Kademlia.NewContact(Kademlia.NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "127.0.0.1:8002"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "127.0.0.1:8003"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004"))

	node := Kademlia.NewNode(Kademlia.NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1", "8015", rt, Kademlia.NewStore())
	contact := Kademlia.NodetoContact(&node)
	contact.String()

	//Running nodes

	fmt.Println("Running node " + node.String())
	node.RunRCP()
	for {
	}
}
