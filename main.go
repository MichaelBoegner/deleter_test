package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	CampaignLister(token)

	// myApp := app.New()
	// myWindow := myApp.NewWindow("Input Example")

	// label := widget.NewLabel("Api Key:")
	// entry := widget.NewEntry()
	// button := widget.NewButton("Submit", func() {
	// 	input := entry.Text
	// 	label.SetText("You entered: " + input)
	// })

	// content := container.NewVBox(label, entry, button)
	// myWindow.SetContent(content)
	// myWindow.ShowAndRun()
}
