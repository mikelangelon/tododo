package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type todo struct {
	Task      string
	CreatedAt time.Time
	DoneAt    *time.Time
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
				Task:      task.Text,
				CreatedAt: time.Now(),
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

	delete := colorButton(
		"Delete",
		color.NRGBA{R: 150, G: 0, B: 0, A: 255},
		func() {
			li := append(todos[:selectedItem], todos[selectedItem+1:]...)
			if err := s.SaveToDos(li); err != nil {
				handleError(err)
			}
			todos, _ = s.GetToDos()
			list.Refresh()
		})
	delete.Hidden = true

	done := colorButton(
		"Done",
		color.NRGBA{R: 0, G: 150, B: 0, A: 255},
		func() {
			// Copy of update except adding DoneAt
			now := time.Now()
			todos[selectedItem].DoneAt = &now
			if err := s.SaveToDos(todos); err != nil {
				handleError(err)
			}
			todos, _ = s.GetToDos()
			list.Refresh()
		})
	done.Hidden = true
	form1 := container.NewVBox(task, submit)
	form2 := container.NewVBox(label, current, update, delete, done)
	separator := widget.NewSeparator()
	separator.Resize(fyne.NewSize(separator.MinSize().Width, 300))
	separator.Refresh()
	form := container.NewVBox(form1, layout.NewSpacer(), form2, layout.NewSpacer())

	split := container.NewHSplit(list, form)
	split.SetOffset(0.4)

	list.OnSelected = func(id widget.ListItemID) {
		selectedItem = id
		label.Hidden = false
		current.Hidden = false
		update.Hidden = false
		delete.Hidden = false
		done.Hidden = false
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

func colorButton(label string, color color.NRGBA, tapped func()) *fyne.Container {
	btn := widget.NewButton(label, tapped)
	rectColor := canvas.NewRectangle(color)
	container := container.New(
		layout.NewMaxLayout(),
		rectColor,
		btn,
	)

	return container
}
