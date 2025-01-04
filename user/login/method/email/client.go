package email

import "github.com/rogelioConsejo/kauyumari/user"

type Client interface {
	SendToken(email user.Email, tk Token) error
}
