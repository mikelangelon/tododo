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
	var selectedItem int
	s := store{filename: "todos.json"}
	todos, err := s.GetToDos()
	if err != nil {
		handleError(err)
	}

	a := app.New()
	w := a.NewWindow("Awesome TODO list")
	w.Resize(fyne.NewSize(400, 500))

	task := widget.NewEntry()
	task.SetPlaceHolder("Write task...")

	label := widget.NewLabel("Selected item")
	label.Hidden = true

	current := widget.NewEntry()
	current.SetPlaceHolder("Write task...")
	current.Hidden = true

	list := widget.NewList(
		func() int { return len(todos) },
		func() fyne.CanvasObject {
			return widget.NewLabel("ToDo item")
		},
		func(l widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(todos[l].Task)
		},
	)

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
			list.Refresh()
		}
	})

	update := widget.NewButton("Update", func() {
		todos[selectedItem].Task = current.Text
		if err := s.SaveToDos(todos); err != nil {
			handleError(err)
		}

		current.Text = ""
		current.Refresh()
		list.Refresh()
	})
	update.Hidden = true

	form := container.NewVBox(task, submit, label, current, update)
	split := container.NewHSplit(list, form)
	split.SetOffset(0.4)

	list.OnSelected = func(id widget.ListItemID) {
		selectedItem = id
		label.Hidden = false
		current.Hidden = false
		update.Hidden = false
		current.Text = todos[id].Task

		label.Refresh()
		current.Refresh()
		form.Refresh()

	}
	w.SetContent(split)

	w.ShowAndRun()
}

func handleError(err error) {
	panic(err)
}
