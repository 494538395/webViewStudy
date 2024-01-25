package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Dynamic Component Example")

	messageEntry := widget.NewMultiLineEntry()
	messageEntry.Text = "Message Entry"

	msgHeaderEntry := widget.NewMultiLineEntry()
	msgHeaderEntry.Text = "Header Entry"

	msgParamEntry := widget.NewMultiLineEntry()
	msgParamEntry.Text = "Param Entry"

	currentComponent := messageEntry

	content := container.NewVBox(
		currentComponent,
	)

	msgJsonEditButton := widget.NewButton("Show Message Entry", func() {
		if currentComponent != messageEntry {
			content.Hide(currentComponent)
			currentComponent = messageEntry
			content.Show(currentComponent)
		}
	})

	msgHeaderEditButton := widget.NewButton("Show Header Entry", func() {
		if currentComponent != msgHeaderEntry {
			content.Hide(currentComponent)
			currentComponent = msgHeaderEntry
			content.Show(currentComponent)
		}
	})

	msgParamEditButton := widget.NewButton("Show Param Entry", func() {
		if currentComponent != msgParamEntry {
			content.Hide(currentComponent)
			currentComponent = msgParamEntry
			content.Show(currentComponent)
		}
	})

	genMsgLine := container.NewGridWithColumns(3,
		msgJsonEditButton,
		msgHeaderEditButton,
		msgParamEditButton,
	)

	content.Append(genMsgLine)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
