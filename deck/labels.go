package deck

import (
	"fmt"

	"godeck/deck/labels"
)

// GetLabels will return the labels belonging to a board
func GetLabels(boardId int) ([]labels.Label, error) {
	board, err := GetBoardById(boardId)
	if err != nil {
		fmt.Println("Could not retrieve boards: ", err)
		return []labels.Label{}, nil
	}

	return board.Labels, nil
}

// GetLabel will return a detailed label of a board
func GetLabel(
	boardId int, labelId int) (labels.Label, error) {

	var label labels.Label

	err := label.Get(*state.client, boardId, labelId)
	if err != nil {
		fmt.Println("Unable to retrieve label: ", err)
		return label, err
	}

	return label, nil
}

func CreateLabel(boardId int, payload map[string]interface{}) (error) {
  var label labels.Label
  return label.Create(*state.client, boardId, payload)
}

func UpdateLabel(boardId int, labelId int,
  payload map[string]interface{}) error {
  var label labels.Label
  return label.Update(*state.client, boardId, labelId, payload)
}

func DeleteLabel(boardId int, labelId int) error {
  var label labels.Label
  return label.Delete(*state.client, boardId, labelId)
}
