package game

import (
	"errors"
	"fs/lib/assets"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type EngineState int

const (
	Menu EngineState = iota
	Loading
	GameLoop
	Paused
	Saving
)

type Engine struct {
	State    EngineState
	Screen   *Screen
	TileSize int
	Tiles    []*Tile
}

func NewEngine(w, h int) *Engine {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetCursorShape(ebiten.CursorShapePointer)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowDecorated(true)
	ebiten.SetWindowTitle("Idk...")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetFullscreen(false)

	screen := NewScreen(float64(w), float64(h), true)
	return &Engine{
		State:    GameLoop,
		Screen:   screen,
		Tiles:    []*Tile{},
		TileSize: 16,
	}
}

func (e *Engine) Run() {
	worldBorderWidth := int(e.Screen.Dim.Width / float64(e.TileSize) / 2)
	worldBorderHeight := int(e.Screen.Dim.Height / float64(e.TileSize) / 2)
	tilemapWidth := 200
	tilemapHeight := 100

	log.Println("World border width:", worldBorderWidth)
	log.Println("World border height:", worldBorderHeight)

	for y := 0; y < 2*worldBorderHeight+tilemapHeight; y++ {
		for x := 0; x < 2*worldBorderWidth+tilemapWidth; x++ {
			if x < worldBorderWidth || x > worldBorderWidth+tilemapWidth-1 || y < worldBorderHeight || y > worldBorderHeight+tilemapHeight-1 {
				tile := NewTile(x*e.TileSize, y*e.TileSize, e.TileSize, e.TileSize, assets.LoadImage("tile/world_border.png"))
				e.Tiles = append(e.Tiles, tile)
				continue
			}

			tile := NewTile(x*e.TileSize, y*e.TileSize, e.TileSize, e.TileSize, assets.LoadImage("tile/grass.png"))
			e.Tiles = append(e.Tiles, tile)
		}
	}

	// os.Exit(0)

	if err := ebiten.RunGame(e); err != nil {
		if err.Error() == "user requested to quit" {
			return
		}
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("user requested to quit")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		e.MoveScreen(0, 1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		e.MoveScreen(1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		e.MoveScreen(0, 1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		e.MoveScreen(1, 0)
	}

	// var wg sync.WaitGroup
	// for _, boid := range e.Boids {
	// 	wg.Add(1)
	// 	go func(b *Boid) {
	// 		defer wg.Done()
	// 		targetHit := b.Update(e.BoidTarget, e.Boids)
	// 		if targetHit {
	// 			randX, randY := rand.Float64()*e.Screen.Width-25, rand.Float64()*e.Screen.Height-25
	// 			e.BoidTarget = phys.NewVec2(randX, randY)
	// 		}
	// 	}(boid)
	// }
	// wg.Wait()

	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	visibleTiles := GetVisibleTiles(e.Screen, e.Tiles)
	for _, tile := range visibleTiles {
		tile.Draw(screen)
	}

	debugPanel := ebiten.NewImage(50, 20)
	ebitenutil.DebugPrint(debugPanel, "FPS: "+strconv.Itoa(int(ebiten.ActualFPS())))
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Scale(3, 3)
	drawOptions.GeoM.Translate(0, 0)
	screen.DrawImage(debugPanel, drawOptions)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(e.Screen.Dim.Width), int(e.Screen.Dim.Height)
}

func GetVisibleTiles(screen *Screen, tiles []*Tile) []*Tile {
	visibleTiles := []*Tile{}
	for _, tile := range tiles {
		t := tile
		if t.Visible(int(screen.Dim.X), int(screen.Dim.Y), int(screen.Dim.Width), int(screen.Dim.Height)) {
			visibleTiles = append(visibleTiles, t)
		}

	}
	return visibleTiles
}

func (e *Engine) MoveScreen(x, y int) {
	e.Screen.Dim.X += float64(x)
	e.Screen.Dim.Y += float64(y)
}
