package main

import "database/sql"

var dsn = ""

func listing1() error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	_ = db
	return nil
}
