package wfc

import (
	"fmt"
	"fs/lib/assets"
	"math/rand"
)

type CollapseArray2d struct {
	Valid *bool
	Tile  [][]*CollapseTile
}

func (wa *CollapseArray2d) GetTile(x, y int) (*assets.Tile, error) {
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

func (wa *CollapseArray2d) Set(x, y int, tile *assets.Tile) error {
	// log.Println("Collapsed", x, y, tile.Name)
	wa.Tile[y][x].Tile = tile
	wa.Tile[y][x].PossibleTiles = nil
	err := wa.UpdatePossible(x, y)
	if err != nil {
		return err
	}
	return nil
}

func (wa *CollapseArray2d) SetPossible(x, y int, possible []*assets.Tile) {
	wa.Tile[y][x].PossibleTiles = possible
}

func (wa *CollapseArray2d) GetPossible(x, y int) ([]*assets.Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("out of bounds")
	}

	if wa.Tile[y][x].Tile != nil {
		return nil, fmt.Errorf("tile already collapsed")
	}

	return wa.Tile[y][x].PossibleTiles, nil
}

func NewCollapseArray2d(width, height int, possible []*assets.Tile) *CollapseArray2d {
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
				// log.Println("Entropy", x, y, entropy)
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
		index := rand.Intn(possible)
		err := wa.Set(x, y, tile.PossibleTiles[index])
		if err != nil {
			return err
		}
	}
	return nil
}

// Recursively update the possible tiles for the neighbors of the collapsed tile
func (wa *CollapseArray2d) UpdatePossible(x, y int) error {
	if wa.Tile[y][x].Tile != nil {
		if x > 0 { // update west
			if wa.Tile[y][x-1].Tile == nil {
				err := wa.updateNeighborPossible(x-1, y, wa.Tile[y][x].Tile.West, "east")
				if err != nil {
					return err
				}
			}
		}

		if x < wa.Width()-1 { // update east
			if wa.Tile[y][x+1].Tile == nil {
				err := wa.updateNeighborPossible(x+1, y, wa.Tile[y][x].Tile.East, "west")
				if err != nil {
					return err
				}
			}
		}

		if y > 0 { // update north
			if wa.Tile[y-1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y-1, wa.Tile[y][x].Tile.North, "south")
				if err != nil {
					return err
				}
			}
		}

		if y < wa.Height()-1 { // update south
			if wa.Tile[y+1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y+1, wa.Tile[y][x].Tile.South, "north")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (ca *CollapseArray2d) updateNeighborPossible(x, y int, otherEdge assets.Edge, compareDirection assets.Edge) error {
	newPossible := []*assets.Tile{}

	for _, tile := range ca.Tile[y][x].PossibleTiles {
		var edge assets.Edge
		stillPossible := false

		// Compare the edge of the neighbor to the edge of the collapsed tile in the compare direction
		switch compareDirection {
		case "east":
			edge = tile.East
		case "west":
			edge = tile.West
		case "north":
			edge = tile.North
		case "south":
			edge = tile.South
		}
		if edge == otherEdge {
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
