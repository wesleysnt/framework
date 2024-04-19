package testing

import (
	"github.com/wesleysnt/framework/contracts/foundation"
	"github.com/wesleysnt/framework/contracts/testing"
	"github.com/wesleysnt/framework/testing/docker"
)

type Application struct {
	app foundation.Application
}

func NewApplication(app foundation.Application) *Application {
	return &Application{
		app: app,
	}
}

func (receiver *Application) Docker() testing.Docker {
	return docker.NewDocker(receiver.app)
}
