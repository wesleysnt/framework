package facades

import (
	"github.com/wesleysnt/framework/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
