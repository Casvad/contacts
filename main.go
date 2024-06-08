package main

import (
	"contacts/jobs"
	"contacts/routers"
	"fmt"
)

type Server interface {
	GetRouter() routers.Router
	GetOnStartUpJobs() []jobs.OnStartUpJob
}

type server struct {
	router routers.Router
	jobs   []jobs.OnStartUpJob
}

func (s *server) GetRouter() routers.Router {
	return s.router
}

func (s *server) GetOnStartUpJobs() []jobs.OnStartUpJob {
	return s.jobs
}

func provideServer(router routers.Router, jobs []jobs.OnStartUpJob) Server {
	return &server{router, jobs}
}

func main() {

	s := initializeDependencies()

	for _, job := range s.GetOnStartUpJobs() {
		job.Execute()
	}

	err := s.GetRouter().CreateRouter().Run()

	if err != nil {
		fmt.Printf("Error stating server: %v", err)
		panic(err)
	}
}
