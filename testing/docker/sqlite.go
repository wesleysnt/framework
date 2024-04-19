package docker

import (
	"fmt"

	"github.com/ory/dockertest/v3"

	contractsconfig "github.com/wesleysnt/framework/contracts/config"
	"github.com/wesleysnt/framework/contracts/database/orm"
	"github.com/wesleysnt/framework/contracts/testing"
	"github.com/wesleysnt/framework/support/file"
)

type Sqlite struct {
	config     contractsconfig.Config
	connection string
}

func NewSqlite(config contractsconfig.Config, connection string) *Sqlite {
	return &Sqlite{
		config:     config,
		connection: connection,
	}
}

func (receiver *Sqlite) Config(resource *dockertest.Resource) testing.Config {
	return testing.Config{
		Database: receiver.config.GetString(fmt.Sprintf("database.connections.%s.database", receiver.connection)),
	}
}

func (receiver *Sqlite) Clear(pool *dockertest.Pool, resource *dockertest.Resource) error {
	return file.Remove(receiver.config.GetString(fmt.Sprintf("database.connections.%s.database", receiver.connection)))
}

func (receiver *Sqlite) Name() orm.Driver {
	return orm.DriverSqlite
}

func (receiver *Sqlite) Image() *dockertest.RunOptions {
	return &dockertest.RunOptions{
		Repository: "nouchka/sqlite3",
		Tag:        "latest",
		Env:        []string{},
	}
}
