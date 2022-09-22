package main

import (
	"fmt"
	"time"

	"github.com/CristinaPaniagua/Kademlia"
)

func main() {

	//NODE A
	rt := Kademlia.NewRoutingTable(Kademlia.NewContact(Kademlia.NewKademliaID("FFFFFFF100000000000000000000000000000000"), "127.0.0.1:8008"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("FFFFFFF200000000000000000000000000000000"), "127.0.0.1:8010"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111100000000000000000000000000000000"), "127.0.0.1:8002"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111200000000000000000000000000000000"), "127.0.0.1:8003"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004"))
	rt.AddContact(Kademlia.NewContact(Kademlia.NewKademliaID("2111111400000000000000000000000000000000"), "127.0.0.1t:8006"))
	nodeA := Kademlia.NewNode(Kademlia.NewKademliaID("FFFFFFF100000000000000000000000000000000"), "127.0.0.1", "8080", rt, Kademlia.NewStore())
	contactA := Kademlia.NodetoContact(&nodeA)
	contactA.String()

	//Running nodes

	fmt.Println("Running node " + nodeA.String())
	nodeA.RunRCP()
	time.Sleep(1 * time.Second)

	//TARGET
	target := Kademlia.NewContact(Kademlia.NewKademliaID("1111111A00000000000000000000000000000000"), "127.0.0.1:8005")
	/*
		contactB := Kademlia.NewContact(Kademlia.NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004")
		n := Kademlia.Network{&contactA, &contactB}
		contactResponse := n.SendFindContactMessage(&target)
		//Priting response
		fmt.Println("Response: ")
		for _, c := range contactResponse.Contacts {
			fmt.Println(c.StringDis())
		}

	*/
	contactResponse := nodeA.LookupContact(&target)
	fmt.Println(" DONEEEEEEEE ")
	fmt.Println("Response: ")
	for _, c := range contactResponse.Contacts {
		fmt.Println(c.StringDis())
	}

}
