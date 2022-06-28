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
	s := store{filename: "todos.json"}
	todos, err := s.GetToDos()
	if err != nil {
		handleError(err)
	}

	a := app.New()
	w := a.NewWindow("Awesome TODO list")
	w.Resize(fyne.NewSize(400, 500))

	list := widget.NewList(
		func() int { return len(todos) },
		func() fyne.CanvasObject {
			return widget.NewLabel("ToDo item")
		},
		func(l widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(todos[l].Task)
		},
	)

	task := widget.NewEntry()
	task.SetPlaceHolder("Write task...")

	submit := widget.NewButton("Submit", func() {
		if task.Text != "" {
			t := todo{
				Task: task.Text,
			}
			todos = append(todos, t)
			if err := s.SaveToDos(todos); err != nil {
				handleError(err)
			}

			task.Text = ""
			task.Refresh()
		}
	})
	split := container.NewHSplit(
		list,
		container.NewVBox(task, submit),
	)
	split.SetOffset(0.4)
	w.SetContent(split)

	w.ShowAndRun()
}

func handleError(err error) {
	panic(err)
}
