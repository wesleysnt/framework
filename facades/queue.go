package facades

import (
	"github.com/wesleysnt/framework/contracts/queue"
)

func Queue() queue.Queue {
	return App().MakeQueue()
}
