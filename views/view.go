package views

import (
	"html/template"
)

// View type contains generic template
type View struct {
	Template *template.Template
}

// NewView generates new View with input file/s as basis
// and appends to it other neccessary layout files.
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}
