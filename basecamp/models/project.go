package models

import (
	"sort"
	"time"
	"unicode"
)

// Project wraps the json response from bascamp projects into a go struct
type Project struct {
	ID             int       `json:"id" db:"id"`
	Status         string    `json:"status" db:"status"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Purpose        string    `json:"purpose" db:"purpose"`
	ClientsEnabled bool      `json:"clients_enabled" db:"clients_enabled"`
	BookmarkURL    string    `json:"bookmark_url" db:"bookmark_url"`
	URL            string    `json:"url" db:"url"`
	AppURL         string    `json:"app_url" db:"app_url"`
	Dock           []struct {
		ID       int    `json:"id" db:"id"`
		Title    string `json:"title" db:"title"`
		Name     string `json:"name" db:"name"`
		Enabled  bool   `json:"enabled" db:"enabled"`
		Position int    `json:"position" db:"position"`
		URL      string `json:"url" db:"url"`
		AppURL   string `json:"app_url" db:"app_url"`
	} `json:"dock" db:"dock"`
	Bookmarked bool `json:"bookmarked" db:"bookmarked"`
	Todos      []Todo
}

// Projectno returns the projectnumber or an empty string if not available
func (p *Project) Projectno() string {
	if len(p.Name) < 14 {
		return ""
	}
	nr := p.Name[:14]
	for i := 0; i < 4; i++ {
		r := rune(nr[i])
		if !unicode.IsUpper(r) {
			return ""
		}
	}
	return nr
}

// SortTodos sorts todos using the CreatedAt property
func (p *Project) SortTodos() {
	sort.Slice((*p).Todos, func(i, j int) bool {
		return ((*p).Todos)[i].CreatedAt.Before(((*p).Todos)[j].CreatedAt)
	})
}
