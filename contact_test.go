package Kademlia

import (
	"fmt"
	"testing"
)

func TestRemoveDoupes(t *testing.T) {
	contacts := []Contact{}
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1:8014")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contacts = append(contacts, (NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "127.0.0.1:8016")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")))

	contacts = RemoveDupes(contacts)
	if len(contacts) != 6 {
		t.Error("BAD")
	}

}

func TestRemove(t *testing.T) {
	contacts := []Contact{}
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFEFFF00000000000000000000000000000001"), "127.0.0.1:8011")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111121100000000000000000000000000000000"), "127.0.0.1:8012")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111141300000000000000000000000000000000"), "127.0.0.1:8014")))
	contacts = append(contacts, (NewContact(NewKademliaID("1111511400000000000000000000000000000000"), "127.0.0.1:8015")))
	contacts = append(contacts, (NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")))
	contacts = append(contacts, (NewContact(NewKademliaID("2111511400000000000000000000000000000000"), "127.0.0.1:8016")))
	contactCandidates := ContactCandidates{contacts}
	target := NewContact(NewKademliaID("FFFFFFF300000000000000000000000000000000"), "127.0.0.1:8020")

	contactCandidates.RemoveContact(target)
	for m := 0; m < len(contactCandidates.Contacts); m++ {
		fmt.Println(contactCandidates.Contacts[m].StringDis())
	}
	if len(contactCandidates.Contacts) != 5 {
		t.Error("BAD")
	}

}
