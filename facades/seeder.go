package facades

import (
	"github.com/wesleysnt/framework/contracts/database/seeder"
)

func Seeder() seeder.Facade {
	return App().MakeSeeder()
}
