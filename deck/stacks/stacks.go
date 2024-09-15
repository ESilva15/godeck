package stacks

import (
	"encoding/json"
	"fmt"

	"godeck/deck/api"
	"godeck/deck/cards"
)

type Stack struct {
	Id           int                   `json:"id"`
	Title        string                `json:"title"`
	BoardId      int                   `json:"boardId"`
	DeletedAt    int                   `json:"deletedAt"`
	LastModified int                   `json:"lastModified"`
	Cards        cards.CardDetailsList `json:"cards"`
	Order        int                   `json:"order"`
	ETag         string                `json:"ETag"`
}

func (b *Stack) Show() {
	toPrint, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		fmt.Println("Unable to marshal data: ", err)
		return
	}

	fmt.Println(string(toPrint))
}

func (s *Stack) Create(client api.DeckAPI, boardId int,
	payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/stacks", boardId)
	resp, err := client.Post(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (s *Stack) Update(client api.DeckAPI, boardId int, stackId int,
	payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/stacks/%d", boardId, stackId)
	resp, err := client.Put(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (s *Stack) Delete(client api.DeckAPI, boardId int, stackId int) (string, error) {
	url := fmt.Sprintf("/boards/%d/stacks/%d", boardId, stackId)
	resp, err := client.Delete(url, nil)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}


func GetById(client api.DeckAPI, boardId int, stackId int) (data *Stack, err error) {
	/*
		curl -X GET 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37' \
		  -H 'Accept: application/json' \
		  -H "Content-Type: application/json" \
		  -H 'OCS-APIRequest: true'
	*/

	endpoint := fmt.Sprintf("/boards/%d/stacks/%d", boardId, stackId)
	rawData, err := client.Get(endpoint, nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return nil, err
	}

	var stack Stack
	err = json.Unmarshal([]byte(rawData), &stack)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return nil, err
	}

	return &stack, nil
}

func Get(client api.DeckAPI, boardId int) (data map[int]*Stack, err error) {
	/*
		curl -X GET 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks' \
		  -H 'Accept: application/json' \
		  -H "Content-Type: application/json" \
		  -H 'OCS-APIRequest: true'
	*/

	endpoint := fmt.Sprintf("/boards/%d/stacks", boardId)
	rawData, err := client.Get(endpoint, nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return nil, err
	}

	var stacks []Stack
	err = json.Unmarshal([]byte(rawData), &stacks)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return nil, err
	}

	stacksMap := make(map[int]*Stack)
	for _, v := range stacks {
		stacksMap[v.Id] = &v
	}

	return stacksMap, err
}

func Create(client api.DeckAPI, boardId int) {
}

func Update(client api.DeckAPI, boardId int, stackId int) {
}

func Delete(client api.DeckAPI, boardId int, stackId int) {
}

func GetArchived(client api.DeckAPI, boardId int) {
}
