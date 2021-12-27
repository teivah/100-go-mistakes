package main

import (
	"database/sql"
	"log"
)

const query = "..."

func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Use rows
	return 0, nil
}

func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()

	// Use rows
	return 0, nil
}

func getBalance3(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close rows: %v", err)
			}
			return
		}
		err = closeErr
	}()

	// Use rows
	return 0, nil
}
