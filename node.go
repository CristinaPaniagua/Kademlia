package Kademlia

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/go-ping/ping"
)

type Node struct {
	NodeID    *KademliaID
	IPAddress string
	Port      string
	Rt        *RoutingTable
	St        *Store
}

func NewNode(id *KademliaID, address string, port string, rt *RoutingTable, st *Store) Node {
	return Node{id, address, port, rt, st}
}

func NodetoContact(node *Node) Contact {
	address := node.IPAddress + ":" + node.Port
	//fmt.Println(address)
	contact := NewContact(node.NodeID, address)
	return contact
}

func (node *Node) JoinNetwork(IPadress string, port string, origin *Contact) *RoutingTable {
	node.IPAddress = IPadress
	node.Port = port
	meContact := NodetoContact(node)
	routingTable := NewRoutingTable(meContact)
	routingTable.AddContact(*origin)

	closestContacts := node.LookupContact(&meContact)

	for _, contact := range closestContacts.Contacts {
		routingTable.AddContact(contact)
	}

	return routingTable
}

func (node *Node) RunRCP() {

	address := node.IPAddress + ":" + node.Port
	fmt.Println("Node address: " + address)

	// Set up a listener
	rpc.Register(node)
	//to use multiple times
	oldMux := http.DefaultServeMux
	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	//---
	rpc.HandleHTTP()
	//---
	http.DefaultServeMux = oldMux

	ln, err := net.Listen("tcp", address)
	//fmt.Println(ln.Addr().String())
	// check if server was successfully created
	if err != nil {
		fmt.Println("The following error occured", err)
	} else {
		fmt.Println("The node is running:", ln)
	}

	go http.Serve(ln, mux)

}

func (node *Node) RPCPing(contact *Contact, reply *Contact) error {
	fmt.Printf("Ping from %s", contact.String())
	if contact == nil {
		return errors.New("couldn't hash IP address")
	}
	//add to the routing table
	node.Rt.AddContact(*contact)
	//TODO: UPDATE K-bucket
	//response
	contactMe := NodetoContact(node)
	fmt.Println(contactMe.String())
	*reply = contactMe
	return nil

}

func (node *Node) FindNode(contact *Contact) (ContactCandidates, error) {
	closeCandidates := ContactCandidates{}
	closeContacts := []Contact{}
	if contact == nil {
		return closeCandidates, errors.New("couldn't hash IP address")
	} else {
		closeContacts = node.Rt.FindClosestContacts(contact.ID, k)

	}
	fmt.Println("closest contacts from the node: " + node.String())
	for i := range closeContacts {
		fmt.Println(closeContacts[i].StringDis())
	}
	closeCandidates.Contacts = closeContacts
	return closeCandidates, nil
}

func (node *Node) RPCFindNode(contact *Contact, reply *ContactCandidates) error {
	fmt.Println(" searching contact:" + contact.String())
	closeCandidates, err := node.FindNode(contact)
	*reply = closeCandidates
	return err
}

type FindValueReply struct {
	Val      []byte
	Contacts []Contact
	found    bool
}

func (node *Node) FindValue(key string) (FindValueReply, bool, error) {
	reply := FindValueReply{}
	val, ok := node.St.Get(key)
	found := false
	if ok {
		reply.Val = val
		found = true
		reply.found = found
		return reply, found, nil

	} else {
		contact := NewContact(NewKademliaID(key), "")
		closestContacts, er := node.FindNode(&contact)
		reply.Contacts = closestContacts.Contacts
		return reply, found, er
	}

}

func (node *Node) RPCFindValue(key string, reply *FindValueReply) error {

	valueReply, _, err := node.FindValue(key)
	*reply = valueReply
	return err
}

func (node *Node) StoreKV(key string, value []byte) {
	node.St.Add(key, value)
}

func (node *Node) RPCStoreKV(key string, value []byte, reply bool) error {
	node.St.Add(key, value)
	reply = true
	return nil
}

func (node *Node) Ping(target *Node) {
	pinger, err := ping.NewPinger(target.IPAddress + ":" + target.Port)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}
func (node *Node) String() string {
	return fmt.Sprintf(`node("%s", "%s", "%s", routing table and store not printed)`, node.NodeID.String(), node.IPAddress, node.Port)
}
