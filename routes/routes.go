package routes

import (
	"stashcurrency.dev/controllers"
)

var Routes = []route{
	route{
		Path:       "/",
		HandleFunc: controllers.Index,
	},
}
