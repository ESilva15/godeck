package users

import (
	"encoding/json"
	"fmt"
)

type User struct {
	PrimaryKey  string `json:"primaryKey"`
	UId         string `json:"uid"`
	Displayname string `json:"displayname"`
}

type Participant struct {
	OnlyName    string `json:"participant"`
	PrimaryKey  string `json:"primaryKey,omitempty"`
	UID         string `json:"uid,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Type        int    `json:"type,omitempty"`
}

type ParticipantWrapper struct {
	Participant
	OnlyStr bool `json:"-"`
}

type Owner struct {
	PrimaryKey  string `json:"primaryKey"`
	Uid         string `json:"uid"`
	Displayname string `json:"displayname"`
	Type        int    `json:"type"`
}

type Permissions struct {
	Read   bool `json:"PERMISSION_READ"`
	Edit   bool `json:"PERMISSION_EDIT"`
	Manage bool `json:"PERMISSION_MANAGE"`
	Share  bool `json:"PERMISSION_SHARE"`
}

func (p *ParticipantWrapper) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err == nil {
		p.OnlyName = str
    p.OnlyStr = true
		return nil
	}

  p.OnlyStr = false
	return json.Unmarshal(data, &p.Participant)
}

func (p *ParticipantWrapper) MarshalJSON() ([]byte, error) {
  if p.OnlyStr {
    return []byte(fmt.Sprintf(`{"participant": "%s"}`, p.OnlyName)), nil
  }

  return json.Marshal(p.Participant)
}

func (u *User) Show() {
	toPrint, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("Unable to marshal data: ", err)
		return
	}

	fmt.Println(string(toPrint))
}
