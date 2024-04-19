package facades

import (
	"github.com/wesleysnt/framework/contracts/database/orm"
)

func Orm() orm.Orm {
	return App().MakeOrm()
}
