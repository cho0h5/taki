package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var name string

func StartClient() {
	fmt.Print("\nEnter your name: ")
	fmt.Scanln(&name)

	conn, err := net.Dial("tcp", ":5252")
	if nil != err {
		log.Println(err)
	}

	go func() {
		_, err := conn.Write([]byte(name))
		if err != nil {
			log.Fatal(err)
		}

		data := make([]byte, 4096)
		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(string(data[:n]))
		}
	}()

	in := bufio.NewReader(os.Stdin)
	for {
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		_, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
		}
	}
}
