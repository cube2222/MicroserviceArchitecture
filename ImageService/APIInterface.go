package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func setUpRoutes(exitChan chan bool, databaseAddress string) {
	http.HandleFunc("/newImage", newImage)
	http.HandleFunc("/getImage", getImage)
	http.HandleFunc("/finishedImage", finishedImage)
	http.Handle("/registerWorkerSupervisor", registerWorkerSupervisorHandler{databaseAddress})
	http.Handle("/exit", exitHandler{exitChan})
	http.ListenAndServe(":3000", nil)
}

type exitHandler struct {
	exitChan chan bool
}

func (i exitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.exitChan <- true
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
		//TODO: Register upcoming work on file to do with use of parsedQuery["id"][0]
	} else {
		fmt.Fprintln(w, "ERROR: Only POST accepted.")
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
		file, err := getFile(parsedQuery["id"][0], "finished")
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		io.Copy(w, file)
	} else {
		fmt.Fprintln(w, "ERROR: Only GET accepted.")
	}
}

func finishedImage(w http.ResponseWriter, r *http.Request) {
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
		err = saveFile(parsedQuery["id"][0], r.Body, "finished")
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		//TODO: Register finished work on file to do with use of parsedQuery["id"][0]
	} else {
		fmt.Fprintln(w, "ERROR: Only POST accepted.")
	}
}

type registerWorkerSupervisorHandler struct {
	databaseAddress string
}

func (i registerWorkerSupervisorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
		newSupervisor := WorkerPoolSupervisorRegister{}
		newSupervisor.Address = parsedQuery["address"][0]
		newSupervisor.Port = parsedQuery["port"][0]
		registerWorkerPoolSupervisor(newSupervisor, i.databaseAddress)
	} else {
		fmt.Fprintln(w, "ERROR: Only GET accepted.")
	}
}
