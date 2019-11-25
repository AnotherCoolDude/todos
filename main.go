package main

import (
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	basecamp, err := DefaultBasecamp()
	if err != nil {
		log.Fatal(err)
	}

	proad, err := DefaultProad()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "todos",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(basecamp)
	app.Bind(proad)
	app.Run()
}
