package main

import (
	"net/http"
)

func quellenHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "quellen", "base")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", "base")
}
