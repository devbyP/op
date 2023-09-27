package sampledb

import (
	"os"
	"testing"
)

var db *Storage

func TestMain(m *testing.M) {
	store := New("./test_temp.csv")
	db = store
	os.Exit(m.Run())
}
