package server

import (
	"math/rand"
	"time"
)

var mainDeck = []string{
	"R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9", "RP", "RP", "RR", "RJ", "RT", "RT",
	"Y1", "Y2", "Y3", "Y4", "Y5", "Y6", "Y7", "Y8", "Y9", "YP", "YP", "YR", "YJ", "YT", "YT",
	"G1", "G2", "G3", "G4", "G5", "G6", "G7", "G8", "G9", "GP", "GP", "GR", "GJ", "GT", "GT",
	"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "BP", "BP", "BR", "BJ", "BT", "BT",
	"AC", "AC",
}
var tableDeck = []string{}

func shuffleMainDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mainDeck), func(i, j int) {mainDeck[i], mainDeck[j] = mainDeck[j], mainDeck[i]})
}

func initializeTableDeck() {
	tableDeck = append(tableDeck, mainDeck[:1]...)
	mainDeck = mainDeck[1:]
}