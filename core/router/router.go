package router

import (
	"net/http"
	"stashcurrency.dev/routes"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Intitialize() {
	for _, route := range routes.Routes {
		http.HandleFunc(route.Path, route.HandleFunc)
	}
}
