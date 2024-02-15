package game

import (
	"fs/lib/db"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

type Engine struct {
	Database *db.Database
}

func NewEngine() *Engine {
	return &Engine{
		Database: db.NewDatabase("file:data/map.db?cache=shared&mode=rwc"),
	}
}

func (e *Engine) Run() {
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
	e.Database.Connection.Close()
}

func (e *Engine) Update() error {
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
