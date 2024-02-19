package wfc

import (
	"fmt"
	"fs/lib/assets"
	"math/rand"
)

type CollapseArray2d struct {
	Valid *bool
	Tiles [][]*CollapseTile
}

func (wa *CollapseArray2d) GetTile(x, y int) (*assets.Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("out of bounds")
	}

	if wa.Tiles[y][x].Tile == nil {
		return nil, fmt.Errorf("tile not collapsed yet")
	}

	return wa.Tiles[y][x].Tile, nil
}

func (wa *CollapseArray2d) Width() int {
	return len(wa.Tiles[0])
}

func (wa *CollapseArray2d) Height() int {
	return len(wa.Tiles)
}

func (wa *CollapseArray2d) Set(x, y int, tile *assets.Tile) error {
	// log.Println("Collapsed", x, y, tile.Name)
	wa.Tiles[y][x].Tile = tile
	wa.Tiles[y][x].Possible = nil
	err := wa.UpdatePossible(x, y)
	if err != nil {
		return err
	}
	return nil
}

func (wa *CollapseArray2d) SetPossible(x, y int, possible []*assets.Tile) {
	wa.Tiles[y][x].Possible = possible
}

func (wa *CollapseArray2d) GetPossible(x, y int) ([]*assets.Tile, error) {
	if x < 0 || x >= wa.Width() || y < 0 || y >= wa.Height() {
		return nil, fmt.Errorf("out of bounds")
	}

	if wa.Tiles[y][x].Tile != nil {
		return nil, fmt.Errorf("tile already collapsed")
	}

	return wa.Tiles[y][x].Possible, nil
}

func NewCollapseArray2d(width, height int, possible []*assets.Tile) *CollapseArray2d {
	wa := CollapseArray2d{
		Tiles: make([][]*CollapseTile, height),
		Valid: nil,
	}
	for y := 0; y < height; y++ {
		wa.Tiles[y] = make([]*CollapseTile, width)
		for x := 0; x < width; x++ {
			wa.Tiles[y][x] = &CollapseTile{
				Tile:     nil,
				Possible: possible,
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
	for i, row := range wa.Tiles {
		for j, tile := range row {
			if len(tile.Possible) < entropy && len(tile.Possible) > 1 {
				entropy = len(tile.Possible)
				x, y = j, i
				// log.Println("Entropy", x, y, entropy)
			}
		}
	}
	return x, y
}

// Collapse the tile at x, y to a random possible tile
func (wa *CollapseArray2d) Collapse(x, y int) error {
	tile := wa.Tiles[y][x]
	possible := len(tile.Possible)
	if possible == 1 {
		err := wa.Set(x, y, tile.Possible[0])
		if err != nil {
			return err
		}
	} else {
		index := rand.Intn(possible)
		err := wa.Set(x, y, tile.Possible[index])
		if err != nil {
			return err
		}
	}
	return nil
}

// Recursively update the possible tiles for the neighbors of the collapsed tile
func (wa *CollapseArray2d) UpdatePossible(x, y int) error {
	if wa.Tiles[y][x].Tile != nil {
		if x > 0 { // update west
			if wa.Tiles[y][x-1].Tile == nil {
				err := wa.updateNeighborPossible(x-1, y, wa.Tiles[y][x].Tile.West, "east")
				if err != nil {
					return err
				}
			}
		}

		if x < wa.Width()-1 { // update east
			if wa.Tiles[y][x+1].Tile == nil {
				err := wa.updateNeighborPossible(x+1, y, wa.Tiles[y][x].Tile.East, "west")
				if err != nil {
					return err
				}
			}
		}

		if y > 0 { // update north
			if wa.Tiles[y-1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y-1, wa.Tiles[y][x].Tile.North, "south")
				if err != nil {
					return err
				}
			}
		}

		if y < wa.Height()-1 { // update south
			if wa.Tiles[y+1][x].Tile == nil {
				err := wa.updateNeighborPossible(x, y+1, wa.Tiles[y][x].Tile.South, "north")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (wa *CollapseArray2d) updateNeighborPossible(x, y int, otherEdge assets.Edge, compareDirection assets.Edge) error {
	var edge assets.Edge
	newPossible := []*assets.Tile{}

	for _, tile := range wa.Tiles[y][x].Possible {
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
			newPossible = append(newPossible, tile)
		}
	}
	if len(newPossible) == 0 {
		return fmt.Errorf("no possible tiles")
	}

	if len(newPossible) == 1 {
		wa.Set(x, y, newPossible[0])
	} else {
		wa.Tiles[y][x].Possible = newPossible
	}
	// log.Println("Updated", x, y, wa.Tiles[y][x].Possible)
	wa.UpdatePossible(x, y) // Be cautious with recursion; ensure it won't cause infinite loops
	return nil
}
