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

// find the
func (node *Node) LookupContact(target *Contact) ContactCandidates {

	contactMe := Contact{node.NodeID, node.IPAddress + ":" + node.Port, nil}
	//List of the nodes already contacted
	contacted := []string{}
	contacted = append(contacted, node.NodeID.String())

	//channel
	contactChan := make(chan ContactCandidates)
	removalChan := make(chan Contact)
	//start point
	//List for clostest contacts- need to be updated in the loop
	contactList, err := node.FindNode(target) //first in my own bucket

	if err == nil {
		fmt.Printf("Found %d contacts", contactList.Len())
		//loop
		for { //until no new closest nodes

			contactList.Sort() //order the list in starting with the closest
			//updated list of contacts
			var updatedCandidates ContactCandidates
			updatedCandidates.Contacts = contactList.Contacts
			//get alpha nodes to contact that haven been already
			nextAlpha := make([]Contact, 0, alpha)
			found := 0
			for f := 0; f < contactList.Len(); f++ {

				if !ExistsIn(&contactList.Contacts[f], contacted) {
					nextAlpha = append(nextAlpha, contactList.Contacts[f])
					found++
				}
				if found == alpha {
					break
				}

			}

			fmt.Printf("nextAlpha len= %d \n ", len(nextAlpha))
			for e := 0; e < len(nextAlpha); e++ {
				fmt.Println(nextAlpha[e].String() + ", distance:" + nextAlpha[e].distance.String())
			}

			for i := 0; i < len(nextAlpha); i++ {
				//send  requests to alpha nodes
				go func(ind int) {

					fmt.Println("Sending message to: " + nextAlpha[ind].String())
					n := Network{&contactMe, &nextAlpha[ind]}
					responseList, succesfulConnection := n.SendFindContactMessage(target) //response of k closests nodes from the node i

					contacted = append(contacted, nextAlpha[ind].ID.String())
					if !succesfulConnection {

						removalChan <- nextAlpha[ind]
					} else {
						fmt.Printf("Response list len= %d \n ", len(responseList.Contacts))

						responseList.Sort() //order the list in starting with the closest
						contactChan <- responseList
					}

				}(i)

			}

			for d := 0; d < len(nextAlpha); d++ {
				fmt.Println("waiting chanel communicaiton...")
				//reding the responses from the channel

				select {
				case toRemove := <-removalChan:
					contactList.RemoveContact(toRemove)
					updatedCandidates.RemoveContact(toRemove)
				case r := <-contactChan:
					if r.Len() > 0 {
						//update list
						fmt.Println("....Updating list")
						updatedCandidates = UpdateList(&updatedCandidates, &r)
						updatedCandidates.CalDistance(*target)
						updatedCandidates.Sort()

					}
				}

			}
			fmt.Printf("UpdatedList len= %d \n ", len(updatedCandidates.Contacts))
			for m := 0; m < len(updatedCandidates.Contacts); m++ {
				fmt.Println(updatedCandidates.Contacts[m].StringDis())
			}

			//Compare if the list of contacts hasn't changed
			if !reflect.DeepEqual(contactList.Contacts, updatedCandidates.Contacts) {
				fmt.Println("....List of contacts changed")
				contactList.Contacts = updatedCandidates.Contacts
			} else {
				fmt.Println("....List of contacts has not changed")
				// Check if all nodes have been contacted
				AllSeen := false
				for _, c := range contactList.Contacts {

					AllSeen = ExistsIn(&c, contacted)
					//fmt.Println(AllSeen)
				}

				fmt.Println("Final allseen:")
				fmt.Println(AllSeen)
				if AllSeen {
					//if all have been contacted- FINISH
					contactList.Contacts = updatedCandidates.Contacts
					break
				} // if not continue

			}

		}
	}
	return contactList
}

// return an update list with the k closests nodes from two lists already order from closer to farder
// WARNING- RESPONSE NOT IN ORDER!!!
func UpdateList(c1 *ContactCandidates, c2 *ContactCandidates) ContactCandidates {
	var updated ContactCandidates
	c1.Append(c2.Contacts)
	contacts := RemoveDupes(c1.Contacts)
	/*
		fmt.Println("after removing????")
		for m := 0; m < len(contacts); m++ {
			fmt.Println(contacts[m].StringDis())
		}
	*/
	//return maximum k elements
	Index := k
	if len(contacts) < k {
		Index = len(contacts)
	}
	fmt.Println(Index)
	updated.Contacts = contacts[:Index]

	return updated
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

/*
func (node *Node) LookupData(hash string) []byte {

	contactMe := Contact{node.NodeID, node.IPAddress + ":" + node.Port, nil}

	//List of the nodes already contacted
	contacted := []string{}
	contacted = append(contacted, node.NodeID.String())

	//channels
	contactChan := make(chan []Contact)
	valueChan := make(chan []byte)

	//start point
	//List for clostest contacts- need to be updated in the loop
	reply, found, err := node.FindValue(hash) //first in my own bucket
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
			updatedlist := contactList.Contacts

			//get alpha nodes to contact that haven been already
			nextAlpha := make([]Contact, 0, alpha)
			found := 0
			for i := 0; i < contactList.Len(); i++ {

				if !ExistsIn(&contactList.Contacts[i], contacted) {
					nextAlpha = append(nextAlpha, contactList.Contacts[i])
					found++
				}
				if found == alpha {
					break
				}

			}

			for i := 0; i < alpha; i++ {
				//send  requests to alpha nodes
				go func(ind int) {
					timeout := false
					n := Network{&contactMe, &nextAlpha[ind]}
					response, found, _ := n.SendFindDataMessage(hash) //search for the value
					contacted = append(contacted, nextAlpha[ind].ID.String())
					if found {
						valueChan <- response.Val
					} else {
						if timeout {
							contactList.RemoveContact(nextAlpha[ind])
						} else {
							responseList := ContactCandidates{response.Contacts}
							responseList.Sort() //order the list in starting with the closest

							Index := k
							if responseList.Len() < k {
								Index = responseList.Len()
							}
							contactChan <- responseList.Contacts[:Index]

						}
					}
				}(i)
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
			if !reflect.DeepEqual(contactList.Contacts, updatedlist) {
				contactList.Contacts = updatedlist
			} else {
				return nil
			}

		}
	}

}
*/

func (node *Node) Store(target *Contact, kv *KV) []error {

	contactMe := Contact{node.NodeID, node.IPAddress + ":" + node.Port, nil}
	// Channels:
	errorChan := make(chan error)
	//store
	node.St.Add(kv.Key, kv.Value)
	// get k closest contacts to the key
	contacts := node.LookupContact(target)
	// send STORE request to 3 of them
	for i := 0; i < 3; i++ {
		go func(index int) {
			contact := contacts.Contacts[index]
			n := Network{&contactMe, &contact}
			var err error
			err = nil
			err = n.SendStoreMessage(kv.Key, kv.Value)
			errorChan <- err
		}(i)
	}
	var errors []error
	for i := 0; i < 3; i++ {

		er := <-errorChan
		fmt.Println(er.Error())
		if er != nil {
			errors[i] = er
		}
	}

	return errors
}
