package main

import "gopkg.in/mgo.v2"

type Task struct {
	workId      string
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
	c := session.DB("Work").C("InWork")
	newTask := Task{}
	newTask.workId = id
	newTask.bInProgress = "false"
	newTask.bFinished = "false"
	err = c.Insert(&newTask)
	//TODO: For some reason it's just not working...
	if err != nil {
		logError(err)
		return
	}
}
