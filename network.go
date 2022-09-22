package Kademlia

import (
	"fmt"
	"net/rpc"
)

type Network struct {
	Sender   *Contact
	Receiver *Contact
}

func (network *Network) Listen() {

	//TODO
}

func (network *Network) SendPingMessage(contact *Contact) bool {
	fmt.Println("Ping to -- " + contact.Address)
	client, err := rpc.DialHTTP("tcp", contact.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}
	var reply Contact
	err = client.Call("Node.RPCPing", contact, &reply)
	success := false
	if err != nil {
		fmt.Println("RCP Ping:", err)
	} else {
		fmt.Printf("Call successful, reply contact: %s\n", reply.String())
		success = true
	}

	return success

}

func (network *Network) SendFindContactMessage(target *Contact) ContactCandidates {
	fmt.Println("receiver address " + network.Receiver.Address)
	client, err := rpc.DialHTTP("tcp", network.Receiver.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply ContactCandidates

	err = client.Call("Node.RPCFindNode", target, &reply)
	reply.CalDistance(*target)
	if err != nil {
		fmt.Println("RCP Find Node:", err)
	} else {
		fmt.Printf("Call successful, reply length: %d\n", reply.Len())
	}
	client.Close()
	return reply
}

func (network *Network) SendFindDataMessage(hash string) (FindValueReply, bool, error) {
	client, err := rpc.DialHTTP("tcp", network.Receiver.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply FindValueReply
	err = client.Call("Node.RPCFindValue", hash, &reply)
	if err != nil {
		fmt.Println("RCP Find Node:", err)
	}

	return reply, reply.found, err
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
