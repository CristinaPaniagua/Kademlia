package Kademlia

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func runNodeEndpoints() {

	router := mux.NewRouter()

	router.HandleFunc("/objects", GetObjects).Methods("GET")
	router.HandleFunc("/objects/{hash}", PostObjects).Methods("POST")
	http.Handle("/", router)

	fmt.Println("API running...")
	http.ListenAndServe(":4000", router)
}

type valueReply struct {
	value []byte `json:"value"`
}

func GetObjects(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	hash := vars["hash"]
	fmt.Println(hash)
	n := Node{}
	reply, found, err := n.FindValue(hash)
	if err == nil {
		if found {
			valuereply := valueReply{reply.Val}
			response.Header().Add("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(valuereply)

		} else {
			//TODO Handle not found
		}

	} else {
		//TODO Handle error
	}

}

type Object struct {
	hash  string
	value []byte
}

func PostObjects(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	hash := vars["hash"]
	value := vars["value"]

	n := Node{}
	n.StoreKV(hash, []byte(value))

	response.Header().Add("Location", "/objects/{"+hash+"/")
	response.WriteHeader(http.StatusCreated)

}
