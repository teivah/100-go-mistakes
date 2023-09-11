package main

import (
	"database/sql"
	"errors"
)

func listing1() {
	err := query()
	if err != nil {
		if err == sql.ErrNoRows {
			// ...
		} else {
			// ...
		}
	}
}

func listing2() {
	err := query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// ...
		} else {
			// ...
		}
	}
}

func query() error {
	return nil
}
