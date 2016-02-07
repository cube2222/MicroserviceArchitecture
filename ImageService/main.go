package main

func main() {
	exitChan := make(chan bool)
	go setUpRoutes(exitChan, "localhost:27017")
	<-exitChan
}
