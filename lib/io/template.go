package io

import (
	"fmt"
	"github.com/ivan-iver/diaspora/lib/time"
	"path/filepath"
	"text/template"
)

var migrationTpl = template.Must(template.New("migration").Parse(`
-- @Up
-- SQL in section 'Up' is executed when this migration is applied


-- @Down
-- SQL section 'Down' is executed when this migration is rolled back

`))

const sep = string(filepath.Separator)

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
	t.File.Name = fmt.Sprintf("%s_%s.sql", t.ID, t.Name)
	path, err := t.File.Pwd()

	t.File.Path = fmt.Sprintf("%s%s%s%s", path, sep, t.Path, sep)
	if err != nil {
		return
	}
	err = t.File.Create(t)
	return
}
