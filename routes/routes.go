package routes

import (
	"stashcurrency.dev/controllers"
)

var Routes = []route{
	route{
		Path:       "/",
		HandleFunc: controllers.Index,
	},
	route{
		Path:       "/convert",
		HandleFunc: controllers.Convert,
	},
}
