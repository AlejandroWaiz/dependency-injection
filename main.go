package main

import (
	"github.com/AlejandroWaiz/dependency-injection/database"
	"github.com/AlejandroWaiz/dependency-injection/service"
)

type Program struct {
	dataBase database.IDatabase
	service  service.IService
}

func (p Program) execute() {

	user := p.dataBase.GetUserName()
	p.service.SayHi(user)

}

func main() {

	db := database.Database{}
	serv := service.Service{}

	p := Program{
		dataBase: db,
		service:  serv,
	}

	p.execute()

}
