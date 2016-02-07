package main

import "gopkg.in/mgo.v2"

type WorkerPoolSupervisorRegister struct {
	Address string
	Port    string
}

func registerWorkerPoolSupervisor(Supervisor WorkerPoolSupervisorRegister) {
	session, err := mgo.Dial("")
	if err != nil {
		logError(err)
		return
	}
	defer session.Close()
	c := session.DB("configuration").C("WorkerPoolSupervisors")

	err = c.Insert(&Supervisor)
	if err != nil {
		logError(err)
		return
	}

}
