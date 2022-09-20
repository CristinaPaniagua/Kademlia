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
	nodeID    *KademliaID
	IPAddress string
	port      string
	rt        *RoutingTable
	st        Store
}

func NewNode(id *KademliaID, address string, port string, rt *RoutingTable, st Store) Node {
	return Node{id, address, port, rt, st}
}

func (node *Node) RunRCP() {

	address := node.IPAddress + ":" + node.port
	// Set up a listener

	rpc.Register(node)
	rpc.HandleHTTP()
	ln, err := net.Listen("tcp", address)

	// check if server was successfully created
	if err != nil {
		fmt.Println("The following error occured", err)
	} else {
		fmt.Println("The node is running:", ln)
	}

	go http.Serve(ln, nil)

}

func (node *Node) FindNode(contact *Contact) (ContactCandidates, error) {
	closeCandidates := ContactCandidates{}
	closeContacts := []Contact{}
	if contact == nil {
		return closeCandidates, errors.New("couldn't hash IP address")
	} else {
		closeContacts = node.rt.FindClosestContacts(contact.ID, k)

	}

	for i := range closeContacts {
		fmt.Println(closeContacts[i].String())
	}
	closeCandidates.contacts = closeContacts
	return closeCandidates, nil
}

func (node *Node) RPCFindNode(contact *Contact, reply *ContactCandidates) error {

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
	val, ok := node.st.get(key)
	found := false
	if ok {
		reply.Val = val
		found = true
		reply.found = found
		return reply, found, nil

	} else {
		contact := NewContact(NewKademliaID(key), "")
		closestContacts, er := node.FindNode(&contact)
		reply.Contacts = closestContacts.contacts
		return reply, found, er
	}

}

func (node *Node) RPCFindValue(key string, reply *FindValueReply) error {

	valueReply, _, err := node.FindValue(key)
	*reply = valueReply
	return err
}

func (node *Node) StoreKV(key string, value []byte) {
	node.st.add(key, value)
}

func (node *Node) RPCStoreKV(key string, value []byte, reply bool) error {
	node.st.add(key, value)
	reply = true
	return nil
}

func (node *Node) Ping(target *Node) {
	pinger, err := ping.NewPinger(target.IPAddress + ":" + target.port)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}
