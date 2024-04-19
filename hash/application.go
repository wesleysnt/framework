package hash

import (
	"github.com/wesleysnt/framework/contracts/config"
	"github.com/wesleysnt/framework/contracts/hash"
)

const (
	DriverBcrypt string = "bcrypt"
)

type Application struct {
}

func NewApplication(config config.Config) hash.Hash {
	driver := config.GetString("hashing.driver", "argon2id")

	if driver == DriverBcrypt {
		return NewBcrypt(config)
	}

	return NewArgon2id(config)
}
