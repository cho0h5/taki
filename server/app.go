package server

import (
	"strings"
	"strconv"
	"log"
	"fmt"
)

func commender(u *user, data string) {
	d := strings.Split(data, " ")
	log.Print(data)

	s := strings.Split(d[1], " ")
	switch s[0] {
	case "START":
		n, _ := strconv.Atoi(s[1])
		gameStart(n)
	case "push":
		push()
	case "pull":
		pull()
	}
}

func gameStart(numberOfCard int) {
	broadcast("Game Start!\n")

	shuffleMainDeck()
	initializeTableDeck()

	for _, u := range users {
		u.deck = mainDeck[:numberOfCard]
		mainDeck = mainDeck[numberOfCard:]
	}

	for _, u := range users {    // TODO
		str := fmt.Sprint("Your Deck: ", u.deck)
		sendMessage(u.name, str)
	}
}

func push() {
}

func pull() {
}