package facades

import (
	"github.com/wesleysnt/framework/contracts/route"
)

func Route() route.Route {
	return App().MakeRoute()
}
