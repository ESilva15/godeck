package main

import (
	"godeck/cmd"
	"godeck/deck"
)

func main() {
	deck.InitDeck()

	cmd.Execute()
}
