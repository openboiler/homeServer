package zone

import (
	"github.com/jinzhu/gorm"
)

type Zone struct {
	ID   uint
	Name string
}

func FindById(db *gorm.DB, id uint) (*Zone, error) {
	var zone Zone
	res := db.First(&zone, id)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &zone, res.Error
}

func (e *Zone) Save(db *gorm.DB) error {
	var res *gorm.DB
	if e.ID == 0 {
		res = db.Create(e)
	} else {
		res = db.Save(e)
	}
	return res.Error
}

func (e *Zone) Delete(db *gorm.DB) error {
	res := db.Delete(e)
	return res.Error
}

func (e *Zone) Update(db *gorm.DB) error {
	res := db.Save(e)
	return res.Error
}
