package creational

import "fmt"

type Storage interface {
	Open()
}

type Postgres struct {
	ConnString string
}

func (s Postgres) Open() {
	fmt.Println("Postgres Connection")
}

type Redis struct {
	ConnString string
}

func (s Redis) Open() {
	fmt.Println("Redis Connection")
}


func NewStorage(target string) Storage {
	switch target {
	case "postgres":
		return &Postgres{ConnString: ""}
	case "redis":
		return &Redis{ConnString: ""}
	default:
		return nil
	}
}