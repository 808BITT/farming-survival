package tests

import (
	"fs/lib/assets"
	"fs/lib/wfc"
	"strings"
	"testing"
)

func TestBlank(t *testing.T) {
	tiles := assets.LoadTestTiles()

	ca := wfc.NewCollapseArray2d(2, 1, tiles)

	var blankTile *assets.Tile
	for _, tile := range tiles {
		if tile.Name == "blank" {
			blankTile = tile
		}
	}

	ca.Set(0, 0, blankTile)

	tile, err := ca.GetTile(0, 0)
	if err != nil {
		t.Error(err)
	}

	if tile != blankTile {
		t.Errorf("expected %v, got %v", blankTile, tile)
	}

	_, err = ca.GetTile(1, 0)
	if err == nil {
		t.Error("expected error, got nil")
	}

	_, err = ca.GetTile(0, 1)
	if err == nil {
		t.Error("expected error, got nil")
	}

	possible, err := ca.GetPossible(0, 0)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if possible != nil {
		t.Errorf("expected nil, got %v", tiles)
	}

	possible, err = ca.GetPossible(1, 0)
	if err != nil {
		t.Error("expected nil, got", possible)
	}
	possibleStrings := make([]string, len(possible))
	for i, tile := range possible {
		possibleStrings[i] = tile.Name
	}

	if len(possible) != 4 {
		t.Errorf("expected 4, got %v", len(possible))
	}

	if !strings.Contains(strings.Join(possibleStrings, " "), "blank") {
		t.Errorf("expected blank, got %v", possibleStrings)
	}
	if !strings.Contains(strings.Join(possibleStrings, " "), "vertical road") {
		t.Errorf("expected vertical road, got %v", possibleStrings)
	}
	if !strings.Contains(strings.Join(possibleStrings, " "), "bottom right road") {
		t.Errorf("expected bottom right road, got %v", possibleStrings)
	}
	if !strings.Contains(strings.Join(possibleStrings, " "), "top right road") {
		t.Errorf("expected top right road, got %v", possibleStrings)
	}
}
