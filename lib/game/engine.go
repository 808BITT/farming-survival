package game

import (
	"embed"
	"errors"
	"fs/lib/db"
	"fs/lib/entity"
	"fs/lib/wfc"
	"image/png"
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
	screenWidth  int
	screenHeight int
	Assets       *embed.FS
	Db           *db.Database
	Entities     *entity.Manager
	TileMap      *TileMap
	WaveArray2d  *wfc.CollapseArray2d
}

type TileMap struct {
	Width  int
	Height int
	Tiles  [][]int
}

func NewEngine(assets *embed.FS) *Engine {
	tiles := loadTiles(assets)

	ebiten.SetFullscreen(true)
	width, height := 1920, 1080

	wa := wfc.Generate2dMap(width/16, height/16, tiles)
	wa.Set(10, 10, tiles[1])

	return &Engine{
		screenWidth:  width,
		screenHeight: height,
		Assets:       assets,
		TileMap:      NewTileMap(width/16, height/16),
		WaveArray2d:  wa,
	}
}

func loadTiles(assets *embed.FS) []*wfc.Tile {
	blankTile := wfc.NewTile(
		LoadImage(assets, "assets/blank.png"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
	)

	horizontalRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/horizontal.png"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
	)

	verticalRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/vertical.png"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
	)

	bottomLeftRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/bottom_left.png"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
	)

	bottomRightRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/bottom_right.png"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
		wfc.NewEdge("road"),
	)

	topLeftRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/top_left.png"),
		wfc.NewEdge("road"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
	)

	topRightRoadTile := wfc.NewTile(
		LoadImage(assets, "assets/top_right.png"),
		wfc.NewEdge("road"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("grass"),
		wfc.NewEdge("road"),
	)

	tiles := []*wfc.Tile{
		&blankTile,
		&horizontalRoadTile,
		&verticalRoadTile,
		&bottomLeftRoadTile,
		&bottomRightRoadTile,
		&topLeftRoadTile,
		&topRightRoadTile,
	}

	return tiles
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
		e.WaveArray2d.Iterate()
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	for y := 0; y < e.WaveArray2d.Height(); y++ {
		for x := 0; x < e.WaveArray2d.Width(); x++ {
			tile, err := e.WaveArray2d.Get(x, y)
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

func LoadImage(assets *embed.FS, path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("Failed to decode image file: %v", err)
	}

	image := ebiten.NewImageFromImage(img)
	return image
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
