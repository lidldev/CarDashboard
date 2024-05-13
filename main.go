package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeUI(button *widget.Button) *fyne.Container {
	return container.New(layout.NewGridLayout(0),
		button,
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("Dash")

	button := widget.NewButton("Press Me ;)", func() {
		fmt.Println("Hello world")
	})

	w.SetContent(makeUI(button))

	a.Run()
	w.Show()
}
