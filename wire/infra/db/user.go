package db

import (
	"github.com/gmidorii/api-design/wire/domain"
)

type User struct {
}

func NewUser() domain.UserRepository {
	return &User{}
}

func (u *User) Get(id string) (domain.User, error) {
	// TODO: fetch db
	return domain.User{}, nil
}
