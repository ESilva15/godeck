package cards

import (
	"encoding/json"
	"errors"
	"fmt"

	"godeck/deck/api"
	"godeck/deck/labels"
	"godeck/deck/users"
)

type AssignedUser struct {
	Id          int         `json:"id"`
	Participant users.Owner `json:"participant"`
	CardId      int         `json:"cardId"`
	Type        int         `json:"type"`
}

// TODO
// Merge both Card structs so that we can mock them more easily if necessary
// I could also make it so instead of relying on automatic marshalling for
// everything I can create "cards" with a given boardId and stackId and pass
// them around instead of passing the three values all the time

// The stacks and cards endpoints retrieve differente information regarding
// the cards. On the stacks the cards use a string to represent the owner while
// on the cards it uses the user.Owner struct
// This is how the card is represented on the stacks card list
type Card struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	StackId         int            `json:"stackId"`
	Type            string         `json:"type"`
	LastModified    int            `json:"lastModified"`
	LastEditor      string         `json:"lastEditor"`
	CreatedAt       int            `json:"createdAt"`
	Labels          []labels.Label `json:"labels"`
	AssignedUsers   []AssignedUser `json:"assignedUsers"`
	Attachments     string         `json:"attachments"`
	AttachmentCount int            `json:"attachmentCount"`
	Owner           string         `json:"owner"`
	Order           int            `json:"order"`
	Archived        bool           `json:"archived"`
	Done            bool           `json:"done"`
	DueDate         string         `json:"duedata"`
	DeletedAt       int            `json:"deletedAt"`
	CommentsUnread  int            `json:"commentsUnread"`
	CommentsCount   int            `json:"commentsCount"`
	ETag            string         `json:"Etag"`
	Overdue         int            `json:"overdue"`
}
type CardList []Card

// This is how the card details are represented
type CardDetails struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	StackId         int            `json:"stackId"`
	Type            string         `json:"type"`
	LastModified    int            `json:"lastModified"`
	LastEditor      string         `json:"lastEditor"`
	CreatedAt       int            `json:"createdAt"`
	Labels          []labels.Label `json:"labels"`
	AssignedUsers   []AssignedUser `json:"assignedUsers"`
	Attachments     []string       `json:"attachments"`
	AttachmentCount int            `json:"attachmentCount"`
	Owner           string         `son:"owner"`
	Order           int            `json:"order"`
	Archived        bool           `json:"archived"`
	Done            string         `json:"done"`
	DueDate         string         `json:"duedata"`
	DeletedAt       int            `json:"deletedAt"`
	CommentsUnread  int            `json:"commentsUnread"`
	CommentsCount   int            `json:"commentsCount"`
	ETag            string         `json:"Etag"`
	Overdue         int            `json:"overdue"`
}
type CardDetailsList []CardDetails

func (c *CardDetails) Show() {
	toPrint, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("Unable to marshal data: ", err)
		return
	}

	fmt.Println(string(toPrint))
}

// Will assign a label to the card
func (c *CardDetails) AssignLabel(client api.DeckAPI, boardId int, stackId int,
	cardId int, labelId int) error {
	/*
			  curl -X PUT 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/89/assignLabel' \
			   -H 'Accept: application/json' \
			   -H "Content-Type: application/json" \
			   -H 'OCS-APIRequest: true' \
		     --data-raw "{\"labelId\":60}"
	*/

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d/assignLabel",
		boardId, stackId, cardId)

	jsonPayload, err := json.Marshal(map[string]int{"labelId": labelId})
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

	return err
}

func (c *CardDetails) UnasignLabel(client api.DeckAPI, boardId int, stackId int,
	cardId int, labelId int) error {
	/*
			  curl -X PUT 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/89/removeLabel' \
			   -H 'Accept: application/json' \
			   -H "Content-Type: application/json" \
			   -H 'OCS-APIRequest: true' \
		     --data-raw "{\"labelId\":60}"
	*/

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d/removeLabel",
		boardId, stackId, cardId)

	jsonPayload, err := json.Marshal(map[string]int{"labelId": labelId})
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

	return err
}

func (c *CardDetails) AssignUser(client api.DeckAPI, boardId int, stackId int,
	cardId int, userId string) error {
	/*
	  curl -X PUT 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/89/assignUser' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true' \
	    --data-raw "{\"userId\":60}"
	*/

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d/assignUser",
		boardId, stackId, cardId)

	jsonPayload, err := json.Marshal(map[string]string{"userId": userId})
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

	return err
}

func (c *CardDetails) UnassignUser(client api.DeckAPI, boardId int, stackId int,
	cardId int, userId string) error {
	/*
	  curl -X PUT 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/89/assignUser' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true' \
	    --data-raw "{\"userId\":60}"
	*/

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d/unassignUser",
		boardId, stackId, cardId)

	jsonPayload, err := json.Marshal(map[string]string{"userId": userId})
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

	return err
}

func (c *CardDetails) Reorder(client api.DeckAPI, boardId int, stackId int,
	cardId int, order int, stack int) error {

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d/reorder",
		boardId, stackId, cardId)

	jsonPayload, err := json.Marshal(map[string]int{
		"order":   order,
		"stackId": stack,
	})
	if err != nil {
		return err
	}

	_, err = client.Put(url, jsonPayload)

	return err
}

func Get(client api.DeckAPI, boardId int, stackId int, cardId int) (c *CardDetails, err error) {
	/*
	  curl -X GET 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/89' \
	    -H 'Accept: application/json' \
	    -H "Content-Type: application/json" \
	    -H 'OCS-APIRequest: true' | jq
	*/

	endpoint := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d",
		boardId, stackId, cardId)
	rawData, err := client.Get(endpoint, nil)
	if err != nil {
		fmt.Println("Could not retrieve the data: ", err)
		return nil, err
	}

	var card CardDetails
	err = json.Unmarshal([]byte(rawData), &card)
	if err != nil {
		fmt.Println("Cannot unmarshal json: ", err)
		return nil, err
	}

	if &card == nil {
		inputData := fmt.Sprintf("New card is nil:\n  Board: %d, Stack: %d, Card: %d.",
			boardId, stackId, cardId)
		return nil, errors.New(inputData)
	}

	return &card, nil
}

func Create(client api.DeckAPI, boardId int, stackId int,
	payload map[string]interface{}) (r string, err error) {
	/*
				Example cURL request:
		    curl -X POST 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards' \
		      -H 'Accept: application/json' \
		      -H "Content-Type: application/json" \
		      -H 'OCS-APIRequest: true' \
		      --data-raw '{"title":"Work on this","type":"plain","order":1,"description":"A very intriguing description\nmultiplelines tho?","duedate": "'$(date -d"+1 week" +"%Y-%m-%dT%H:%M:%S+00:00")'"}'	
  */

	// Maybe add some checks to the input of the user here
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards", boardId, stackId)
	resp, err := client.Post(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func Delete(client api.DeckAPI, boardId int, stackId int, cardId int) (r string, err error) {
	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d", boardId, stackId, cardId)
	resp, err := client.Delete(url, nil)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func Update(client api.DeckAPI, boardId int,
	stackId int, cardId int, payload map[string]interface{}) (r string, err error) {
	// For some stupid ass reason we need provide the owner into the data. Even tho
	// The API documentation doesn't say such a thing
	/*
	   curl -X PUT 'https://username:password@cloud.org/index.php/apps/deck/api/v1.0/boards/15/stacks/37/cards/176' \
	   -H 'Accept: application/json' \
	   -H "Content-Type: application/json" \
	   -H 'OCS-APIRequest: true' \
	   --data-raw '{"title":"Work on this","type":"plain","order":1,"description":"A very intriguing description\nmultiplelines tho?","duedate":"2024-12-24T19:29:30+00:00","owner":"username"}'

	   The duedate format is in ISO-8601, we can use the following date command to
	   get the date:
	     date +"%Y-%m-%dT%H:%M:%S%:z"
	   For example, we can also do:
	     date -d "now + 1 week" +"%Y-%m-%dT%H:%M:%S%:z"
	*/

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/stacks/%d/cards/%d", boardId, stackId, cardId)
	resp, err := client.Put(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
