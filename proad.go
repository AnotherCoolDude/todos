package main

import (
	"github.com/AnotherCoolDude/todos/helper"
	"github.com/AnotherCoolDude/todos/proad"
	"github.com/AnotherCoolDude/todos/proad/models"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

// Proad wraps the proad client into a struct
type Proad struct {
	client  *proad.Client
	runtime *wails.Runtime
	logger  *wails.CustomLogger
}

// DefaultProad returns a initialized Proad instance
func DefaultProad() (*Proad, error) {
	pr := Proad{}
	pr.client = proad.DefaultClient()
	return &pr, nil
}

// WailsInit initializes helpful runtime features
func (pr *Proad) WailsInit(runtime *wails.Runtime) error {
	pr.runtime = runtime
	pr.logger = pr.runtime.Log.New("Proad")
	pr.logger.Info("Logger initiated")
	return nil
}

// GetEmployees returns the employees from proad
func (pr *Proad) GetEmployees() interface{} {
	type Employee struct {
		Name string `json:"name"`
		Urno int    `json:"urno"`
	}
	ee := []Employee{}
	for k, v := range pr.client.Employees {
		ee = append(ee, Employee{Name: k, Urno: v})
	}
	return ee
}

// CreateTodo creates a new proad todo
func (pr *Proad) CreateTodo(todoInterface map[string]interface{}) error {
	var todo Todo
	mapstructure.Decode(todoInterface, &todo)
	var project models.Project
	err := pr.client.FetchProject(todo.Projectnr, &project)
	if err != nil {
		pr.logger.Errorf("could not find project %s: %s", todo.Projectnr, err)
		return err
	}
	ptodo := helper.ProadTodo(todo.Title, todo.StartDate, todo.EndDate, pr.client.ManagerUrno, project.Urno, todo.Assignee.Urno)
	return pr.client.PostTodo(ptodo)
}
