package zoneLog

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ZoneLog struct {
	ZoneID       uint      `gorm:"primary_key"`
	Aquired      time.Time `gorm:"primary_key"`
	CurrentTemp  float32
	SettedTemp   float32
	ExternalTemp float32
	Unit         string
}

func FindByZoneId(db *gorm.DB, zoneId uint) ([]ZoneLog, error) {
	zoneLogs := make([]ZoneLog, 0)
	res := db.Where("zone_id = ? ", zoneId).Find(&zoneLogs)
	return zoneLogs, res.Error
}

func FindByZoneIdTime(db *gorm.DB, zoneId uint, since, to time.Time) ([]ZoneLog, error) {
	zoneLogs := make([]ZoneLog, 0)
	res := db.Where("zone_id = ? AND aquired>=? AND aquired<=?", zoneId, since, to).Find(&zoneLogs)
	return zoneLogs, res.Error
}

func (e *ZoneLog) Save(db *gorm.DB) error {
	return db.Create(e).Error
}

func (e *ZoneLog) Delete(db *gorm.DB) error {
	return db.Delete(e).Error
}

func (e *ZoneLog) Update(db *gorm.DB) error {
	return db.Save(e).Error
}
