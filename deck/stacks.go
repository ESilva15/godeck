package deck

import (
	"fmt"

	"godeck/deck/stacks"
)

// GetStacks will return all the stacks from a board
func GetStacks(boardId int) (data map[int]*stacks.Stack, err error) {
	s, err := stacks.Get(*state.client, boardId)
	if err != nil {
		fmt.Println("Could not get stacks: ", err)
		return nil, err
	}

	return s, err
}

// GetStack returns a single stack
func GetStack(boardId int, stackId int) (*stacks.Stack, error) {
  stack, err := stacks.GetById(*state.client, boardId, stackId)
  if err != nil {
    return nil, err
  }

  return stack, nil
}

// CreateStack will create a new stack given the payload
func CreateStack(boardId int, payload map[string]interface{}) (string, error) {
  var stack stacks.Stack
  return stack.Create(*state.client, boardId, payload)
}

// UpdateStack
func UpdateStack(boardId int, stackId int,
  payload map[string]interface{}) (string, error) {
  var stack stacks.Stack
  return stack.Update(*state.client, boardId, stackId, payload)
}

// DeleteStack
func DeleteStack(boardId int, stackId int) (string, error) {
  var stack stacks.Stack
  return stack.Delete(*state.client, boardId, stackId)
}
