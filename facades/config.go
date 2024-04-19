package facades

import (
	"github.com/wesleysnt/framework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
