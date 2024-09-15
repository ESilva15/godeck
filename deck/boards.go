package deck

import (
	"fmt"

	"godeck/deck/boards"
)

// CreateBoard given a title and a hex color code creates a new board
func CreateBoard(title string, color string) {
	_, err := boards.CreateBoard(*state.client, title, color)
	if err != nil {
		fmt.Println("Unable to create board: ", err)
		return
	}
}

// DeleteBoards deletes a list of boards from Deck
func DeleteBoards(ids []int) (failed map[int]error) {
	failed = make(map[int]error, len(ids))

	for _, id := range ids {
		err := boards.DeleteBoardById(*state.client, id)
		if err != nil {
			failed[id] = err
		}
	}

	return failed
}

// UndoDeleteBoards reverses the deletion of the boards
// func UndoDeleteBoards(ids []int) (failed map[int]error) {
// 	failed = make(map[int]error, len(ids))
//
//   deckApi := api.DeckAPI{}
// 	for _, id := range ids {
// 		err := boards.UndoDeleteBoardById(deckApi, id)
// 		if err != nil {
// 			failed[id] = err
// 		}
// 	}
//
// 	return failed
// }

// GetBoards will return all the boards in the Deck instance
func GetBoards() (data map[int]*boards.Board, err error) {
	b, err := boards.GetBoards(*state.client)
	if err != nil {
		fmt.Println("Could not get boards: ", err)
		return nil, err
	}
	return b, err
}

// GetBoardById will return a single board
func GetBoardById(boardId int) (*boards.Board, error) {
	b, err := boards.GetBoard(*state.client, boardId)
	if err != nil {
		fmt.Println("Failed to retrieve board: ", err)
		return nil, err
	}

	return b, nil
}

// UpdateBoard
func UpdateBoard(boardId int, payload map[string]interface{}) (string, error) {
  var board boards.Board
  return board.Update(*state.client, boardId, payload)
}
