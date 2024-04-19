package facades

import (
	"github.com/wesleysnt/framework/contracts/testing"
)

func Testing() testing.Testing {
	return App().MakeTesting()
}
