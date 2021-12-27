package main

import (
	"database/sql"
	"log"
)

func listing1() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM CUSTOMERS")
	if err != nil {
		return err
	}

	// Use rows
	_ = rows

	return nil
}

func listing2() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM CUSTOMERS")
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	// Use rows
	_ = rows

	return nil
}

var dataSourceName = ""
