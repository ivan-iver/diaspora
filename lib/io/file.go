// Package io encapsulate any file system access
package io

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// File models an entity file
type File struct {
	Path string
	Name string
	file *os.File
}

// Create method storage a new File into file system.
func (f *File) Create(t *Template) (err error) {
	if _, err = os.Stat(f.Path); err != nil {
		if err = os.Mkdir(f.Path, 0755); err != nil {
			return
		}
	}
	if f.file, err = os.Create(f.AbsoluteName()); err != nil {
		return
	}
	defer f.file.Close()
	err = t.Data.Execute(f.file, t.ID)
	return
}

// Read content file
func (f *File) Read() (data string, err error) {
	var content []byte
	filename := f.AbsoluteName()
	if content, err = ioutil.ReadFile(filename); err != nil {
		return data, err
	}
	return string(content), err
}

// AbsoluteName method returns absolute file name
func (f *File) AbsoluteName() string {
	if len(f.Path) == 0 {
		f.Path, _ = f.Pwd()
	}
	if strings.HasPrefix(f.Path, "/") {
		return path.Join(f.Path, f.Name)
	}
	var p string
	p, _ = f.Pwd()
	return path.Join(p, f.Path, f.Name)
}

// Pwd method gets current directory
func (f *File) Pwd() (path string, err error) {
	path, err = os.Getwd()
	return
}

// Exists method indicates whether file exists
func (f *File) Exists() (ok bool) {
	filename := f.AbsoluteName()
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

// NotExists method indicates whether file not exists
func (f *File) NotExists() (ok bool) {
	filename := f.AbsoluteName()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return true // Not exists
	}
	return false // File exists
}

// Lines return each line from file as []string.
func (f *File) Lines() (lines []string, err error) {
	var data string
	if data, err = f.Read(); err != nil {
		return lines, err
	}
	lines = strings.Split(data, "\n")
	return lines, err
}
