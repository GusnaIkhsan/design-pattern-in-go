package main

import "fmt"

type DatabaseConnection struct{}

func (dc *DatabaseConnection) Open() {
	fmt.Println("Opening database connection...")
}

func (dc *DatabaseConnection) SetConfiguration() {
	fmt.Println("Setting configuration...")
}

type DatabaseSession struct{}

func (ds *DatabaseSession) Init() {
	fmt.Println("Initializing session...")
}

type DatabaseFacade struct {
	connection *DatabaseConnection
	session    *DatabaseSession
}

func (df *DatabaseFacade) Start() {
	df.connection.Open()
	df.connection.SetConfiguration()
	df.session.Init()
	fmt.Println("Database started!")
}

func main() {
	dbFacade := &DatabaseFacade{
		connection: &DatabaseConnection{},
		session:    &DatabaseSession{},
	}
	dbFacade.Start()
}
