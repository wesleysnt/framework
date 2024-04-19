package facades

import "github.com/wesleysnt/framework/contracts/mail"

func Mail() mail.Mail {
	return App().MakeMail()
}
