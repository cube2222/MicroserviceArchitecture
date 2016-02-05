package main

import (
	"fmt"
	"net/http"
)

func setUpRoutes() {
	http.HandleFunc("/newImage", newImage)
	http.HandleFunc("/getImage", getImage)
	http.HandleFunc("/finishedImage", finishedImage)
}

func newImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Do stuff...
	} else {
		fmt.Fprintln(w, "ERROR: Only POST accepted.")
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Do stuff...
	} else {
		fmt.Fprintln(w, "ERROR: Only GET accepted.")
	}
}

func finishedImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Do stuff...
	} else {
		fmt.Fprintln(w, "ERROR: Only POST accepted.")
	}
}
