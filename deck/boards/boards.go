package boards

import (
	"encoding/json"
	"fmt"

	"godeck/deck/acl"
	"godeck/deck/api"
	"godeck/deck/labels"
	"godeck/deck/users"
)

type Stack struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	BoardID      int    `json:"boardId"`
	DeletedAt    int    `json:"deleteAt"`
	LastModified int    `json:"lastModified"`
	Order        int    `json:"order"`
	ETag         string `json:"ETag"`
}

type Settings struct {
	NotifyDue string `json:"notify-due"`
	Calendar  bool   `json:"calendar"`
}

type Board struct {
	Id             int               `json:"id"`
	Acl            []acl.ACLRule     `json:"acl"`
	Users          []users.User      `json:"users"`
	Labels         []labels.Label    `json:"labels"`
	Title          string            `json:"title"`
	Owner          users.Owner       `json:"owner"`
	Color          string            `json:"color"`
	Archived       bool              `json:"archived"`
	Shared         int               `json:"shared"`
	Permissions    users.Permissions `json:"permissions"`
	ActiveSessions []int             `json:"activeSessions"`
	LastModified   int               `json:"lastModified"`
	DeletedAt      int               `json:"deletedAt"`
	Settings       Settings          `json:"settings"`
	ETag           string            `json:"ETag"`
	Stacks         []Stack           `json:"stacks"`
}
type Boards []Board

func (b *Board) Show() {
	toPrint, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		fmt.Println("Unable to marshal data: ", err)
		return
	}

	fmt.Println(string(toPrint))
}

func (b *Board) Update(client api.DeckAPI, boardId int,
	payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d", boardId)
	resp, err := client.Put(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func GetBoards(client api.DeckAPI) (data map[int]*Board, err error) {
	/*
	  curl -X GET 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true'
	*/
	rawData, err := client.Get("/boards", nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return nil, err
	}

	var boards []Board
	err = json.Unmarshal([]byte(rawData), &boards)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return nil, err
	}

	boardsMap := make(map[int]*Board)
	for _, v := range boards {
		boardsMap[v.Id] = &v
	}

	return boardsMap, err
}

func GetBoard(client api.DeckAPI, boardId int) (data *Board, err error) {
	/*
	  curl -X GET 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/12' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true'
	*/

	rawData, err := client.Get(fmt.Sprintf("/boards/%d", boardId), nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return nil, err
	}

	var board Board
	err = json.Unmarshal([]byte(rawData), &board)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return nil, err
	}

	return &board, err
}

func CreateBoard(client api.DeckAPI, title string, color string) (data string, err error) {
	/*
		Example cURL request:
		curl -X POST 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0' \
		 -H 'Accept: application/json' \
		 -H "Content-Type: application/json" \
		 -H 'OCS-APIRequest: true' \
		 --data-raw '{"title":"new","color":"ff00ff"}'
	*/

	payload := map[string]string{
		"title": title,
		"color": color,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := client.Post("/boards", jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

// UpdateBoard allows us to update board fields
func UpdateBoard(boardId int) {

}

func DeleteBoardById(client api.DeckAPI, boardId int) (err error) {
	/*
	  It will return 403 if the board has already been deleted, so mind that

	  curl -X DELETE 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/25' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true'
	*/
	endpoint := fmt.Sprintf("/boards/%d", boardId)
	_, err = client.Delete(endpoint, nil)
	if err != nil {
		return err
	}

	return nil
}

// TODO
// THIS FUNCTION IS NOT WORKING, HAVE TO FIGURE OUT WHY
// func UndoDeleteBoardById(client api.DeckAPI, boardId int) (err error) {
// 	endpoint := fmt.Sprintf("/boards/%d/undo_delete", boardId)
// 	_, err = client.Post(endpoint, nil)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func AddNewACLRule(client api.DeckAPI, boardId int) {
}

func UpdateACLRule(client api.DeckAPI, boardId int, aclId int) {
}

func DeleteACLRule(client api.DeckAPI, boardId int, aclId int) {
}
