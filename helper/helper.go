package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AnotherCoolDude/todos/proad/models"
)

const (
	proadDateTimeFormat = "2006-01-02T15:04:05"
)

// Query is a convenience for map[string]string
type Query map[string]string

// ResponseBytes returns the json body of response as bytes
func ResponseBytes(response *http.Response) ([]byte, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return bytes, nil
}

// ProadTodo returns an todo that can be used to create a new todo
func ProadTodo(title string, startDate, endDate time.Time, hoursPlanned, hoursLeft float64, managerUrno, projectUrno, responsibleUrno int) models.PostTodo {
	return models.PostTodo{
		Shortinfo:       title,
		ProjectUrno:     projectUrno,
		ManagerUrno:     managerUrno,
		ResponsibleUrno: responsibleUrno,
		FromDatetime:    startDate.Format(proadDateTimeFormat),
		UntilDatetime:   endDate.Format(proadDateTimeFormat),
		HoursPlanned:    hoursPlanned,
		HoursLeft:       hoursLeft,
	}
}

// PrettyPrintBytes prints out bytes (e.g. from a response) in a readable way
func PrettyPrintBytes(bb []byte) error {
	var jsonPretty bytes.Buffer
	err := json.Indent(&jsonPretty, bb, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonPretty.Bytes()))
	return nil
}
