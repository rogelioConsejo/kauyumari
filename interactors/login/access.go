package login

import (
	"errors"
	"github.com/rogelioConsejo/kauyumari/entities/user"
)

func NewAccess(dam AuthenticationMethod) Access {
	return access{
		defaultAuthenticationMethod: dam,
	}
}

type Access interface {
	PrepareAuthentication(user.User) error
	PerformAuthentication(user.User, Credential) (bool, error)
}

type access struct {
	defaultAuthenticationMethod AuthenticationMethod
}

func (a access) PrepareAuthentication(u user.User) error {
	err := a.defaultAuthenticationMethod.SetupAuthenticationAttempt(u)
	if err != nil {
		return errors.Join(ErrPreparingAuthenticationAttempt, err)
	}
	return nil
}

func (a access) PerformAuthentication(u user.User, credential Credential) (bool, error) {
	authenticated, err := a.defaultAuthenticationMethod.Authenticate(u, credential)
	if err != nil {
		return false, errors.Join(ErrPerformingAuthentication, err)
	}
	return authenticated, nil
}

var ErrPreparingAuthenticationAttempt = errors.New("error preparing authentication attempt")
var ErrPerformingAuthentication = errors.New("error performing authentication")
