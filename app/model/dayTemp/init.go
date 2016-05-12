package dayTemp

import (
	"github.com/jinzhu/gorm"
)

type DayTemp struct {
	DayWeekID uint `gorm:"primary_key"`
	DayID     uint `gorm:"primary_key"`
	ZoneID    uint `gorm:"primary_key"`
}

func NewDayTemp(zoneId, dayId uint) *DayTemp {
	return &DayTemp{DayID: dayId, ZoneID: zoneId}
}

func NewDayWeekTemp(zoneId, dayWeekId uint) *DayTemp {
	return &DayTemp{DayWeekID: dayWeekId, ZoneID: zoneId}
}

func FindByDayId(db *gorm.DB, zoneId, dayId uint) (*DayTemp, error) {
	var dayTemp DayTemp
	res := db.Where("zone_id = ? AND day_id = ?", zoneId, dayId).First(&dayTemp)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &dayTemp, res.Error
}

func FindByDayWeekId(db *gorm.DB, zoneId, dayWeekId uint) (*DayTemp, error) {
	var dayTemp DayTemp
	res := db.Where("zone_id = ? AND day_week_id = ?", zoneId, dayWeekId).First(&dayTemp)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &dayTemp, res.Error
}

func (e *DayTemp) Save(db *gorm.DB) error {
	res := db.Create(e)
	return res.Error
}

func (e *DayTemp) Delete(db *gorm.DB) error {
	res := db.Delete(e)
	return res.Error
}

func (e *DayTemp) Update(db *gorm.DB) error {
	res := db.Save(e)
	return res.Error
}
