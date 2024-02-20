package wfc

import (
	"fmt"
	"fs/lib/tilemap"
	"math/rand"
)

type CollapseArray2d struct {
	Valid *bool
	Tile  [][]*CollapseTile
}

func (wa *CollapseArray2d) IsValid() bool {
	return *wa.Valid
}

func (wa *CollapseArray2d) IsCollapsed() bool {
	if wa.Valid == nil {
		valid := true
		for _, row := range wa.Tile {
			for _, tile := range row {
				if tile.Tile == nil {
					if len(tile.PossibleTiles) == 0 {
						valid = false
						break
					}
				}
			}
		}
		wa.Valid = &valid
	}
	return *wa.Valid
}

func (wa *CollapseArray2d) GetTile(x, y int) (*tilemap.Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("out of bounds")
	}

	if wa.Tile[y][x].Tile == nil {
		return nil, fmt.Errorf("tile not collapsed yet")
	}

	return wa.Tile[y][x].Tile, nil
}

func (wa *CollapseArray2d) Width() int {
	return len(wa.Tile[0])
}

func (wa *CollapseArray2d) Height() int {
	return len(wa.Tile)
}

func (wa *CollapseArray2d) Set(x, y int, tile *tilemap.Tile) error {
	// log.Println("Collapsed", x, y, tile.Name)
	wa.Tile[y][x].Tile = tile
	wa.Tile[y][x].PossibleTiles = nil
	err := wa.UpdatePossible(x, y)
	if err != nil {
		return err
	}
	return nil
}

func (wa *CollapseArray2d) SetPossible(x, y int, possible []*tilemap.Tile) {
	wa.Tile[y][x].PossibleTiles = possible
}

func (wa *CollapseArray2d) GetPossible(x, y int) ([]*tilemap.Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("out of bounds")
	}

	if wa.Tile[y][x].Tile != nil {
		return nil, fmt.Errorf("tile already collapsed")
	}

	return wa.Tile[y][x].PossibleTiles, nil
}

func NewCollapseArray2d(width, height int, possible []*tilemap.Tile) *CollapseArray2d {
	wa := CollapseArray2d{
		Tile:  make([][]*CollapseTile, height),
		Valid: nil,
	}
	for y := 0; y < height; y++ {
		wa.Tile[y] = make([]*CollapseTile, width)
		for x := 0; x < width; x++ {
			wa.Tile[y][x] = &CollapseTile{
				Tile:          nil,
				PossibleTiles: possible,
			}
		}
	}
	return &wa
}

// Iterate through the wave array and collapse the tiles with the lowest entropy
func (wa *CollapseArray2d) Iterate() (int, int, error) {
	x, y := wa.FindLowestEntropy()
	if x != -1 && y != -1 {
		err := wa.Collapse(x, y)
		if err != nil {
			return x, y, err
		}
	}
	return x, y, nil
}

// Find the tile with the lowest entropy
func (wa *CollapseArray2d) FindLowestEntropy() (int, int) {
	entropy := 1000
	x, y := -1, -1
	for i, row := range wa.Tile {
		for j, tile := range row {
			if len(tile.PossibleTiles) < entropy && len(tile.PossibleTiles) > 1 {
				entropy = len(tile.PossibleTiles)
				x, y = j, i
			}
		}
	}
	return x, y
}

// Collapse the tile at x, y to a random possible tile
func (wa *CollapseArray2d) Collapse(x, y int) error {
	tile := wa.Tile[y][x]
	possible := len(tile.PossibleTiles)
	if possible == 1 {
		err := wa.Set(x, y, tile.PossibleTiles[0])
		if err != nil {
			return err
		}
	} else {
		// i need to use the probability of the tiles to determine the random tile
		randomTile := randomTileFromProbablilities(tile.PossibleTiles)
		err := wa.Set(x, y, randomTile)
		if err != nil {
			return err
		}
	}
	return nil
}

func randomTileFromProbablilities(tiles []*tilemap.Tile) *tilemap.Tile {
	p := 0.0
	for _, t := range tiles {
		p += t.Probability
	}
	r := rand.Float64() * p
	shuffledTiles := shuffleTiles(tiles)
	for _, t := range shuffledTiles {
		r -= t.Probability
		if r <= 0 {
			return t
		}
	}
	return tiles[0]
}

func shuffleTiles(tiles []*tilemap.Tile) []*tilemap.Tile {
	for i := range tiles {
		j := rand.Intn(i + 1)
		tiles[i], tiles[j] = tiles[j], tiles[i]
	}
	return tiles
}

// Recursively update the possible tiles for the neighbors of the collapsed tile
func (wa *CollapseArray2d) UpdatePossible(x, y int) error {
	if wa.Tile[y][x].Tile != nil {
		if x > 0 { // update west
			if wa.Tile[y][x-1].Tile == nil {
				err := wa.updateNeighborPossible(x-1, y, wa.Tile[y][x].Tile.Type.Edge.West(), "east")
				if err != nil {
					return err
				}
			}
		}

		if x < wa.Width()-1 { // update east
			if wa.Tile[y][x+1].Tile == nil {
				err := wa.updateNeighborPossible(x+1, y, wa.Tile[y][x].Tile.Type.Edge.East(), "west")
				if err != nil {
					return err
				}
			}
		}

		if y > 0 { // update north
			if wa.Tile[y-1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y-1, wa.Tile[y][x].Tile.Type.Edge.North(), "south")
				if err != nil {
					return err
				}
			}
		}

		if y < wa.Height()-1 { // update south
			if wa.Tile[y+1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y+1, wa.Tile[y][x].Tile.Type.Edge.South(), "north")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (ca *CollapseArray2d) updateNeighborPossible(x, y int, otherEdge string, neighborDir string) error {
	newPossible := []*tilemap.Tile{}

	for _, tile := range ca.Tile[y][x].PossibleTiles {
		var matchingEdge string
		stillPossible := false

		// Compare the edge of the neighbor to the edge of the collapsed tile in the compare direction
		switch neighborDir {
		case "east":
			matchingEdge = tile.Type.Edge.East()
		case "west":
			matchingEdge = tile.Type.Edge.West()
		case "north":
			matchingEdge = tile.Type.Edge.North()
		case "south":
			matchingEdge = tile.Type.Edge.South()
		}
		if matchingEdge == otherEdge {
			stillPossible = true
		}

		if stillPossible {
			newPossible = append(newPossible, tile)
		}
	}
	if len(newPossible) == 0 {
		return fmt.Errorf("no possible tiles")
	}

	if len(newPossible) == 1 {
		ca.Set(x, y, newPossible[0])
	} else if len(newPossible) != len(ca.Tile[y][x].PossibleTiles) {
		ca.Tile[y][x].PossibleTiles = newPossible
		ca.UpdatePossible(x, y) // Recursively update the neighbor neighbors since their possible tiles have changed as well
	}

	return nil
}
