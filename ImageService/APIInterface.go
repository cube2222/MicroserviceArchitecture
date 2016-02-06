package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func setUpRoutes() {
	http.HandleFunc("/newImage", newImage)
	http.HandleFunc("/getImage", getImage)
	http.HandleFunc("/finishedImage", finishedImage)
	http.HandleFunc("/registerWorkerSupervisor", registerWorkerSupervisor)
}

func newImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		parsedUrl, err := url.Parse(r.URL.String())
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		parsedQuery, err := url.ParseQuery(parsedUrl.RawQuery)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		err = saveFile(parsedQuery["id"][0], r.Body, "working")
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
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

func registerWorkerSupervisor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Do stuff...
	} else {
		fmt.Fprintln(w, "ERROR: Only GET accepted.")
	}
}
