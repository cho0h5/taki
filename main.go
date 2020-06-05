package main

import (
	"fmt"
	"os"

	"github.com/dominoyh5/taki/client"
	"github.com/dominoyh5/taki/server"
)

const (
	serverFlag = 1 + iota
	clientFlag
	exitFlag
)

func main() {
	fmt.Println("Welcome to TAKI")
	fmt.Println("1. Server")
	fmt.Println("2. Client")
	fmt.Println("3. Exit")
	fmt.Print("Enter: ")

	for {
		var mode int
		fmt.Scan(&mode)

		switch mode {
		case serverFlag:
			server.StartServer()
			os.Exit(0)
		case clientFlag:
			client.StartClient()
			os.Exit(0)
		case exitFlag:
			os.Exit(0)
		}

		fmt.Print("\nEnter again: ")
	}
}
