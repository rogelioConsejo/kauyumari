package login

import (
	"github.com/rogelioConsejo/kauyumari/user"
	"testing"
)

func TestNewUserRegistry(t *testing.T) {
	var r UserRegistry = NewUserRegistry()
	if r == nil {
		t.Fatal("user registry is nil")
	}
}

func TestUserRegistry_CreateUser(t *testing.T) {
	registry := NewUserRegistry()
	var userName user.Name = "testelio"
	var email user.Email = "testelio@emailprovider.com"
	u, err := user.New(userName, email)
	if err != nil {
		t.Fatal("unexpected error when creating user entity: ", err)
	}
	t.Run("it should create a user", func(t *testing.T) {
		_ = registry.CreateUser(u)
		if !registry.UserExists(u) {
			t.Fatal("user was not created")
		}
	})
	t.Run("it should return an error if the user name already exists", func(t *testing.T) {
		err = registry.CreateUser(u)
		if err == nil {
			t.Fatal("expected error when creating repeated user")
		}
	})
}

func TestUserRegistry_UserExists(t *testing.T) {
	registry := NewUserRegistry()
	var userName user.Name = "testelio"
	var email user.Email = "testelio@email.com"
	u, err := user.New(userName, email)
	if err != nil {
		t.Fatal("unexpected error when creating user entity: ", err)
	}
	if registry.UserExists(u) {
		t.Fatal("user should not exist yet")
	}
	_ = registry.CreateUser(u)
	if !registry.UserExists(u) {
		t.Fatal("user was not created")
	}
}
