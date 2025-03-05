package connection

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	host := "localhost"
	port := "3306"
	user := "root"
	password := "password"
	dbName := "dbcontrol"

	db := NewDB(host, port, user, password, dbName)

	if db == nil {
		t.Error("Expected a non-nil DB instance, but got nil")
	}

}

func TestGetConnection(t *testing.T) {
	host := "localhost"
	port := "3306"
	user := "root"
	password := "root"
	dbName := "dbcontrol"

	db := NewDB(host, port, user, password, dbName)

	// Test with SQLite
	connection := db.GetConnection("sqlite")
	if connection == nil {
		t.Error("Expected a non-nil SQLite connection, but got nil")
	}
}
