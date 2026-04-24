package models

import (
	"errors"
	"log"
	"logAnalyzer/pkg"
	"math"
	"time"
	"gorm.io/gorm"
)

type ErrorGroup struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	GroupKey    string `gorm:"size:500;uniqueIndex"`
	Count       int
	BaselineAVG float64
	LastSeen    time.Time
	Confidence  string `gorm:"size:6"`
	AISummary   string `gorm:"type:text"`
}

type Period struct {
	T1    time.Time
	T2    time.Time
	Count int
}

func (eg *ErrorGroup) CreateErrorG(lastSeen time.Time, key string) error {
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
	//get the latest records in 10 min on logs table
	groupKey := pkg.DecodeGroupKey(eg.GroupKey)
	if groupKey == nil {
		return errors.New("Error While Decoding groupKey")
	}
	count, err := GetLatestLogsCount(groupKey["service"],
		groupKey["level"])
	if err != nil {
		return err
	}
	eg.Count = count
	eg.UpdateErrorG()
	return err
}

func (eg *ErrorGroup) CalBaselineAVG() {
	var periods []Period

	var time1 time.Time
	var time2 time.Time
	groupKey := pkg.DecodeGroupKey(eg.GroupKey)

	for t := -10; t > -60; t += -10 {
		var count int64
		if t == -10 {
			time1 = eg.LastSeen.Add(time.Duration(t) * time.Minute)
			time2 = eg.LastSeen
		} else {
			time1 = eg.LastSeen.Add(time.Duration(t) * time.Minute)
			time2 = time1.Add(10 * time.Minute)
		}
		if err := db.Model(&Log{}).
			Where("service = ? AND level= ? AND created_at >= ? AND created_at < ?",
				groupKey["service"], groupKey["level"], time1, time2).
			Count(&count).Error; err != nil {
			log.Println(err)
		}
		if count != 0{
			periods = append(periods, Period{
			T1: time1,
			T2: time2,
			Count:  int(count),
		})
		}
	}
	sum := 0
	for _, p := range periods {
		sum += p.Count
	}
	if len(periods) > 0 {
		eg.BaselineAVG = math.Round(float64(sum) / float64(len(periods)))
	}
	eg.UpdateErrorG()
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

func (eg *ErrorGroup) UpdateErrorG() {
	gorm.G[ErrorGroup](db).Where("id=?", eg.ID).Updates(ctx, *eg)
}

// func (eg *ErrorGroup) GetErrorG() error {
// 	mg, err := gorm.G[ErrorGroup](db).Where("mp_id=?", id).Find(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return mg, err
// }
