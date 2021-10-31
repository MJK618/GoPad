package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"strconv"
)

var fileCount = 1

func main() {
	fmt.Println("Welcome to GoPad")
	a := app.New()
	w := a.NewWindow("GoPad")
	w.Resize(fyne.NewSize(500, 500))

	firstLabel := widget.NewLabel("Welcome to GoPad!")
	content := container.NewVBox(
		container.NewHBox(
			firstLabel,
		))
	content.Add(widget.NewButton("Add", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(fileCount)))
		fileCount += 1
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Go Enter Your Text...")

	saveButton := widget.NewButton("Save File", func() {
		saveFileDialogBox := dialog.NewFileSave(
			func(closer fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				closer.Write(textData)
			}, w)
		saveFileDialogBox.SetFileName("New File" + strconv.Itoa(fileCount-1) + ".txt")
		saveFileDialogBox.Show()
	})

	openButton := widget.NewButton("Open File", func() {
		openFileDialogBox := dialog.NewFileOpen(
			func(readCloser fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(readCloser)

				output := 	fyne.NewStaticResource("New File ", readData)
				showReadData := widget.NewMultiLineEntry()
				showReadData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(string(output.StaticName))
				w.SetContent(container.NewScroll(showReadData))
				w.Resize(fyne.NewSize(400,400))
				w.Show()

			},w)
		openFileDialogBox.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialogBox.Show()
	})

	w.SetContent(container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveButton,
			openButton,
		),
	))

	w.ShowAndRun()
}
