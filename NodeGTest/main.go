package main

import (
	"fmt"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	//NODE G
	rt := Kademlia.NewRoutingTable(Kademlia.NewContact(Kademlia.NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1:8014"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "127.0.0.1:8002"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "127.0.0.1:8003"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111141200000000000000000000000000000000"), "127.0.0.1:8033"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004"))

	node := Kademlia.NewNode(Kademlia.NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1", "8014", rt, Kademlia.NewStore())
	contact := Kademlia.NodetoContact(&node)
	contact.String()

	//Running nodes

	fmt.Println("Running node " + node.String())
	node.RunRCP()
	for {
	}
}
