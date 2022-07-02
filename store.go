package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type store struct {
	filename string
}

func (s store) GetToDos() ([]todo, error) {
	todos, err := s.getToDos()
	if err != nil {
		return nil, err
	}
	var todosNotDone []todo
	for _, v := range todos {
		if v.DoneAt == nil {
			todosNotDone = append(todosNotDone, v)
		}
	}
	return todosNotDone, nil
}

func (s store) SaveToDos(todos []todo) error {
	b, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(s.filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s store) getToDos() ([]todo, error) {
	res, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	var todos []todo
	err = json.Unmarshal(res, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
