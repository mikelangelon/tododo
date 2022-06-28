package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
)

type todo struct {
	Task string
}

func main() {
	a := app.New()
	w := a.NewWindow("Awesome TODO list")
	w.Resize(fyne.NewSize(400, 500))

	res, err := ioutil.ReadFile("todos.json")
	if err != nil {
		handleError(err)
	}

	var myTodos []todo
	json.Unmarshal(res, &myTodos)

	task := widget.NewEntry()
	task.SetPlaceHolder("Write task...")

	submit := widget.NewButton("Submit", func() {
		if task.Text != "" {
			t := todo{
				Task: task.Text,
			}
			myTodos = append(myTodos, t)
			b, err := json.MarshalIndent(myTodos, "", " ")
			if err != nil {
				handleError(err)
			}
			err = os.WriteFile("todos.json", b, 0644)
			if err != nil {
				handleError(err)
			}

			task.Text = ""
			task.Refresh()
		}
	})
	w.SetContent(container.NewVBox(task, submit))

	w.ShowAndRun()
}

func handleError(err error) {
	panic(err)
}
