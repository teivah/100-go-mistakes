package main

import "database/sql"

func listing1(db *sql.DB, id string) error {
	stmt, err := db.Prepare("SELECT * FROM ORDER WHERE ID = ?")
	if err != nil {
		return err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}
	_ = rows
	return nil
}
