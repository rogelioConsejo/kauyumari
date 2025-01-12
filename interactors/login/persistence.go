package login

import "github.com/rogelioConsejo/kauyumari/entities/user"

type RegistryPersistence interface {
	SaveUser(u user.User) error
	UserWasSaved(u user.Name) (bool, error)
	SaveConfirmationCode(u user.Name, c ConfirmationCode) error
}

type AccessPersistence interface {
}

type DoorLockPersistence interface {
}
