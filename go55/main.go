package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB
