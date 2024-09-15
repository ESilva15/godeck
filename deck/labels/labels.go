package labels

import (
	"encoding/json"
	"fmt"
	"godeck/deck/api"
)

type Label struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Color        string `json:"color"`
	BoardId      int    `json:"boardId"`
	CardId       int    `json:"cardId"`
	LastModified int    `json:"lastModified"`
	ETag         string `json:"ETag"`
}

type DeckLabelInterface interface {
	GetById(boardId int) ([]Label, error)
	Show()
}

func (l *Label) Show() {
	toPrint, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		fmt.Println("Unable to marshal data: ", err)
		return
	}

	fmt.Println(string(toPrint))
}

func (l *Label) Get(client api.DeckAPI, boardId int, labelId int) error {
	url := fmt.Sprintf("/boards/%d/labels/%d", boardId, labelId)

	rawData, err := client.Get(url, nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return err
	}

	err = json.Unmarshal([]byte(rawData), l)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return err
	}

	return nil
}

func (l *Label) Create(client api.DeckAPI, boardId int, 
  payload map[string]interface{}) error {

  url := fmt.Sprintf("/boards/%d/labels", boardId)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = client.Post(url, jsonPayload)

  return err
}

func (l *Label) Update(client api.DeckAPI, boardId int, labelId int,
  payload map[string]interface{}) error {

  url := fmt.Sprintf("/boards/%d/labels/%d", boardId, labelId)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

  return err
}

func (l *Label) Delete(client api.DeckAPI, boardId int, stackId int) error {
  url := fmt.Sprintf("/boards/%d/stacks/%d", boardId, stackId)

  _, err := client.Delete(url, nil)

  return err
}
