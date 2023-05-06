package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello fyne!")
	// Campo de entrada input
	input := widget.NewEntry()
	input.SetPlaceHolder("Ingrese su nombre")
	// Mostramos el input
	hello := widget.NewLabel("")
	// Cuando hagamos clic en boton mostramos
	send := widget.NewButton("Enviar", func() {
		hello.SetText("Hola " + input.Text)
	})

	content := container.NewVBox(input, send, hello)

	// myWindow.SetContent(widget.NewLabel("Hola mundo"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
