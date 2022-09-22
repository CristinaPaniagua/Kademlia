package Kademlia

import (
	"fmt"
	"testing"
)

func TestRunRCP(t *testing.T) {

	ID := NewRandomKademliaID()
	node := NewNode(ID, "localhost", "8080", &RoutingTable{}, &Store{})
	node.RunRCP()

}

func TestNodeToContact(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	node := NewNode(NewKademliaID("FFFFFFF100000000000000000000000000000000"), "localhost", "8080", rt, NewStore())
	contact := NodetoContact(&node)
	fmt.Println(contact.String())
	fmt.Println(node.String())

	if contact.ID != node.NodeID {
		t.Error("IDs are not equal")
	}

}

func TestFindNode(t *testing.T) {

	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8003"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8004"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8005"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8006"))
	node := NewNode(NewKademliaID("FFFFFFF100000000000000000000000000000000"), "localhost", "8080", rt, NewStore())
	target := NewContact(NewKademliaID("1111111A00000000000000000000000000000000"), "localhost:8005")
	contactCandidates, err := node.FindNode(&target)
	if err != nil {
		t.Errorf("ERROR: %e", err)
	}
	fmt.Println("Response: ")
	for _, c := range contactCandidates.Contacts {
		fmt.Println(c.String())
	}

}

func TestRPCFindNode(t *testing.T) {

	//NODE A
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8003"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8004"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8005"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8006"))
	nodeA := NewNode(NewKademliaID("FFFFFFF100000000000000000000000000000000"), "localhost", "8080", rt, NewStore())
	contactA := NodetoContact(&nodeA)
	// NODE B
	rtb := NewRoutingTable(NewContact(NewKademliaID("FFFFFF2F00000000000000000000000000000000"), "localhost:8000"))
	rtb.AddContact(NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "localhost:8011"))
	rtb.AddContact(NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "localhost:8012"))
	rtb.AddContact(NewContact(NewKademliaID("1111131200000000000000000000000000000000"), "localhost:8013"))
	rtb.AddContact(NewContact(NewKademliaID("1111141300000000000000000000000000000000"), "localhost:8014"))
	rtb.AddContact(NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "localhost:8015"))
	rtb.AddContact(NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "localhost:8016"))
	nodeB := NewNode(NewKademliaID("FFFFFFF200000000000000000000000000000000"), "localhost", "8081", rt, NewStore())
	contactB := NodetoContact(&nodeB)

	//TARGET
	target := NewContact(NewKademliaID("1111111A00000000000000000000000000000000"), "localhost:8005")

	//Running server-node
	go func() {
		fmt.Println("Running node " + nodeA.String())
		nodeA.RunRCP()
	}()

	//Calling RPC
	n := Network{&contactB, &contactA}
	contactResponse := n.SendFindContactMessage(&target)
	//Priting response
	fmt.Println("Response: ")
	for _, c := range contactResponse.Contacts {
		fmt.Println(c.String())
	}

}
func TestRPCPing(t *testing.T) {

	//NODE A
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8003"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8004"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8005"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8006"))
	nodeA := NewNode(NewKademliaID("FFFFFFF100000000000000000000000000000000"), "127.0.0.1", "8080", rt, NewStore())
	contactA := NodetoContact(&nodeA)
	// NODE B
	rtb := NewRoutingTable(NewContact(NewKademliaID("FFFFFF2F00000000000000000000000000000000"), "localhost:8000"))
	rtb.AddContact(NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "localhost:8001"))
	rtb.AddContact(NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "localhost:8002"))
	rtb.AddContact(NewContact(NewKademliaID("1111131200000000000000000000000000000000"), "localhost:8003"))
	rtb.AddContact(NewContact(NewKademliaID("1111141300000000000000000000000000000000"), "localhost:8004"))
	rtb.AddContact(NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "localhost:8005"))
	rtb.AddContact(NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "localhost:8006"))
	nodeB := NewNode(NewKademliaID("FFFFFFF200000000000000000000000000000000"), "127.0.0.1", "8081", rt, NewStore())
	contactB := NodetoContact(&nodeB)

	//Running server-node
	go func() {
		fmt.Println("Running node " + nodeA.String())
		nodeA.RunRCP()
	}()

	//Calling RPC
	n := Network{&contactB, &contactA}
	n.SendPingMessage(&contactA)

}
func TestCreateNetwork15Nodes(t *testing.T) {

	//create origin node
	origin := Contact{}
	n1 := NewNode(NewKademliaID("0000000000000000000000000000000000000001"), "localhost", "8080", NewRoutingTable(origin), NewStore())
	origin = NodetoContact(&n1)

	//new node
	n2 := Node{}
	n2.NodeID = NewKademliaID("0000000000000000000000000000000000000002")
	n2.JoinNetwork("localhost", "8081", &origin)

	fmt.Println(n2)

	//new node
	n3 := Node{}
	n3.NodeID = NewKademliaID("0000000000000000000000000000000000000003")
	n3.JoinNetwork("localhost", "8082", &origin)

	fmt.Println(n3.NodeID.String() + " - Address" + n3.IPAddress + ":" + n3.Port)

}
