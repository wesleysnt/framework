package facades

import (
	"github.com/wesleysnt/framework/contracts/crypt"
)

func Crypt() crypt.Crypt {
	return App().MakeCrypt()
}
