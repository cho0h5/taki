package server

import (
	"math/rand"
	"time"
)

var mainDeck = []string{
	"R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9",
	"Y1", "Y2", "Y3", "Y4", "Y5", "Y6", "Y7", "Y8", "Y9",
	"G1", "G2", "G3", "G4", "G5", "G6", "G7", "G8", "G9",
	"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9",
}
var tableDeck = []string{}

func shuffleMainDeck() {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(mainDeck), func(i, j int) {mainDeck[i], mainDeck[j] = mainDeck[j], mainDeck[i]})
}