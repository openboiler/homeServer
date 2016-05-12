package day

import (
	"github.com/jinzhu/gorm"
	"github.com/openboiler/homeServer/app/model"
	"time"
)

type Day struct {
	model.CommonTime
	Day time.Time
}

func FindById(db *gorm.DB, id uint) (*Day, error) {
	var day Day
	res := db.First(&day, id)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &day, res.Error
}

func (e *Day) Save(db *gorm.DB) error {
	var res *gorm.DB
	if e.ID == 0 {
		res = db.Create(e)
	} else {
		res = db.Save(e)
	}
	return res.Error
}

func (e *Day) Delete(db *gorm.DB) error {
	res := db.Delete(e)
	return res.Error
}

func (e *Day) Update(db *gorm.DB) error {
	res := db.Save(e)
	return res.Error
}
