package profile

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

	db.CreateTable(&Profile{})

	p := Profile{Name: "testProfile"}
	err = p.Save(db)
	if err != nil {
		t.Fatal(err)
	}
	if p.ID == 0 {
		t.Fatal("id null")
	}

	p.Name = "modified Profile"
	err = p.Update(db)
	if err != nil {
		t.Fatal(err)
	}

	p2, err := FindById(db, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if p2 == nil {
		t.Fatal("FindById Failed")
	}
	if p2.Name != p.Name {
		t.Fatal("Updated Failed")
	}

	if p.Delete(db) != nil {
		t.Fatalf("Delete Failed")
	}
}
