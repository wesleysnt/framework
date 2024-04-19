package facades

import (
	"github.com/wesleysnt/framework/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
