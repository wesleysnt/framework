package facades

import (
	"github.com/wesleysnt/framework/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
