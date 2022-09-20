package Kademlia

import (
	"testing"
)

func TestNode(t *testing.T) {

	ID := NewRandomKademliaID()
	node := newNode(ID, "localhost", "8080", &RoutingTable{}, Store{})
	node.runRCP()
	
}
