package main

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/readline.v1"
)

func main() {

	fmt.Println("Node started")
	//take arguments
	//args := os.Args[1:]
	contact := NewContact()
	node := Node{}

	//REDO WITH MY METHODS
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}
		input := strings.Split(line, " ")
		switch input[0] {
		case "help":
			displayHelp()
		case "store":
			if len(input) != 2 {
				displayHelp()
				continue
			}
			id, err := node.Store([]byte(input[1]))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Stored ID: " + id)
		case "get":
			if len(input) != 2 {
				displayHelp()
				continue
			}
			data, exists, err := dht.Get(input[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Searching for", input[1])
			if exists {
				fmt.Println("..Found data:", string(data))
			} else {
				fmt.Println("..Nothing found for this key!")
			}
		case "info":
			nodes := dht.NumNodes()
			self := dht.GetSelfID()
			addr := dht.GetNetworkAddr()
			fmt.Println("Addr: " + addr)
			fmt.Println("ID: " + self)
			fmt.Println("Known Nodes: " + strconv.Itoa(nodes))
		}
	}
}

func displayFlagHelp() {
	fmt.Println(`cli-example
Usage:
	cli-example --port [port]
Options:
	--help Show this screen.
	--ip=<ip> Local IP [default: 0.0.0.0]
	--port=[port] Local Port [default: 0]
	--bip=<ip> Bootstrap IP
	--bport=<port> Bootstrap Port
	--stun=<bool> Use STUN protocol for public addr discovery [default: true]`)
}

func displayHelp() {
	fmt.Println(`
help - This message
store <message> - Store a message on the network
get <key> - Get a message from the network
info - Display information about this node
	`)
}
