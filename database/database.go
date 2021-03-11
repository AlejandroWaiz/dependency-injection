package database

type IDatabase interface {
	GetUserName() string
}

type Database struct{}

func (db Database) GetUserName() string {

	return "Makunouchi Ippo"

}
