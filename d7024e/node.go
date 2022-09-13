package d7024e

import (
	"errors"
	"fmt"

	"github.com/go-ping/ping"
)

type Node struct {
	nodeID    *KademliaID
	IPAddress string
	port      string
	rt        *RoutingTable
	st        Store
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

type FindValueReply struct {
	Val      []byte
	Contacts []Contact
}

func (node *Node) FindValue(contact *Contact, key string) (FindValueReply, error) {

	reply := FindValueReply{}
	if contact == nil {
		return reply, errors.New("couldn't hash IP address")

	} else {
		val, ok := node.st.get(key)
		if ok {
			reply.Val = val
			return reply, nil
		} else {

			closestContacts, er := node.FindNode(contact)
			reply.Contacts = closestContacts.contacts
			return reply, er
		}

	}
}

func (node *Node) StoreKV(key string, value []byte) {
	node.st.add(key, value)
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
