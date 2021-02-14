package views

import (
	"html/template"
)

// View type contains generic template
type View struct {
	Template *template.Template
	Layout   string
}

// NewView generates new View with input file/s as basis
// and appends to it other neccessary layout files.
func NewView(layout string, files ...string) *View {
	files = append(
		files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml",
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout: layout,
	}
}
