package email

import (
	"errors"
	"github.com/google/uuid"
	"github.com/rogelioConsejo/kauyumari/user"
	"github.com/rogelioConsejo/kauyumari/user/login"
	"time"
)

func GetEmailMethod() login.AuthenticationMethod {
	//TODO: inject Persistence and Client or pass it? Probably it is better to inject them so that nobody else needs to know about them. Otherwise we would need a third party to create the emailMethod.
	return &emailMethod{
		expiration: 15 * time.Minute,
	}
}

type emailMethod struct {
	persistence Persistence
	emailClient Client
	expiration  time.Duration
}

func (e emailMethod) SetupAuthenticationAttempt(u user.User) error {
	tk := Token(uuid.NewString())
	hashedToken := login.HashCredential(login.Credential(tk))
	err := e.persistence.SaveLoginToken(u, HashedToken(hashedToken), time.Now().Add(e.expiration))
	if err != nil {
		return errors.Join(ErrSavingLoginToken, err)
	}
	err = e.emailClient.SendToken(u.Email(), tk)
	if err != nil {
		return errors.Join(ErrSendingLoginToken, err)
	}
	return nil
}

func (e emailMethod) Authenticate(u user.User, credential login.Credential) (bool, error) {
	hashedToken, tokenExpiration, err := e.persistence.GetLoginToken(u.Name())
	if err != nil {
		return false, errors.Join(ErrAuthenticating, ErrRetrievingLoginToken, err)
	}
	if time.Now().After(tokenExpiration) {
		return false, nil
	}

	return login.HashedCredential(hashedToken).Check(credential), nil
}

var ErrSavingLoginToken = errors.New("error saving login token")
var ErrSendingLoginToken = errors.New("error sending login token")
var ErrAuthenticating = errors.New("could not authenticate user")
var ErrRetrievingLoginToken = errors.New("failed to retrieve login token")
