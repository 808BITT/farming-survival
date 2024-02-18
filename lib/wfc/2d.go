package wfc

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Id    uuid.UUID
	Name  string
	Img   *ebiten.Image
	North Edge
	West  Edge
	South Edge
	East  Edge
}

func NewTile(name string, image *ebiten.Image, north, west, south, east Edge) Tile {
	return Tile{
		Id:    uuid.New(),
		Name:  name,
		Img:   image,
		North: north,
		West:  west,
		South: south,
		East:  east,
	}
}

func (t Tile) Equals(other Tile) bool {
	if t.Id != other.Id {
		return false
	}

	if t.Name != other.Name {
		return false
	}

	if t.North != other.North {
		return false
	}

	if t.West != other.West {
		return false
	}

	if t.South != other.South {
		return false
	}

	if t.East != other.East {
		return false
	}

	return true
}

func (t Tile) String() string {
	return t.Id.String()
}

type Edge string

func NewEdge(name string) Edge {
	return Edge(name)
}

type CollapseArray2d struct {
	Tiles [][]*CollapseTile
}

type CollapseTile struct {
	Tile     *Tile
	Possible []*Tile
}

func (wa *CollapseArray2d) Get(x, y int) (*Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("Out of bounds")
	}

	if wa.Tiles[y][x].Tile == nil {
		return nil, fmt.Errorf("Tile not collapsed yet")
	}

	return wa.Tiles[y][x].Tile, nil
}

func (wa *CollapseArray2d) CollapeTile(x, y int, tile *Tile) {
	wa.Tiles[y][x].Tile = tile
	wa.Tiles[y][x].Possible = nil
	wa.updatePossible(x, y)
}

func (wa *CollapseArray2d) Width() int {
	return len(wa.Tiles[0])
}

func (wa *CollapseArray2d) Height() int {
	return len(wa.Tiles)
}

func (wa *CollapseArray2d) Set(x, y int, tile *Tile) {
	wa.Tiles[y][x].Tile = tile
	wa.Tiles[y][x].Possible = nil
	wa.updatePossible(x, y)
}

func NewCollapseTile(tile *Tile) *CollapseTile {
	return &CollapseTile{
		Tile:     tile,
		Possible: nil,
	}
}

func (wa *CollapseArray2d) SetPossible(x, y int, possible []*Tile) {
	wa.Tiles[y][x].Possible = possible
}

func (wa *CollapseArray2d) GetPossible(x, y int) ([]*Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("Out of bounds")
	}

	if wa.Tiles[y][x].Tile != nil {
		return nil, fmt.Errorf("Tile already collapsed")
	}

	return wa.Tiles[y][x].Possible, nil
}

func Generate2dMap(width, height int, possible []*Tile) *CollapseArray2d {
	wa := CollapseArray2d{
		Tiles: make([][]*CollapseTile, height),
	}
	for y := 0; y < height; y++ {
		wa.Tiles[y] = make([]*CollapseTile, width)
		for x := 0; x < width; x++ {
			wa.Tiles[y][x].Possible = possible
		}
	}
	return &wa
}
func (wa *CollapseArray2d) Iterate() {
	x, y := wa.findLowestEntropy()
	log.Println("Lowest entropy", x, y)
	if x == -1 && y == -1 {
		log.Println("No solution")
		return
	}
	wa.collapse(x, y)
	wa.updatePossible(x, y)
}

func (wa *CollapseArray2d) findLowestEntropy() (int, int) {
	entropy := 1000
	x, y := -1, -1
	for i, row := range wa.Tiles {
		for j, tile := range row {
			if len(tile.Possible) < entropy && len(tile.Possible) > 1 {
				entropy = len(tile.Possible)
				x, y = j, i
				log.Println("Entropy", x, y, entropy)
			}
		}
	}
	return x, y
}

func (wa *CollapseArray2d) collapse(x, y int) {
	tile := wa.Tiles[y][x]
	possible := len(tile.Possible)
	if possible == 1 {
		wa.Tiles[y][x].Tile = tile.Possible[0]
		wa.Tiles[y][x].Possible = nil
	} else {
		index := rand.Intn(possible)
		wa.Tiles[y][x].Tile = tile.Possible[index]
		wa.Tiles[y][x].Possible = nil
	}
	log.Println("Collapsed", x, y, wa.Tiles[y][x])
}

func (wa *CollapseArray2d) updatePossible(x, y int) {
	if x > 0 { // update west
		collapasedWestEdge := wa.Tiles[y][x].Tile.West
		for i, tile := range wa.Tiles[y][x-1].Possible {
			if tile.East != collapasedWestEdge {
				wa.Tiles[y][x-1].Possible = append(wa.Tiles[y][x-1].Possible[:i], wa.Tiles[y][x-1].Possible[i+1:]...)
				log.Println("Removed", tile)

			}
		}
		log.Println("Updated west, now", len(wa.Tiles[y][x-1].Possible))
	}

	if x < len(wa.Tiles[y])-1 { // update east
		collapasedEastEdge := wa.Tiles[y][x].Tile.East
		for i, tile := range wa.Tiles[y][x+1].Possible {
			if tile.West != collapasedEastEdge {
				wa.Tiles[y][x+1].Possible = append(wa.Tiles[y][x+1].Possible[:i], wa.Tiles[y][x+1].Possible[i+1:]...)
				log.Println("Removed", tile)
			}
		}
	}

	if y > 0 { // update north
		collapasedNorthEdge := wa.Tiles[y][x].Tile.North
		for i, tile := range wa.Tiles[y-1][x].Possible {
			if tile.South != collapasedNorthEdge {
				wa.Tiles[y-1][x].Possible = append(wa.Tiles[y-1][x].Possible[:i], wa.Tiles[y-1][x].Possible[i+1:]...)
				log.Println("Removed", tile)
			}
		}
	}

	if y < len(wa.Tiles)-1 { // update south
		collapasedSouthEdge := wa.Tiles[y][x].Tile.South
		for i, tile := range wa.Tiles[y+1][x].Possible {
			if tile.North != collapasedSouthEdge {
				wa.Tiles[y+1][x].Possible = append(wa.Tiles[y+1][x].Possible[:i], wa.Tiles[y+1][x].Possible[i+1:]...)
				log.Println("Removed", tile)
			}
		}
	}
}
