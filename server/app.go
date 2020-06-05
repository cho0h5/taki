package server

import (
	"strings"
	"strconv"
	"log"
	"fmt"
)

func commender(data string) {
	d := strings.Split(data, ":")
	log.Print(data)

	s := strings.Split(d[1], " ")
	switch s[0] {
	case "START":
		shuffleMainDeck()

		for _, u := range users {
			n, _ := strconv.Atoi(s[1])
			u.deck = mainDeck[:n]
			mainDeck = mainDeck[n:]
		}
		broadcast("Game Start!\n")
		for _, u := range users {
			str := fmt.Sprint("Your Deck: ", u.deck)
			sendMessage(u.name, str)
		}
	}
}