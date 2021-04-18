package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DBconfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func GetDB() *sql.DB {
	dbConfig := newDBconfig()
	db := dbConfig.Connect()
	return db
}

func newDBconfig() DBconfig {
	var (
		host     = "80.78.246.133"
		user     = "admin"
		port     = "5432"
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = "main"
	)
	return DBconfig{host, port, user, password, dbname}
}

func (db *DBconfig) Connect() *sql.DB {
	database, err := sql.Open("postgres", db.buildConfigString())
	if err != nil {
		log.Fatal("cant connect to database:", err)
	}
	if err = database.Ping(); err != nil {
		log.Fatal("cant ping database:", err)
	}
	return database
}

func (db *DBconfig) buildConfigString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.user, db.password, db.dbname,
	)
}
