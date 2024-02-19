package game

import (
	"errors"
	"fs/lib/assets"
	"fs/lib/db"
	"fs/lib/entity"
	"fs/lib/wfc"
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
	screenWidth   int
	screenHeight  int
	Db            *db.Database
	Entities      *entity.Manager
	TileMap       *TileMap
	CollapseArray *wfc.CollapseArray2d
}

type TileMap struct {
	Width  int
	Height int
	Tiles  [][]int
}

func NewEngine() *Engine {
	tiles := assets.LoadTestTiles()
	// tiles := assets.LoadTiles()

	ebiten.SetFullscreen(true)
	width, height := 192, 108

	wa := wfc.NewCollapseArray2d(width/16, width/16, tiles)
	// var x, y int
	// for x != -1 && y != -1 {
	// 	x, y = wa.Iterate()
	// }

	return &Engine{
		screenWidth:   width,
		screenHeight:  height,
		CollapseArray: wa,
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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		_, _, err := e.CollapseArray.Iterate()
		if err != nil {
			wa := wfc.NewCollapseArray2d(e.TileMap.Width, e.TileMap.Height, assets.LoadTiles())
			e.CollapseArray = wa
		}
	}

	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	for y := 0; y < e.CollapseArray.Height(); y++ {
		for x := 0; x < e.CollapseArray.Width(); x++ {
			tile, err := e.CollapseArray.GetTile(x, y)
			if err != nil {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*16), float64(y*16))
			screen.DrawImage(tile.Img, op)
		}
	}
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return e.screenWidth, e.screenHeight
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
