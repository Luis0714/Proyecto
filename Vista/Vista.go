package Vista

import (
	"proyecto/Controller"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Run(){
		APP := app.New()
		WINDOW := APP.NewWindow("VENTANA")

		LabelGramatica := widget.NewLabel("")
		LabelGramatica.Alignment = fyne.TextAlignLeading
		LabelGramatica.TextStyle = fyne.TextStyle{Bold: true}

		LabelPrimeros:= widget.NewLabel("")
		LabelPrimeros.Alignment = fyne.TextAlignLeading
		LabelPrimeros.TextStyle = fyne.TextStyle{Bold: true}

		PrimerosButton := widget.NewButtonWithIcon("ESCOGER GRAMATICA", theme.DocumentIcon(), func() {
			primeros := Controller.ObtenerCadenas()
			gramatica := Controller.MostrarGramatica()
			LabelGramatica.SetText(gramatica)
			LabelPrimeros.SetText(primeros)
			LabelPrimeros.Size()
		})
		PrimerosButton.Cursor() 
		PrimerosButton.Size()
	
		content := container.NewVBox(
			PrimerosButton,
			LabelGramatica,
			LabelPrimeros,
		)

		WINDOW.SetContent(content)
		WINDOW.Resize(fyne.NewSize(800, 600)) // Establece un tamaño predeterminado para la ventana
		WINDOW.Show()
		APP.Run()

}