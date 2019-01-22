package db

import (
	"github.com/gmidorii/api-design/wire/repository"
)

type User struct {

}

func NewUser() repository.User {
	return &User{}
}

func (u *User) Get(id string) (string, error) {
	// TODO: fetch db
	return "hogehoge", nil
}