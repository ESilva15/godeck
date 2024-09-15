package deck

import (
	"fmt"

	"godeck/deck/cards"
)

// GetCards will return all cards of a given stack
func GetCards(boardId int,
	stackId int, cardIds []int) (cardList cards.CardDetailsList, err error) {

	cardList = make(cards.CardDetailsList, len(cardIds))
	for k, c := range cardIds {
		card, err := cards.Get(*state.client, boardId, stackId, c)
		if err != nil {
			fmt.Println("Could not get stack: ", err)
			return nil, err
		}

		cardList[k] = *card
	}

	return cardList, err
}

// CreateCard will create a new card in a the selected stack and board
func CreateCard(boardId int,
	stackId int, payload map[string]interface{}) (r string, err error) {
	return cards.Create(*state.client, boardId, stackId, payload)
}

// DeleteCard will delete a card given its coordinates
func DeleteCard(boardId int, stackId int, cardId int) (r string, err error) {
	return cards.Delete(*state.client, boardId, stackId, cardId)
}

// UpdateCard will delete a card given its coordinates
func UpdateCard(boardId int,
	stackId int, cardId int, payload map[string]interface{}) (r string, err error) {
	return cards.Update(*state.client, boardId, stackId, cardId, payload)
}

// AssignLabelToCard will assign a selected label to the selected card
func AssignLabelToCard(boardId int, stackId int,
	cardId int, labelId int) error {

	var card cards.CardDetails
	return card.AssignLabel(*state.client, boardId, stackId, cardId, labelId)
}

// UnassignLabelFromCard will remove a selected label from the selected card
func UnassignLabelFromCard(boardId int, stackId int,
	cardId int, labelId int) error {

	var card cards.CardDetails
	return card.UnasignLabel(*state.client, boardId, stackId, cardId, labelId)
}

// AssignUserToCard will assign a given user to a given card
func AssignUserToCard(boardId int, stackId int,
	cardId int, userId []string) error {

	var card cards.CardDetails
	for _, u := range userId {
		card.AssignUser(*state.client, boardId, stackId, cardId, u)
	}

	return nil
}

// UnassignUserFromCard will assign a given user to a given card
func UnassignUserFromCard(boardId int, stackId int,
	cardId int, userId []string) error {

	var card cards.CardDetails
	for _, u := range userId {
		card.UnassignUser(*state.client, boardId, stackId, cardId, u)
	}

	return nil
}

// ReorderCard will reorder a specific card
func ReorderCard(boardId int, stackId int,
	cardId int, order int, stack int) error {

	var card cards.CardDetails
	card.Reorder(*state.client, boardId, stackId, cardId, order, stack)

	return nil
}

// Archive a card
func ArchiveCard(boardId int, stackId int, cardId int) (string, error) {
  card, err := GetCards(boardId, stackId, []int{cardId})
  if err != nil {
    return "", err
  }

  c := card[0]

  payload := map[string]interface{} {
    "title": c.Title,
    "type": c.Type,
    "order": c.Order,
    "description": c.Description,
    "duedate": c.DueDate,
    "archived": true,
  }

  resp, err := UpdateCard(boardId, stackId, cardId, payload)

  return resp, nil
}
