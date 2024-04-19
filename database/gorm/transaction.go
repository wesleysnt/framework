package gorm

import (
	"gorm.io/gorm"

	"github.com/wesleysnt/framework/contracts/config"
	"github.com/wesleysnt/framework/contracts/database/orm"
)

type Transaction struct {
	orm.Query
	instance *gorm.DB
}

func NewTransaction(tx *gorm.DB, config config.Config) *Transaction {
	return &Transaction{Query: NewQueryImplByInstance(tx, &QueryImpl{
		config:        config,
		withoutEvents: false,
	}), instance: tx}
}

func (r *Transaction) Commit() error {
	return r.instance.Commit().Error
}

func (r *Transaction) Rollback() error {
	return r.instance.Rollback().Error
}
