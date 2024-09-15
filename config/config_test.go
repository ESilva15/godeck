package config

import ( 
  "testing"
)

func TestGetInstance(t *testing.T) {
  config := GetInstance()
  t.Log(config)
}
