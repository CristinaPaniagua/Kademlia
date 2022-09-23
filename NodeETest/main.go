package main

import (
	"fmt"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	//NODE E
	rt := Kademlia.NewRoutingTable(Kademlia.NewContact(Kademlia.NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "127.0.0.1:8006"))

	node := Kademlia.NewNode(Kademlia.NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1", "8012", rt, Kademlia.NewStore())
	contact := Kademlia.NodetoContact(&node)
	contact.String()

	//Running nodes

	fmt.Println("Running node " + node.String())
	node.RunRCP()
	for {
	}
}
