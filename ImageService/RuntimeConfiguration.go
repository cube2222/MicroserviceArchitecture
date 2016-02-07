package main

import "gopkg.in/mgo.v2"

type WorkerPoolSupervisorRegister struct {
	Address string
	Port    string
}

func registerWorkerPoolSupervisor(Supervisor WorkerPoolSupervisorRegister) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		logError(err)
		return
	}
	defer session.Close()
	c := session.DB("Configuration").C("WorkerPoolSupervisors")

	err = c.Insert(&Supervisor)
	if err != nil {
		logError(err)
		return
	}

}
