package Kademlia

import (
	"fmt"
	"testing"
)

func TestUpdateList(t *testing.T) {
	contacts := []Contact{}
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1:8014")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contacts = append(contacts, (NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "127.0.0.1:8016")))
	var ca ContactCandidates
	ca.Contacts = contacts

	contactsB := []Contact{}
	contactsB = append(contactsB, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "127.0.0.1:8004")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "127.0.0.1:8005")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "127.0.0.1:8016")))
	contactsB = append(contactsB, (NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "127.0.0.1t:8006")))
	var cb ContactCandidates
	cb.Contacts = contactsB

	updated := UpdateList(&ca, &cb)

	//TARGET
	target := NewContact(NewKademliaID("1111111A00000000000000000000000000000000"), "127.0.0.1:8005")
	updated.CalDistance(target)

	for m := 0; m < len(updated.Contacts); m++ {
		fmt.Println(updated.Contacts[m].StringDis())
	}
}
