package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

func FromRequest(r *http.Request) (Task, error) {
	if r.Body == nil {
		return Task{}, errors.New("Empty body is not allowed")
	}

	var newTask Task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("An error has occurred getting request body", err)
	}

	err = json.Unmarshal(reqBody, &newTask)
	if err != nil {
		return Task{}, errors.New("Error parsing the JSON encoded data")
	}

	return newTask, nil
}
