package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func New() *DB {
	db, err := sql.Open("sqlite3", "./expense_tracker.db")
	if err != nil {
		return nil
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		f_name varchar(20), l_name varchar(20),
		email varchar(25) unique,
		password varchar(255)
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name varchar(20) unique, 
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE DEFAULT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE, 
		category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE NOT NULL, 
		title varchar(20), 
		amount real, 
		due_date varchar(10),
		is_paid tinyint
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS incomes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE, 
		title varchar(20), 
		amount real, 
		date varchar(10))`)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}
