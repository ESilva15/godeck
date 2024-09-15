package deck

import (
	"sync"

	"godeck/config"
	"godeck/deck/api"
)

type deckState struct {
	client *api.DeckAPI
}

var (
	state  *deckState = nil
	once   sync.Once
)

func InitDeck() {
	once.Do(func() {
    config := config.GetInstance()

		state = &deckState{
			client: &api.DeckAPI{
				URL:     config.Url,
				User:    config.Auth.User,
				Pass:    config.Auth.AppPassword,
				LogedIn: false,
				Token:   "",
			},
		}
	})
}
