package server

import (
	"net"
	"log"
)

type user struct {
	name string
	conn net.Conn
	deck []string
}

var users []user

func createUser(name string, conn net.Conn) *user {
	newUser := &user{name: name, conn: conn, deck: make([]string, 10)}
	users = append(users, *newUser)

	return newUser
}

func (u *user) startRead() {
	bs := make([]byte, 4096)
	for {
		n, err := u.conn.Read(bs)
		if err != nil {
			var n int
			for i, s := range users{    // Remove user from users
				if u.name == s.name {
					n = i
					break
				}
			}
			users = append(users[:n], users[n+1:]...)

			log.Println(u.name, "Exit")

			break
		}
		if 0 < n {
			data := bs[:n]
			commender(string(data))
		}
	}
}