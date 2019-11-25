package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/AnotherCoolDude/todos/basecamp/models"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails"

	"github.com/AnotherCoolDude/todos/basecamp"
)

// Basecamp wraps the basecamp api in a struct for the frontend
type Basecamp struct {
	Projects []Project `json:"projects"`
	client   *basecamp.Client
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
}

// DefaultBasecamp returns an initialized basecamp instance
func DefaultBasecamp() (*Basecamp, error) {
	bs := Basecamp{}
	bs.Projects = []Project{}

	fmt.Println("loadin .env file...")
	p, err := os.Executable()
	if err != nil {
		return &bs, err
	}
	envPath := path.Join(path.Dir(p), ".env")
	fmt.Printf("expecting .env at path %s\n", envPath)
	err = godotenv.Load(envPath)
	if err != nil {
		return &bs, err
	}
	fmt.Println(os.Getenv("BASECAMP_CALLBACK"))

	fmt.Println("Initiating client")
	bs.client = basecamp.DefaultClient()
	return &bs, nil
}

// WailsInit initializes helpful runtime features
func (bs *Basecamp) WailsInit(runtime *wails.Runtime) error {
	bs.runtime = runtime
	bs.logger = bs.runtime.Log.New("Basecamp")
	bs.logger.Info("Logger initiated")
	bs.runtime.Window.SetTitle("basecamp -> Proad")
	return nil
}

// Login handles the login for basecamp
func (bs *Basecamp) Login() error {
	bs.logger.Info("Loggin in to Basecamp...")
	err := bs.client.Authenticate()
	if err != nil {
		bs.logger.Error(err.Error())
		return err
	}
	bs.logger.Info("Authentication successfull")
	return nil
}

// FetchProjects fetches the projects from basecamp
func (bs *Basecamp) FetchProjects(reset bool) ([]Project, error) {
	if len(bs.Projects) > 0 && !reset {
		bs.logger.Info("returning cached projects...")
		return bs.Projects, nil
	}
	bs.logger.Info("Fetching Projects...")
	basecamppp, err := bs.client.FetchProjects()
	if err != nil {
		bs.logger.Error(err.Error())
		return nil, err
	}
	bs.logger.Info("Fetching todos...")
	for i, p := range basecamppp {
		fmt.Printf("project #%d: %s\n", i, p.Name)
		err = bs.client.FetchTodos(&p)
		if err != nil {
			bs.logger.Error(err.Error())
			return nil, err
		}
		np := newProject(&p)
		bs.Projects = append(bs.Projects, *np)
	}
	bs.logger.Info("Fetching successfull")
	return bs.Projects, nil
}

// GetProjects returns projects fetched by the client
func (bs *Basecamp) GetProjects() []Project {
	for _, p := range bs.Projects {
		fmt.Printf("%d todos in Project %s\n", len(p.Todos), p.Jobnr)
	}
	return bs.Projects
}

// Project wraps the basecamp api in a struct for the frontend
type Project struct {
	Jobnr string `json:"nr"`
	Todos []Todo `json:"todos"`
}

// NewProject creates a new Project from a basecampclient project
func newProject(basecampProject *models.Project) *Project {
	tt := []Todo{}
	defaultDays := 10
	if days, err := strconv.Atoi(os.Getenv("DEFAULT_DAYS")); err == nil {
		defaultDays = days
	}
	for _, bct := range basecampProject.Todos {
		t := Todo{
			Title:           bct.Title,
			StartDate:       bct.CreatedAt,
			EndDate:         bct.CreatedAt.Add(time.Duration(defaultDays) * 24 * time.Hour),
			WorkAmountDone:  Workamount{Hours: "00", Minutes: "00"},
			WorkAmountTotal: Workamount{Hours: "10", Minutes: "00"},
			Projectnr:       bct.Projectno(),
		}
		tt = append(tt, t)
	}

	return &Project{
		Jobnr: basecampProject.Projectno(),
		Todos: tt,
	}
}

// Todo wraps the basecamp api in a struct for the frontend
type Todo struct {
	Title           string     `json:"title"`
	StartDate       time.Time  `json:"startDate"`
	EndDate         time.Time  `json:"endDate"`
	WorkAmountTotal Workamount `json:"workAmountTotal"`
	WorkAmountDone  Workamount `json:"workAmountDone"`
	Projectnr       string     `json:"projectnr"`
	Assignee        Assignee   `json:"assignee"`
}

// Assignee wraps the assignee of the todo into a struct
type Assignee struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Urno      int    `json:"urno"`
}

// Workamount wraps the basecamp api in a struct for the frontend
type Workamount struct {
	Hours   string `json:"HH"`
	Minutes string `json:"mm"`
}
