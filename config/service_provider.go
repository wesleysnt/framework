package config

import (
	"github.com/wesleysnt/framework/contracts/foundation"
	"github.com/wesleysnt/framework/support"
)

const Binding = "goravel.config"

type ServiceProvider struct {
}

func (config *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(support.EnvPath), nil
	})
}

func (config *ServiceProvider) Boot(app foundation.Application) {

}
