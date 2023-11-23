package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// nolint: golint
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func getConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
	}
}

func New() *DB {
	cfg := getConfig()
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		return nil
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println("Connected!")

	runMigration(db)

	return &DB{db}
}

func runMigration(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		f_name varchar(20), l_name varchar(20),
		email varchar(25) unique,
		password varchar(255)
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY, 
		name varchar(20) unique, 
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE DEFAULT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE, 
		category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE NOT NULL, 
		title varchar(20), 
		amount real, 
		due_date varchar(10),
		is_paid boolean
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS incomes (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE, 
		title varchar(20), 
		amount real, 
		date varchar(10))`)
	if err != nil {
		log.Fatal(err)
	}
}
