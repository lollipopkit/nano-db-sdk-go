package ndb_test

import (
	"testing"

	"github.com/lollipopkit/nano-db-sdk-go"
)

const (
	dbUrl          = "http://localhost:3770"
	dbCookie       = "FHYmGdNwfiJngvF2z"
)

var (
	db = ndb.NewClient(dbUrl, dbCookie)
)

func TestWrite(t *testing.T) {
	err := db.Write("novel", "23", "27145", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestRead(t *testing.T) {
	data, err := db.Read("novel", "23", "27145")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestDelete(t *testing.T) {
	err := db.Delete("novel", "23", "27145")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestReadNoSuchFile(t *testing.T) {
	_, err := db.Read("novel", "23", "27146")
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)
}

func TestDeleteNoSuchFile(t *testing.T) {
	err := db.Delete("novel", "23", "27146")
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)
}

func TestDirs(t *testing.T) {
	dirs, err := db.Read("novel")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirs)
}

func TestFiles(t *testing.T) {
	files, err := db.Read("novel", "23")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(files))
}

func TestDeleteDir(t *testing.T) {
	err := db.Delete("novel", "23")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestDeleteDB(t *testing.T) {
	err := db.Delete("novel")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestIllegalPath(t *testing.T) {
	_, err := db.Do("GET", "novel/../../123", nil)
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)

	_, err = db.Read(".")
	if err == nil {
		t.Fatal("should be error")
	}
	t.Log(err)
}
