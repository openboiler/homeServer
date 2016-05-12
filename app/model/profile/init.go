package profile

import (
	"github.com/jinzhu/gorm"
)

type Profile struct {
	ID   uint `gorm:"AUTOINCREMENT"`
	Name string
}

func FindById(db *gorm.DB, id uint) (*Profile, error) {
	var profile Profile
	res := db.First(&profile, id)
	if res.RecordNotFound() {
		return nil, nil
	}
	return &profile, res.Error
}

func (e *Profile) Save(db *gorm.DB) error {
	var res *gorm.DB
	if e.ID == 0 {
		res = db.Create(e)
	} else {
		res = db.Save(e)
	}
	return res.Error
}

func (e *Profile) Delete(db *gorm.DB) error {
	res := db.Delete(e)
	return res.Error
}

func (e *Profile) Update(db *gorm.DB) error {
	res := db.Save(e)
	return res.Error
}
