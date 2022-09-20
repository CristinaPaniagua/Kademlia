package Kademlia

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	alpha = 3
	b     = 160
	k     = 20
)

type Kademlia struct {
	node *Node
}

// find the
func (kademlia *Kademlia) LookupContact(target *Contact) ContactCandidates {
	me := kademlia.node
	contactMe := Contact{me.nodeID, me.IPAddress + ":" + me.port, nil}
	//List of the nodes already contacted
	contacted := []string{}
	contacted = append(contacted, me.nodeID.String())

	//channel
	contactChan := make(chan []Contact)

	//start point
	//List for clostest contacts- need to be updated in the loop
	contactList, err := me.FindNode(target) //first in my own bucket

	if err == nil {
		fmt.Printf("Found %d contacts", contactList.Len())
		//loop
		for { //until no new closest nodes

			contactList.Sort() //order the list in starting with the closest
			//updated list of contacts
			updatedlist := contactList.contacts

			//get alpha nodes to contact that haven been already
			nextAlpha := make([]Contact, 0, alpha)
			found := 0
			for i := 0; i < contactList.Len(); i++ {

				if !ExistsIn(&contactList.contacts[i], contacted) {
					nextAlpha = append(nextAlpha, contactList.contacts[i])
					found++
				}
				if found == alpha {
					break
				}

			}

			for i := 0; i < alpha; i++ {
				//send  requests to alpha nodes
				go func() {
					timeout := false
					n := Network{&contactMe, &nextAlpha[i]}
					responseList := n.SendFindContactMessage(target) //response of k closests nodes from the node i
					contacted = append(contacted, nextAlpha[i].ID.String())
					if timeout {
						contactList.RemoveContact(nextAlpha[i])
					} else {
						responseList.Sort() //order the list in starting with the closest

						Index := k
						if responseList.Len() < k {
							Index = responseList.Len()
						}
						contactChan <- responseList.contacts[:Index]

					}

				}()
			}

			for i := 0; i < alpha; i++ {
				//reding the responses from the channel
				r := <-contactChan
				//update list
				updatedlist = UpdateList(updatedlist, r)

			}

			//Compare if the list of contacts hasn't changed
			if !reflect.DeepEqual(contactList.contacts, updatedlist) {
				contactList.contacts = updatedlist
			} else {
				break
			}

		}
	}
	return contactList
}

// return an update list with the k closests nodes from two lists already order from closer to farder
func UpdateList(list1 []Contact, list2 []Contact) []Contact {
	updated := []Contact{}
	a := 0
	for i := 0; i < len(list1); i++ {
		for j := a; j < len(list2); j++ {
			if list2[j].Less(&list1[i]) {
				updated = append(updated, list2[j])
			} else {
				a = j
				updated = append(updated, list2[i])
				break
			}

		}

	}
	//return maximum k elements
	Index := k
	if len(updated) < k {
		Index = len(updated)
	}
	return updated[:Index]
}

// EvistsIn retunr true if the contact exists in the list
func ExistsIn(contact *Contact, Clist []string) bool {
	exists := false
	Id := contact.ID.String()
	for i := range Clist {
		if strings.Compare(Clist[i], Id) == 0 {
			exists = true
			break
		}
	}
	return exists
}

func (kademlia *Kademlia) LookupData(hash string) []byte {
	me := kademlia.node
	contactMe := Contact{me.nodeID, me.IPAddress + ":" + me.port, nil}

	//List of the nodes already contacted
	contacted := []string{}
	contacted = append(contacted, me.nodeID.String())

	//channels
	contactChan := make(chan []Contact)
	valueChan := make(chan []byte)

	//start point
	//List for clostest contacts- need to be updated in the loop
	reply, found, err := me.FindValue(hash) //first in my own bucket
	if err != nil {
		//TODO HANDLE ERRROR

	}
	if found {
		return reply.Val
	} else {
		contactList := ContactCandidates{reply.Contacts}

		for { //until value found

			contactList.Sort() //order the list in starting with the closest
			//updated list of contacts
			updatedlist := contactList.contacts

			//get alpha nodes to contact that haven been already
			nextAlpha := make([]Contact, 0, alpha)
			found := 0
			for i := 0; i < contactList.Len(); i++ {

				if !ExistsIn(&contactList.contacts[i], contacted) {
					nextAlpha = append(nextAlpha, contactList.contacts[i])
					found++
				}
				if found == alpha {
					break
				}

			}

			for i := 0; i < alpha; i++ {
				//send  requests to alpha nodes
				go func() {
					timeout := false
					n := Network{&contactMe, &nextAlpha[i]}
					response, found, _ := n.SendFindDataMessage(hash) //search for the value
					contacted = append(contacted, nextAlpha[i].ID.String())
					if found {
						valueChan <- response.Val
					} else {
						if timeout {
							contactList.RemoveContact(nextAlpha[i])
						} else {
							responseList := ContactCandidates{response.Contacts}
							responseList.Sort() //order the list in starting with the closest

							Index := k
							if responseList.Len() < k {
								Index = responseList.Len()
							}
							contactChan <- responseList.contacts[:Index]

						}
					}
				}()
			}

			for i := 0; i < alpha; i++ {
				//reding the responses from the channel
				var r []Contact
				select {
				case val := <-valueChan:
					return val
				case r = <-contactChan:
					//update list
					updatedlist = UpdateList(updatedlist, r)
				}

			}

			//Compare if the list of contacts hasn't changed
			if !reflect.DeepEqual(contactList.contacts, updatedlist) {
				contactList.contacts = updatedlist
			} else {
				return nil
			}

		}
	}

}

func (kademlia *Kademlia) Store(target *Contact, data []byte) {
	// me
	me := kademlia.node
	contactMe := Contact{me.nodeID, me.IPAddress + ":" + me.port, nil}
	// get k closest contacts to the key
	contacts := kademlia.LookupContact(target)
	// send STORE request
	for _, contact := range contacts.contacts {
		go func(contact Contact) {
			n := Network{&contactMe, &contact}
			n.SendStoreMessage(data)
		}(contact)
	}
}
