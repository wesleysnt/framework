package facades

import (
	"github.com/wesleysnt/framework/contracts/schedule"
)

func Schedule() schedule.Schedule {
	return App().MakeSchedule()
}
