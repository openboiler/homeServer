package tiepidyn

import (
	"github.com/jinzhu/gorm"
	"github.com/openboiler/homeServer/app/model/zoneLog"
	"github.com/openboiler/homeServer/app/util"
	"time"
)

func DeltaTime(db *gorm.DB, fromTemp, toTemp, extTemp float32, zoneId uint) time.Duration {
	now := time.Now()
	logs, err := zoneLog.FindByZoneIdTime(db, zoneId, now.AddDate(0, 0, -5), now)

	for i, log := range logs {

	}
}
