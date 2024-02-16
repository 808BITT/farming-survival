package db

import (
	"database/sql"
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
	d.Connection = conn

	// create the settings table if it doesn't exist
	err = d.CreateSettingsTable()
	if err != nil {
		log.Fatal(err)
	}

	// create the tiles table if it doesn't exist
	err = d.createTileTypeTable()
	if err != nil {
		log.Fatal(err)
	}

	// create the map chunk tables if they don't exist
	err = d.CreateMapChunkTable()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) Close() {
	d.Connection.Close()
}

func (d *Database) query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.Connection.Query(query, args...)
}

func (d *Database) exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Connection.Exec(query, args...)
}

func (d *Database) queryRow(query string, args ...interface{}) *sql.Row {
	return d.Connection.QueryRow(query, args...)
}

func (d *Database) CreateSettingsTable() error {
	_, err := d.exec("CREATE TABLE IF NOT EXISTS settings (key TEXT NOT NULL UNIQUE, value TEXT NOT NULL);")
	if err != nil {
		return err
	}

	_, err = d.exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_settings_key ON settings (key);")
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

func (d *Database) GetSetting(key string) (string, error) {
	var value string
	err := d.queryRow("SELECT value FROM settings WHERE key = ?", key).Scan(&value)
	return value, err
}

func (d *Database) SetSetting(key string, value string) error {
	_, err := d.exec("INSERT OR REPLACE INTO settings (key, value) VALUES (?, ?)", key, value)
	return err
}

func (d *Database) GetSettings() (map[string]string, error) {
	rows, err := d.query("SELECT key, value FROM settings")
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

func (d *Database) createTileTypeTable() error {
	_, err := d.exec("CREATE TABLE IF NOT EXISTS tile_types (name TEXT NOT NULL UNIQUE, data BLOB NOT NULL);")
	if err != nil {
		return err
	}

	_, err = d.exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_tile_types_name ON tile_types (name);")
	if err != nil {
		return err
	}

	// create default tile_types
	grassData := `{
		"walkable": true,
		"sprite": "grass",
		"color": "#00FF00"
	}`
	err = d.AddTileType("grass", grassData)
	if err != nil {
		return err
	}

	rockData := `{
		"walkable": false,
		"sprite": "rock",
		"color": "#808080"
	}`
	err = d.AddTileType("rock", rockData)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) AddTileType(name string, data string) error {
	_, err := d.exec("INSERT OR REPLACE INTO tile_types (name, data) VALUES (?, ?)", name, data)
	return err
}

func (d *Database) TileType(name string) (string, error) {
	var data string
	err := d.queryRow("SELECT data FROM tile_types WHERE name = ?", name).Scan(&data)
	return data, err
}

func (d *Database) CreateMapChunkTable() error {
	_, err := d.exec("CREATE TABLE IF NOT EXISTS chunks (id INTEGER PRIMARY KEY, x INTEGER NOT NULL, y INTEGER NOT NULL, data BLOB NOT NULL);")
	if err != nil {
		return err
	}

	_, err = d.exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_chunks_x_y ON chunks (x, y);")
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetChunk(x, y int) (string, error) {
	var data string
	err := d.queryRow("SELECT data FROM chunks WHERE x = ? AND y = ?", x, y).Scan(&data)
	return data, err
}

func (d *Database) SetChunk(x, y int, data string) error {
	_, err := d.exec("INSERT OR REPLACE INTO chunks (x, y, data) VALUES (?, ?, ?)", x, y, data)
	return err
}
