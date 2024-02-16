package game

import (
	"fs/lib/db"
	"fs/lib/entity"
	"image/color"
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
	Db       *db.Database
	Entities *entity.Manager
}

func NewEngine() *Engine {
	manager := entity.NewManager()
	manager.AddEntity(entity.NewEntity(100, 100, &entity.EntityType{
		Name:  "player",
		Color: &[]color.Color{{R: 255, G: 0, B: 0, A: 255}}[0],
	}))

	manager.AddEntity(entity.NewEntity(200, 200, &entity.EntityType{
		Name:  "player",
		Color: &[]color.Color{{R: 0, G: 255, B: 0, A: 255}}[0],
	}))

	return &Engine{
		Entities: entity.NewManager(),
	}
}

func (e *Engine) Run() {
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	if e.Entities == nil {
		e.Entities = entity.NewManager()
	}
	e.Entities.Update()
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	if e.Entities == nil {
		return
	}
	for _, entity := range e.Entities.Entities {
		entity.Draw(screen)
	}
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
