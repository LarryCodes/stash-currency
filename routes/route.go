package routes

import (
	"net/http"
)

type route struct {
	Path       string
	HandleFunc http.HandlerFunc
}
