package akane

import (
	"testing"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func TestSelectOne(t *testing.T) {
	db, err := Open("sqlite3", "./test.db")
	if err != nil {
		t.Error(err)
	}

	name, err := db.SelectOne(`SELECT name FROM sample WHERE id = ?`, 1)
	if err != nil {
		t.Error(err)
	}
	if name != "akane" {
		t.Error("sample's name should be akane")
	}
}

func TestSelectRow(t *testing.T) {
	db, err := Open("sqlite3", "./test.db")
	if err != nil {
		t.Error(err)
	}

	r, err := db.SelectRow(`SELECT * FROM sample WHERE id = ?`, 1)
	if err != nil {
		t.Error(err)
	}
	d := map[string]interface{}{"id":int64(1), "name":"akane"}
	if !reflect.DeepEqual(r, d) {
		t.Error("sample's row different")
	}
}

func TestSelectAll(t *testing.T) {
	db, err := Open("sqlite3", "./test.db")
	if err != nil {
		t.Error(err)
	}

	r, err := db.SelectAll(`SELECT * FROM sample`)
	if err != nil {
		t.Error(err)
	}
	d := []map[string]interface{}{
		map[string]interface{}{"id":int64(1), "name":"akane"},
		map[string]interface{}{"id":int64(2), "name":"miyuki"},
		map[string]interface{}{"id":int64(3), "name":"yayoi"},
	}
	if !reflect.DeepEqual(r, d) {
		t.Error("sample's rows different")
	}
}
