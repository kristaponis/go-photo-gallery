package views

import (
	"html/template"
	"path/filepath"
)

// View type contains generic template
type View struct {
	Template *template.Template
	Layout   string
}

// NewView generates new View with generic template, input file as basis
// and appends to it other neccessary layout files.
func NewView(layout string, files ...string) *View {
	files = append(files, globFiles("views/layouts/", ".gohtml")...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// globFiles function returns a slice of strings,
// whitch are taken from a given directory with the given
// file extention name
func globFiles(path string, ext string) []string {
	globFiles, err := filepath.Glob(path + "*" + ext)
	if err != nil {
		panic(err)
	}
	return globFiles
}