package Vista

import (
	"proyecto/Controller"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run(){
	APP := app.New()
	WINDOW := APP.NewWindow("VENTANA DE SALIDA")
	label := widget.NewLabel("")
	PrimerosButton := widget.NewButton("Obtener primeros", func() {
		primeros := Controller.ObtenerCadenas()
		label.SetText(primeros)
	})
	content := container.NewVBox(PrimerosButton,label)

	WINDOW.SetContent(content)
	WINDOW.Show()
	APP.Run()
}