package controllers

import (
	"fmt"
	"net/http"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Converter!")
}
