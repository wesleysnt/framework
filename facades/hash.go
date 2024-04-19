package facades

import (
	"github.com/wesleysnt/framework/contracts/hash"
)

func Hash() hash.Hash {
	return App().MakeHash()
}
