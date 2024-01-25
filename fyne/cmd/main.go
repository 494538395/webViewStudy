package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Demo")

	label := widget.NewLabel("Welcome to Fyne!")
	button := widget.NewButton("Click me!", func() {
		label.SetText("Button Clicked!")
	})

	content := container.NewVBox(
		label,
		button,
	)

	myWindow.SetContent(content)
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
