package facades

import "github.com/wesleysnt/framework/contracts/event"

func Event() event.Instance {
	return App().MakeEvent()
}
