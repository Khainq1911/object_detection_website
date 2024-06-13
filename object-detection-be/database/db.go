package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	User     string
	Port     int
	Password string
	Dbname   string
}

func (s *Sql) Connect() {
	sqlData := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s sslmode=disable", s.Host, s.User, s.Port, s.Password, s.Dbname)
	s.Db = sqlx.MustConnect("postgres", sqlData)

	if err := s.Db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Connect Succesfully!!")
	}
}

func (s *Sql) Close() {
	s.Db.Close()
}
