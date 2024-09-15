package acl

import (
	"encoding/json"
	"fmt"

	"godeck/deck/api"
	"godeck/deck/users"
)

// The Participant can either be a Participant struct or a simple string
type ACLRule struct {
	Participant      users.ParticipantWrapper `json:"participant"`
	Type             int                      `json:"type"`
	BoardId          int                      `json:"boardId"`
	PermissionEdit   bool                     `json:"permissionEdit"`
	PermissionShare  bool                     `json:"permissionShare"`
	PermissionManage bool                     `json:"permissionManager"`
	Owner            bool                     `json:"owner"`
	Id               int                      `json:"id"`
}

func (a *ACLRule) Create(client api.DeckAPI, boardId int,
	payload map[string]interface{}) (string, error) {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/acl", boardId)
	resp, err := client.Post(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (a *ACLRule) Update(client api.DeckAPI, boardId int, aclId int,
	payload map[string]interface{}) (string, error) {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/boards/%d/acl/%d", boardId, aclId)
	resp, err := client.Put(url, jsonPayload)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (a *ACLRule) Delete(client api.DeckAPI, boardId int, aclId int) (string, error) {
	url := fmt.Sprintf("/boards/%d/acl/%d", boardId, aclId)
	resp, err := client.Delete(url, nil)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
