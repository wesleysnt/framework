package facades

import (
	"github.com/wesleysnt/framework/contracts/auth"
	"github.com/wesleysnt/framework/contracts/auth/access"
)

func Auth() auth.Auth {
	return App().MakeAuth()
}

func Gate() access.Gate {
	return App().MakeGate()
}
