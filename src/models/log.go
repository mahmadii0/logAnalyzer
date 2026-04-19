package models

import (
	"context"
	database "logAnalyzer/data/db"
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Service   string    `gorm:"size:100" json:"service" binding:"required,service,max=100"`
	Level     string    `gorm:"size:20" json:"level" binding:"required,level,max=20"`
	Message   string    `gorm:"type:text" json:"message" binding:"required,max=380"`
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

func GetLatestLogsCount(service string, level string) (int, error) {
	var count int64
	period := time.Now().Add(-10 * time.Minute)
	err := db.Model(&Log{}).
		Where("service = ? AND level= ? AND created_at > ?", service, level, period).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil

}
