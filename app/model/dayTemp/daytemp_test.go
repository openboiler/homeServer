package dayTemp

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func TestBasic(t *testing.T) {
	db, err := gorm.Open("sqlite3", "/tmp/openboiler_test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.CreateTable(&DayTemp{})

}
