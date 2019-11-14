package main

import (
	"fmt"
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

//ErrorOrSuccess is used to receive potential errors from the frontend
func ErrorOrSuccess(success bool) (string, error) {
	if success {
		return "I was successful", nil
	}
	return "", fmt.Errorf("i am an error")
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	myTodoList, err := NewTodos()
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

	app.Bind(myTodoList)

	app.Bind(ErrorOrSuccess)
	app.Run()
}
