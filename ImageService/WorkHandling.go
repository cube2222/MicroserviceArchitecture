package main

import "gopkg.in/mgo.v2"

type Task struct {
	id          string
	bInProgress bool
	bFinished   bool
}

func registerWorkToDo(id string, databaseAddress string) {
	session, err := mgo.Dial(databaseAddress)
	if err != nil {
		logError(err)
		return
	}
	defer session.Close()
	c := session.DB("Work").C("inWork")
	err = c.Insert(Task{id, false, false})
	if err != nil {
		logError(err)
		return
	}
}
