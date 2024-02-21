package game

import (
	"errors"
	"fs/lib/assets"
	"fs/lib/db"
	"fs/lib/entity"
	"fs/lib/wfc"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

type Engine struct {
	ViewportX    float64
	ViewportY    float64
	ScreenWidth  float64
	ScreenHeight float64
	Db           *db.Database
	Entities     *entity.Manager
	TileMap      *wfc.CollapseArray2d
}

type TileMap struct {
	Width  int
	Height int
	Tiles  [][]int
}

func NewEngine() *Engine {
	tiles := assets.LoadTestTiles()

	ebiten.SetFullscreen(true)
	width, height := 1920.0, 1080.0

	wa := wfc.NewCollapseArray2d(int(width/16), int(height/16), tiles)
	for {
		x, y, err := wa.Iterate()
		if err != nil {
			log.Println(err.Error())
		}
		if x == -1 && y == -1 {
			break
		}
	}

	em := entity.NewManager()
	player := entity.NewEntity("Player", width/2, height/2, &color.RGBA{R: 255, G: 0, B: 0, A: 255})
	player.SetSpeed(1.0)
	em.AddEntity(player)

	for i := 0; i < 100; i++ {
		randomEnemy := entity.NewEntity("Enemy", float64(rand.Intn(1920)), float64(rand.Intn(1080)), &color.RGBA{R: 0, G: 255, B: 0, A: 255})
		randomEnemy.SetSpeed(16.0)
		em.AddEntity(randomEnemy)
	}

	return &Engine{
		Entities:     em,
		Db:           db.NewDatabase("file:data/hh.db?cache=shared&mode=rwc"),
		ScreenWidth:  width,
		ScreenHeight: height,
		ViewportX:    0,
		ViewportY:    0,
		TileMap:      wa,
	}
}

func (e *Engine) Run() {
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("user quit")
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		e.Entities.MovePlayer(0, -1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		e.Entities.MovePlayer(0, 1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		e.Entities.MovePlayer(-1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		e.Entities.MovePlayer(1, 0)
	}

	e.Entities.Update()

	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	for y := int(e.ViewportY / 16); y < int(e.ScreenHeight/16); y++ {
		for x := int(e.ViewportX / 16); x < int(e.ScreenWidth/16); x++ {
			tile, err := e.TileMap.GetTile(x, y)
			if err != nil {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*16), float64(y*16))
			screen.DrawImage(tile.Type.Texture, op)
		}
	}

	e.Entities.Draw(screen)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(e.ScreenWidth), int(e.ScreenHeight)
}

func NewTileMap(width, height int) *TileMap {
	tm := &TileMap{
		Width:  width,
		Height: height,
		Tiles:  GenerateRandomTileMap(width, height),
	}
	return tm
}

func GenerateRandomTileMap(width, height int) [][]int {
	tiles := make([][]int, width)
	for x := 0; x < width; x++ {
		tiles[x] = make([]int, height)
		for y := 0; y < height; y++ {
			tiles[x][y] = 0
		}
	}
	return tiles
}

var TileTypeMap = map[int]string{
	0: "Blue_X_16",
	1: "Green_X_16",
	2: "Orange_X_16",
	3: "Purple_X_16",
}

type TileConstraint struct {
	AllowedNeighbors map[string][]int // "north", "south", "east", "west"
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
