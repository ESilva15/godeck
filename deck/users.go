package deck

import (
	"godeck/deck/users"
)

// GetUsers will return a list of the users of a board
func GetUsers(boardId int) ([]users.User, error) {
	board, err := GetBoardById(boardId)
	if err != nil {
		return []users.User{}, nil
	}

	return board.Users, nil
}
