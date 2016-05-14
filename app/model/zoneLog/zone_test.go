package zoneLog

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	db, err := gorm.Open("sqlite3", "/tmp/openboiler_test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.CreateTable(&ZoneLog{})

	p := ZoneLog{ZoneID: 1, Aquired: time.Now(), CurrentTemp: 10.4}
	err = p.Save(db)
	if err != nil {
		t.Fatal(err)
	}

	p.CurrentTemp = 19.9
	err = p.Update(db)
	if err != nil {
		t.Fatal(err)
	}

	p2, err := FindByZoneId(db, p.ZoneID)
	if err != nil {
		t.Fatal(err)
	}
	if len(p2) == 0 {
		t.Fatal("FindById Failed")
	}
	if p2[0].CurrentTemp != p.CurrentTemp {
		t.Fatal("Updated Failed")
	}

	if p.Delete(db) != nil {
		t.Fatalf("Delete Failed")
	}
}
