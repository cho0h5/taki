package main

import (
	"fmt"
	"os"
)

const (
	Server = 1 + iota
	Client
	Exit
)

func main() {
	fmt.Println("Welcome to TAKI\n")
	fmt.Println("1. Server")
	fmt.Println("2. Client")
	fmt.Println("3. Exit\n")
	fmt.Print("Enter: ")

	for {
		var mode int
		fmt.Scan(&mode)

		switch mode {
		case Server:
			StartServer()
			os.Exit(0)
		case Client:
			StartClient()
			os.Exit(0)
		case Exit:
			os.Exit(0)
		}

		fmt.Print("\nEnter again: ")
	}
}
