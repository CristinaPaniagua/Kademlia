package Kademlia

import (
	"fmt"
	"sort"
)

// Contact definition
// stores the KademliaID, the ip address and the distance
type Contact struct {
	ID       *KademliaID
	Address  string
	distance *KademliaID
}

// NewContact returns a new instance of a Contact
func NewContact(id *KademliaID, address string) Contact {
	return Contact{id, address, nil}
}

// CalcDistance calculates the distance to the target and
// fills the contacts distance field
func (contact *Contact) CalcDistance(target *KademliaID) {
	contact.distance = contact.ID.CalcDistance(target)
}

// Less returns true if contact.distance < otherContact.distance
func (contact *Contact) Less(otherContact *Contact) bool {
	return contact.distance.Less(otherContact.distance)
}

// String returns a simple string representation of a Contact
func (contact *Contact) String() string {
	return fmt.Sprintf(`contact("%s", "%s")`, contact.ID, contact.Address)
}

// String returns a simple string representation of a Contact including distance
func (contact *Contact) StringDis() string {
	var distance string
	if contact.distance == nil {
		distance = "not calculated"
	} else {
		distance = contact.distance.String()
	}
	return fmt.Sprintf(`contact("%s", "%s"), Distance: "%s")`, contact.ID, contact.Address, distance)
}

// ContactCandidates definition
// stores an array of Contacts
type ContactCandidates struct {
	Contacts []Contact
}

// calculate distance to a target for a list of contacts
func (candidates *ContactCandidates) CalDistance(target Contact) {

	for m := 0; m < len(candidates.Contacts); m++ {
		contact := candidates.Contacts[m]
		contact.CalcDistance(target.ID)
		candidates.Contacts[m].distance = contact.distance
	}

}

// Append an array of Contacts to the ContactCandidates
func (candidates *ContactCandidates) Append(contacts []Contact) {
	candidates.Contacts = append(candidates.Contacts, contacts...)
}

// remove an contact from the ContactCandidates
func (candidates *ContactCandidates) RemoveContact(contact Contact) {
	//func (n *shortList) RemoveNode(node *NetworkNode) {
	//		for i := 0; i < n.Len(); i++ {
	//			if bytes.Compare(n.Nodes[i].ID, node.ID) == 0 {
	//				n.Nodes = append(n.Nodes[:i], n.Nodes[i+1:]...)
	//				return
	//			}
	//		}
	//	}

}

func RemoveDupes(contacts []Contact) []Contact {
	//fmt.Printf("BEFORE *******************************")
	//for m := 0; m < len(contacts); m++ {
	//	fmt.Println(contacts[m].String())
	//}
	var unduped_slice []Contact
	for i := 0; i < len(contacts); i++ {
		current := contacts[i]
		fmt.Println("current: " + current.String())
		equal := false
		for j := i + 1; j < len(contacts); j++ {
			//fmt.Println("comparing with: " + contacts[j].String())
			if current.ID.Equals(contacts[j].ID) {

				equal = true
			}
			//fmt.Println(equal)
		}
		if !equal {
			unduped_slice = append(unduped_slice, current)
		}
	}
	/*fmt.Println("AFTER *******************************")
	for m := 0; m < len(unduped_slice); m++ {
		fmt.Println(unduped_slice[m].String())
	} */
	return unduped_slice
}

// GetContacts returns the first count number of Contacts
func (candidates *ContactCandidates) GetContacts(count int) []Contact {
	return candidates.Contacts[:count]
}

// Sort the Contacts in ContactCandidates
func (candidates *ContactCandidates) Sort() {
	sort.Sort(candidates)
}

// Len returns the length of the ContactCandidates
func (candidates *ContactCandidates) Len() int {
	return len(candidates.Contacts)
}

// Swap the position of the Contacts at i and j
// WARNING does not check if either i or j is within range
func (candidates *ContactCandidates) Swap(i, j int) {
	candidates.Contacts[i], candidates.Contacts[j] = candidates.Contacts[j], candidates.Contacts[i]
}

// Less returns true if the Contact at index i is smaller than
// the Contact at index j
func (candidates *ContactCandidates) Less(i, j int) bool {
	return candidates.Contacts[i].Less(&candidates.Contacts[j])
}
