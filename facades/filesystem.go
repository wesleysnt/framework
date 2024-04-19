package facades

import "github.com/wesleysnt/framework/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
