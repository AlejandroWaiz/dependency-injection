package service

import (
	"log"

	"github.com/AlejandroWaiz/dependency-injection/database"
)

type IService interface {
	SayHi(userName string)
}

type Service struct {
	Database database.IDatabase
}

func (service Service) SayHi(userName string) {

	log.Printf("Hello %s !, nice to meet you", userName)

}
