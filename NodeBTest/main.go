package main

import (
	"fmt"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	//NODE B
	rt := Kademlia.NewRoutingTable(Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1:8014"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("2111511400000000000000000000000000000000"), "127.0.0.1:8016"))
	node := Kademlia.NewNode(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1", "8004", rt, Kademlia.NewStore())
	contact := Kademlia.NodetoContact(&node)
	contact.String()

	//Running nodes

	fmt.Println("Running node " + node.String())
	node.RunRCP()
	for {
	}
}
