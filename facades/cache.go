package facades

import (
	"github.com/wesleysnt/framework/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
