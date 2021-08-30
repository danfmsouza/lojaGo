package main

import (
	"net/http"
	"std/github.com/danfmsouza/lojaGo/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
