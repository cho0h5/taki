package main

import (
	"fmt"
	"log"
	"net"
)

func StartClient() {
	fmt.Println("StartClient")
	conn, err := net.Dial("tcp", ":5252")
	if nil != err {
		log.Println(err)
	}

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Server send: ", string(data[:n]))
		}
	}()

	var s string
	for {
		fmt.Scanln(&s)
		_, err := conn.Write([]byte(s))
		if err != nil {
			log.Println(err)
		}
	}
}
