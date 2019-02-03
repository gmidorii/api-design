package app

import (
	"fmt"

	"github.com/gmidorii/api-design/wire/domain"
	"github.com/pkg/errors"
)

type Hello struct {
	user domain.UserRepository
}

func NewHello(user domain.UserRepository) Hello {
	return Hello{
		user: user,
	}
}

func (h Hello) GetMessage(id string) (string, error) {
	name, err := h.user.Get(id)
	if err != nil {
		return "", errors.Wrap(err, "failed get name from repository")
	}

	return fmt.Sprintf("Hello %v", name), nil
}
