package deck

import (
  "godeck/deck/acl"
)

// CreateACLRule
func CreateACLRule(boardId int, payload map[string]interface{}) (string, error) {
  var acl acl.ACLRule
	return acl.Create(*state.client, boardId, payload)
}

// UpdateACLRule
func UpdateACLRule(boardId int, aclId int, payload map[string]interface{}) (string, error) {
  var acl acl.ACLRule
	return acl.Update(*state.client, boardId, aclId, payload)
}

// DeleteACLRule
func DeleteACLRule(boardId int, aclId int) (string, error) {
  var acl acl.ACLRule
	return acl.Delete(*state.client, boardId, aclId)
}
