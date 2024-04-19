package console

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"

	configmock "github.com/wesleysnt/framework/contracts/config/mocks"
	consolemocks "github.com/wesleysnt/framework/contracts/console/mocks"
	ormcontract "github.com/wesleysnt/framework/contracts/database/orm"
	"github.com/wesleysnt/framework/database/gorm"
)

func TestMigrateRefreshCommand(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping tests of using docker")
	}

	var (
		mockConfig *configmock.Config
		pool       *dockertest.Pool
		resource   *dockertest.Resource
		query      ormcontract.Query
	)

	beforeEach := func() {
		pool = nil
		mockConfig = &configmock.Config{}
	}

	tests := []struct {
		name  string
		setup func()
	}{
		{
			name: "mysql",
			setup: func() {
				var err error
				docker := gorm.NewMysqlDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createMysqlMigrations()
			},
		},
		{
			name: "postgresql",
			setup: func() {
				var err error
				docker := gorm.NewPostgresqlDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createPostgresqlMigrations()
			},
		},
		{
			name: "sqlserver",
			setup: func() {
				var err error
				docker := gorm.NewSqlserverDocker()
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqlserverMigrations()
			},
		},
		{
			name: "sqlite",
			setup: func() {
				var err error
				docker := gorm.NewSqliteDocker("goravel")
				pool, resource, query, err = docker.New()
				assert.Nil(t, err)
				mockConfig = docker.MockConfig
				createSqliteMigrations()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()

			mockArtisan := &consolemocks.Artisan{}
			mockContext := &consolemocks.Context{}
			mockContext.On("Option", "step").Return("").Once()

			migrateCommand := NewMigrateCommand(mockConfig)
			assert.Nil(t, migrateCommand.Handle(mockContext))

			// Test MigrateRefreshCommand without --seed flag
			mockContext.On("OptionBool", "seed").Return(false).Once()
			migrateRefreshCommand := NewMigrateRefreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateRefreshCommand.Handle(mockContext))

			var agent Agent
			err := query.Where("name", "goravel").First(&agent)
			assert.Nil(t, err)
			assert.True(t, agent.ID > 0)

			mockArtisan = &consolemocks.Artisan{}
			mockContext = &consolemocks.Context{}
			mockContext.On("Option", "step").Return("5").Once()

			migrateCommand = NewMigrateCommand(mockConfig)
			assert.Nil(t, migrateCommand.Handle(mockContext))

			// Test MigrateRefreshCommand with --seed flag and --seeder specified
			mockContext.On("OptionBool", "seed").Return(true).Once()
			mockContext.On("OptionSlice", "seeder").Return([]string{"UserSeeder"}).Once()
			mockArtisan.On("Call", "db:seed --seeder UserSeeder").Return(nil).Once()
			migrateRefreshCommand = NewMigrateRefreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateRefreshCommand.Handle(mockContext))

			mockArtisan = &consolemocks.Artisan{}
			mockContext = &consolemocks.Context{}

			// Test MigrateRefreshCommand with --seed flag and no --seeder specified
			mockContext.On("Option", "step").Return("").Once()
			mockContext.On("OptionBool", "seed").Return(true).Once()
			mockContext.On("OptionSlice", "seeder").Return([]string{}).Once()
			mockArtisan.On("Call", "db:seed").Return(nil).Once()
			migrateRefreshCommand = NewMigrateRefreshCommand(mockConfig, mockArtisan)
			assert.Nil(t, migrateRefreshCommand.Handle(mockContext))

			var agent1 Agent
			err = query.Where("name", "goravel").First(&agent1)
			assert.Nil(t, err)
			assert.True(t, agent1.ID > 0)

			if pool != nil && test.name != "sqlite" {
				assert.Nil(t, pool.Purge(resource))
			}

			mockContext.AssertExpectations(t)
			removeMigrations()
		})
	}
}
