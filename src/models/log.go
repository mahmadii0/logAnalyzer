package models

import (
	"context"
	database "logAnalyzer/data/db"
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Service   string    `gorm:"size:100"`
	Level     string    `gorm:"size:20"`
	Message   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

var (
	db  *gorm.DB
	ctx context.Context
)

func init() {
	db, ctx = database.GetDBctx()
}

func (lg *Log) CreateLog() error {
	err := gorm.G[Log](db).Create(ctx, lg)
	if err != nil {
		return err
	}
	return nil
}
