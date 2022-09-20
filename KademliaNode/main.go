package main

import (
	"fmt"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	fmt.Println("Node started")
	//take arguments
	//args := os.Args[1:]
	ID := Kademlia.NewRandomKademliaID()

	node := Kademlia.Node{ID, "localhost", "8080", &Kademlia.RoutingTable{}, Kademlia.Store{}}
	fmt.Println(node)
	node.runRCP()

}
