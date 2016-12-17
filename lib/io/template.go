package io

import (
	"fmt"
	"github.com/ivan-iver/diaspora/lib/time"
	"text/template"
)

var migrationTpl = template.Must(template.New("migration").Parse(`
-- @Up
-- SQL in section 'Up' is executed when this migration is applied


-- @Down
-- SQL section 'Down' is executed when this migration is rolled back

`))

// Template models a new file with basic template data
type Template struct {
	ID   string
	Name string
	Path string
	Data *template.Template
	*File
}

// NewTemplate returns an empty template with basic tags.
func NewTemplate() (tmp *Template, err error) {
	var serie = time.NewSerie()
	tmp = &Template{
		ID:   serie.Formated,
		Data: migrationTpl,
		File: &File{},
	}
	return
}

// Save persist the information template
func (t *Template) Save() (err error) {
	t.File.Name = fmt.Sprintf("%s_%s", t.ID, t.Name)
	t.File.Extension = ".sql"
	path, err := t.File.Pwd()
	t.File.Path = path
	if err != nil {
		return
	}
	err = t.File.Create(t)
	return
}
