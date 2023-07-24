package mod04

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func GenerateRandomMessage() string {
	var err error
	var out bytes.Buffer
	cmd := exec.Command("curl", "api.quotable.io/random")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return err.Error()
	}
	quote := Quote{}
	err = json.Unmarshal(out.Bytes(), &quote)
	if err != nil {
		return err.Error()
	}
	return quote.Content + " - " + quote.Author
}
