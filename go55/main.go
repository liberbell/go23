package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`

	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = `INSERT INTO person (name, age) VALUES (?, ?)`
	// _, err = DbConnection.Exec(cmd, "Alex", 32)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	cmd = `UPDATE person SET age = ? WHERE name = ?`
	_, err := DbConnection.Exec(cmd, "Eric", 82)
	if err != nil {
		log.Fatalln(err)
	}
}
