package tests

import (
	"fs/lib/db"
	"os"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	// delete test.db if it exists
	err := os.Remove("test.db")
	if err != nil {
		t.Log(err)
	}

	d := db.NewDatabase("file:test.db?cache=shared&mode=rwc")
	// test if the database file exists
	t.Run("Test if the database file exists", func(t *testing.T) {
		if d.FilePath == "" {
			t.Error("Database file path is empty")
		}
	})

	// test if the database connection is created
	t.Run("Test if the database connection is created", func(t *testing.T) {
		if d.Connection == nil {
			t.Error("Database connection is nil")
		}
	})

	// test if the database connection is valid
	t.Run("Test if the database connection is valid", func(t *testing.T) {
		err := d.Connection.Ping()
		if err != nil {
			t.Error(err)
		}
	})

	// test if the default settings are set
	t.Run("Test if the default settings are set", func(t *testing.T) {
		defaultSettings := map[string]string{
			"version":              "0.0.1",
			"screen_width":         "1920",
			"screen_height":        "1080",
			"fullscreen":           "false",
			"window_resizing_mode": "enabled",
			"window_title":         "Homestead Harmony",
		}
		version, err := d.GetSetting("version")
		if err != nil {
			t.Error(err)
		}
		if version != defaultSettings["version"] {
			t.Errorf("Expected version to be %s, got %s", defaultSettings["version"], version)
		}

		screenWidth, err := d.GetSetting("screen_width")
		if err != nil {
			t.Error(err)
		}
		if screenWidth != defaultSettings["screen_width"] {
			t.Errorf("Expected screen_width to be %s, got %s", defaultSettings["screen_width"], screenWidth)
		}

		screenHeight, err := d.GetSetting("screen_height")
		if err != nil {
			t.Error(err)
		}
		if screenHeight != defaultSettings["screen_height"] {
			t.Errorf("Expected screen_height to be %s, got %s", defaultSettings["screen_height"], screenHeight)
		}

		fullscreen, err := d.GetSetting("fullscreen")
		if err != nil {
			t.Error(err)
		}
		if fullscreen != defaultSettings["fullscreen"] {
			t.Errorf("Expected fullscreen to be %s, got %s", defaultSettings["fullscreen"], fullscreen)
		}

		windowResizingMode, err := d.GetSetting("window_resizing_mode")
		if err != nil {
			t.Error(err)
		}
		if windowResizingMode != defaultSettings["window_resizing_mode"] {
			t.Errorf("Expected window_resizing_mode to be %s, got %s", defaultSettings["window_resizing_mode"], windowResizingMode)
		}

		windowTitle, err := d.GetSetting("window_title")
		if err != nil {
			t.Error(err)
		}
		if windowTitle != defaultSettings["window_title"] {
			t.Errorf("Expected window_title to be %s, got %s", defaultSettings["window_title"], windowTitle)
		}
	})

	// test if the setting is retrieved
	t.Run("Test if the setting is retrieved", func(t *testing.T) {
		version, err := d.GetSetting("version")
		if err != nil {
			t.Error(err)
		}
		if version != "0.0.1" {
			t.Errorf("Expected version to be 0.0.1, got %s", version)
		}
	})

	// test if the setting is updated
	t.Run("Test if the setting is updated", func(t *testing.T) {
		err := d.SetSetting("version", "0.0.2")
		if err != nil {
			t.Error(err)
		}
		version, err := d.GetSetting("version")
		if err != nil {
			t.Error(err)
		}
		if version != "0.0.2" {
			t.Errorf("Expected version to be 0.0.2, got %s", version)
		}
	})

	// test if multiple settings are set
	t.Run("Test if multiple settings are set", func(t *testing.T) {
		settings := map[string]string{
			"version": "0.0.3",
			"author":  "John Doe",
		}
		err := d.SetSettings(settings)
		if err != nil {
			t.Error(err)
		}
	})

	// test if multiple settings are retrieved
	t.Run("Test if multiple settings are retrieved", func(t *testing.T) {
		settings, err := d.GetSettings()
		if err != nil {
			t.Error(err)
		}
		if settings["version"] != "0.0.3" {
			t.Errorf("Expected version to be 0.0.3, got %s", settings["version"])
		}
		if settings["author"] != "John Doe" {
			t.Errorf("Expected author to be John Doe, got %s", settings["author"])
		}
	})

	err = d.Connection.Close()
	if err != nil {
		t.Error(err)
	}
}
