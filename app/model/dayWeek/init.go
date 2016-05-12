package dayWeek

import (
	"github.com/jinzhu/gorm"
	"github.com/openboiler/homeServer/app/model"
)

type DayWeek struct {
	model.CommonTime
	DayWeek uint
}

func FindById(db *gorm.DB, id uint) (*DayWeek, error) {
	var day DayWeek
	res := db.First(&day, id)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &day, res.Error
}

func (e *DayWeek) Save(db *gorm.DB) error {
	var res *gorm.DB
	if e.ID == 0 {
		res = db.Create(e)
	} else {
		res = db.Save(e)
	}
	return res.Error
}

func (e *DayWeek) Delete(db *gorm.DB) error {
	res := db.Delete(e)
	return res.Error
}

func (e *DayWeek) Update(db *gorm.DB) error {
	res := db.Save(e)
	return res.Error
}
