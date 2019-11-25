package main

import (
	"fmt"
	"github.com/AnotherCoolDude/todos/helper"
	"github.com/AnotherCoolDude/todos/proad"
	"github.com/AnotherCoolDude/todos/proad/models"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
	"reflect"
	"time"
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
	err := decode(todoInterface, &todo)
	if err != nil {
		pr.logger.Errorf("error decoding todoInterface %v: %s", todoInterface, err)
		return err
	}
	if todo.Assignee.Urno == 0 {
		pr.logger.Errorf("Zuständigen für %s auswählen", todo.Title)
		return fmt.Errorf("error processing todo %s: No assignee found", todo.Title)
	}
	var project models.Project
	err = pr.client.FetchProject(todo.Projectnr, &project)
	if err != nil {
		pr.logger.Errorf("could not find project %s: %s", todo.Projectnr, err)
		return err
	}
	ptodo := helper.ProadTodo(todo.Title, todo.StartDate, todo.EndDate, todo.WorkAmountTotal, todo.WorkAmountTotal-todo.WorkAmountDone, pr.client.ManagerUrno, project.Urno, todo.Assignee.Urno)
	fmt.Printf("%+v\n", ptodo)
	return pr.client.PostTodo(ptodo)
}

// need custom decoder for mapstructure to be able to decode time.Time
func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}
	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
