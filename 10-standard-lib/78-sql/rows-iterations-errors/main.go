package main

import (
	"context"
	"database/sql"
	"log"
)

func (c *conn) get1(ctx context.Context, id string) (string, int, error) {
	rows, err := c.db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}

	return department, age, nil
}

func (c *conn) get2(ctx context.Context, id string) (string, int, error) {
	rows, err := c.db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}
	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	return department, age, nil
}

type conn struct {
	db *sql.DB
}
