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

func (network *Network) SendFindContactMessage(target *Contact) (ContactCandidates, bool) {
	fmt.Println("receiver address " + network.Receiver.Address)
	connection := false
	var reply ContactCandidates
	client, err := rpc.DialHTTP("tcp", network.Receiver.Address)
	if err != nil {
		connection = false
		fmt.Println("dialing:", err)
	} else {
		err = client.Call("Node.RPCFindNode", target, &reply)
		reply.CalDistance(*target)
		if err != nil {
			fmt.Println("RCP Find Node:", err)
			connection = false
		} else {
			fmt.Printf("Call successful, reply length: %d\n", reply.Len())
			connection = true
		}
		client.Close()
	}

	return reply, connection
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

func (network *Network) SendStoreMessage(key string, data []byte) error {

	client, err := rpc.DialHTTP("tcp", network.Receiver.Address)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply FindValueReply
	kv := KV{key, data}
	err = client.Call("Node.RPCStoreKV", kv, &reply)
	if err != nil {
		fmt.Println("RCP store error:", err)
	}

	return err

}
