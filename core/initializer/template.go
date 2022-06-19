package initializer

import (
	"html/template"
)

// Initialize templates
var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./frontend/*.gohtml"))
}
