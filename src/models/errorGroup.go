package models

import (
	"time"

	"gorm.io/gorm"
)

type ErrorGroup struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	GroupKey  string `gorm:"size:500;uniqueIndex"`
	Count     int
	LastSeen  time.Time
	AISummary string `gorm:"type:text"`
}

func (eg *ErrorGroup) CreateErrorG(lastSeen time.Time,key string) error {
	eg.Count = 1
	eg.LastSeen = lastSeen
	eg.GroupKey = key
	err := gorm.G[ErrorGroup](db).Create(ctx, eg)
	if err != nil {
		return err
	}
	return nil
}

func (eg *ErrorGroup) IncraeseCount() error {
	eg.Count++
	eg.LastSeen = time.Now()
	_, err := gorm.G[ErrorGroup](db).Where("id=?", eg.ID).Updates(ctx, *eg)
	return err
}

func GetErrorGbyKey(groupKey string) (*ErrorGroup, error) {
	egs, err := gorm.G[ErrorGroup](db).Where("group_key=?", groupKey).Find(ctx)
	if err != nil {
		return nil, err
	}
	if len(egs) < 1 {
		eg := ErrorGroup{}
		return &eg, nil
	}
	eg := &egs[0]
	return eg, nil
}

// func (eg *ErrorGroup) GetErrorG() error {
// 	mg, err := gorm.G[ErrorGroup](db).Where("mp_id=?", id).Find(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return mg, err
// }
