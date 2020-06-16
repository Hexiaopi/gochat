package database

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDB("../../etc/database/mysql.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
