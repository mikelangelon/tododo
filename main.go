package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type todo struct {
	Task string
}

func main() {
	a := app.New()
	w := a.NewWindow("Awesome TODO list")
	w.Resize(fyne.NewSize(400, 500))

	s := store{filename: "todos.json"}
	myTodos, err := s.GetToDos()
	if err != nil {
		handleError(err)
	}

	task := widget.NewEntry()
	task.SetPlaceHolder("Write task...")

	submit := widget.NewButton("Submit", func() {
		if task.Text != "" {
			t := todo{
				Task: task.Text,
			}
			myTodos = append(myTodos, t)
			if err := s.SaveToDos(myTodos); err != nil {
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
