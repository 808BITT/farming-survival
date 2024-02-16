package db

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

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

func (d *Database) Load() {
	if d.FilePath == "" {
		log.Fatal("Database file path is empty")
	}

	// check if the database file exists
	filePath := strings.Split(d.FilePath, "?")[0]
	filePathParts := strings.Split(filePath, ":")[1:]
	filePath = path.Join(filePathParts...)
	if _, file := os.Stat(filePath); os.IsNotExist(file) {
		err := os.MkdirAll(path.Dir(filePath), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// create a new connection to the database
	conn, err := sql.Open("sqlite3", d.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	// check if the connection is valid
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	d.Connection = conn

	// create the tables if they don't exist
	err = d.CreateSettingsTable()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) CreateSettingsTable() error {
	_, err := d.Exec("CREATE TABLE IF NOT EXISTS settings (key TEXT NOT NULL UNIQUE, value TEXT NOT NULL);")
	if err != nil {
		return err
	}

	_, err = d.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_settings_key ON settings (key);")
	if err != nil {
		return err
	}

	err = d.SetSettings(map[string]string{
		"version":              "0.0.1",
		"screen_width":         "1920",
		"screen_height":        "1080",
		"fullscreen":           "false",
		"window_resizing_mode": "enabled",
		"window_title":         "Homestead Harmony",
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) Close() {
	d.Connection.Close()
}

func (d *Database) LoadQuery(fs embed.FS, filename string) string {
	query, err := fs.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(query)
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.Connection.Query(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Connection.Exec(query, args...)
}

func (d *Database) Prepare(query string) (*sql.Stmt, error) {
	return d.Connection.Prepare(query)
}

func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.Connection.QueryRow(query, args...)
}

func (d *Database) GetVersion() (string, error) {
	return d.GetSetting("version")
}

func (d *Database) SetVersion(version string) error {
	return d.SetSetting("version", version)
}

func (d *Database) GetSetting(key string) (string, error) {
	var value string
	err := d.QueryRow("SELECT value FROM settings WHERE key = ?", key).Scan(&value)
	return value, err
}

func (d *Database) SetSetting(key string, value string) error {
	_, err := d.Exec("INSERT OR REPLACE INTO settings (key, value) VALUES (?, ?)", key, value)
	return err
}

func (d *Database) GetSettings() (map[string]string, error) {
	rows, err := d.Query("SELECT key, value FROM settings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	settings := make(map[string]string)
	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}
		settings[key] = value
	}
	return settings, nil
}

func (d *Database) SetSettings(settings map[string]string) error {
	tx, err := d.Connection.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for key, value := range settings {
		_, err = tx.Exec("INSERT OR REPLACE INTO settings (key, value) VALUES (?, ?)", key, value)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
