package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonDoc := `{
		"environment": "asdf",
		"hostName": ""
     }`
	conf := &ConfigWithPointers{}
	json.Unmarshal([]byte(jsonDoc), conf)
	fmt.Println("-------", conf)
}
