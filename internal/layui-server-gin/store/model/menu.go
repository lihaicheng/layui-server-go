package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	menuTable = "menu"
)

type Menu struct {
	gorm.Model
	Flag  string `gorm:"column:flag; type:string; size:36; not null; unique; <-:create;"`
	Value string `gorm:"column:value; type:longtext; "`
}

// TableName returns table name
func (t Menu) TableName() string {
	return tableName(menuTable)
}

// BeforeSave is table operation hooks
func (t *Menu) BeforeSave(tx *gorm.DB) error {
	zap.L().Info("BeforeSave")
	return nil
}

// BeforeCreate is table operation hooks
func (t *Menu) BeforeCreate(tx *gorm.DB) error {
	zap.L().Info("BeforeCreate")
	return nil
}

// BeforeUpdate is table operation hooks
func (t *Menu) BeforeUpdate(tx *gorm.DB) error {
	zap.L().Info("BeforeUpdate")
	return nil
}

// BeforeDelete is table operation hooks
func (t *Menu) BeforeDelete(tx *gorm.DB) error {
	zap.L().Info("BeforeDelete")
	return nil
}
