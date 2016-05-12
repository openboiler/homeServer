package model

type CommonTime struct {
	ID        uint `gorm:"AUTOINCREMENT"`
	ProfileID uint
	SinceMin  uint
	UntilMin  uint
}

func (c *CommonTime) SetSince(hour, min uint) {
	c.SinceMin = (hour * 60) + min
}

func (c *CommonTime) SetUntil(hour, min uint) {
	c.UntilMin = (hour * 60) + min
}
