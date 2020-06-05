package server

import (
	"fmt"
	"log"
	"net"
)


const (
	waiting = iota
	started
)

var serverState = waiting

// StartServer  start server
func StartServer() {
	fmt.Println("StartServer")

	l, err := net.Listen("tcp", ":5252")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		if serverState == waiting {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go connHandler(conn)
		}
	}
}

func connHandler(conn net.Conn) {
	bs := make([]byte, 4096)
	n, _ := conn.Read(bs)	// Read user name
	name := string(bs[:n])

	newUser := createUser(name, conn)
	log.Println(newUser, "entered")

	newUser.startRead()
}

func sendMessage(name, str string) {
	for _, u := range users {
		if name == u.name {
			u.conn.Write([]byte(str))
			break
		}
	}
}

func broadcast(str string) {
	for _, u := range users {
		u.conn.Write([]byte(str))
	}
}