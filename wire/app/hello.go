package app

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/gmidorii/api-design/wire/repository"
)

type Hello struct {
	repo repository.User
}

func NewHello(repo repository.User) Hello {
	return Hello{
		repo: repo,
	}
}

func (h Hello) GetMessage(id string) (string, error) {
	name, err := h.repo.Get(id)
	if err != nil {
		return "", errors.Wrap(err, "failed get name from repository")
	}

	return fmt.Sprintf("Hello %v", name), nil
}