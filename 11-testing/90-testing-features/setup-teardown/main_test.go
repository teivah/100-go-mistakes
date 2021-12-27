package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestMySQLIntegration(t *testing.T) {
	s := setupMySQL()
	defer teardownMySQL(s)

	// ...
}

func createConnection(t *testing.T, dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.FailNow()
	}
	t.Cleanup(
		func() {
			_ = db.Close()
		})
	return db
}

func TestMain(m *testing.M) {
	s := setupMySQL()
	code := m.Run()
	teardownMySQL(s)
	os.Exit(code)
}

func setupMySQL() interface{} { return nil }

func teardownMySQL(m interface{}) {}
