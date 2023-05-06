package main

import (
	"fmt"
	"time"

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
		name := input.Text
		if name != "" {
			hourOfDay := time.Now().Hour()
			greeting := getGreeting(hourOfDay)
			hello.SetText(fmt.Sprintf("%s, %s!", greeting, name))
		}
	})

	content := container.NewVBox(input, send, hello)

	// myWindow.SetContent(widget.NewLabel("Hola mundo"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}

func getGreeting(hour int) string {
	switch {
	case hour >= 6 && hour < 12:
		return "Buenos días"
	case hour >= 12 && hour < 18:
		return "Buenas tardes"
	default:
		return "Buenas noches"
	}
}
