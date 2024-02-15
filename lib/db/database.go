package db

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed *.sql
var sqlFiles embed.FS

type Database struct {
	FilePath   string
	Connection *sql.DB
}

type SqlLiteConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewDatabase(dbPath string) *Database {
	db := &Database{FilePath: dbPath}
	db.Load()

	return db
}

func (db *Database) Load() {
	if db.FilePath == "" {
		log.Fatal("Database file path is empty")
	}

	// check if the database file exists
	if _, file := os.Stat(db.FilePath); os.IsNotExist(file) {
		err := os.MkdirAll(db.FilePath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create(db.FilePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// create a new connection to the database
	conn, err := sql.Open("sqlite3", db.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	// check if the connection is valid
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.Connection = conn

	// create the tables if they don't exist
	query := db.LoadQuery(sqlFiles, "create_settings.sql")
	_, err = db.Connection.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) LoadQuery(fs embed.FS, filename string) string {
	query, err := fs.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(query)

}

func (d *Database) Close() {
	d.Connection.Close()
}
