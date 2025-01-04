package login

import (
	"errors"
	"github.com/rogelioConsejo/kauyumari/user"
)

func NewUserRegistry() UserRegistry {
	return userRegistry{
		users: make(map[user.Name]user.User),
	}
}

type UserRegistry interface {
	CreateUser(u user.User) error
	UserExists(u user.User) bool
}

type userRegistry struct {
	users map[user.Name]user.User
}

func (ur userRegistry) CreateUser(u user.User) error {
	if _, ok := ur.users[u.Name()]; ok {
		return ErrUserAlreadyExists
	}
	ur.users[u.Name()] = u
	return nil
}

func (ur userRegistry) UserExists(u user.User) bool {
	_, ok := ur.users[u.Name()]
	return ok
}

var ErrUserAlreadyExists = errors.New("user already exists")
