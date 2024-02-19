package assets

import (
	"embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed tile/*
var assets embed.FS

func LoadImage(path string) *ebiten.Image {
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

func LoadTestTiles() []*Tile {
	blankTile1 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile2 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile3 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile4 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile5 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile6 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile7 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile8 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile9 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	blankTile10 := NewTile(
		"blank",
		LoadImage("tile/blank.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	horizontalRoadTile1 := NewTile(
		"horizontal-road",
		LoadImage("tile/horizontal.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	horizontalRoadTile2 := NewTile(
		"horizontal-road",
		LoadImage("tile/horizontal.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	horizontalRoadTile3 := NewTile(
		"horizontal-road",
		LoadImage("tile/horizontal.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	horizontalRoadTile4 := NewTile(
		"horizontal-road",
		LoadImage("tile/horizontal.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	horizontalRoadTile5 := NewTile(
		"horizontal-road",
		LoadImage("tile/horizontal.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	verticalRoadTile1 := NewTile(
		"vertical-road",
		LoadImage("tile/vertical.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	verticalRoadTile2 := NewTile(
		"vertical-road",
		LoadImage("tile/vertical.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	verticalRoadTile3 := NewTile(
		"vertical-road",
		LoadImage("tile/vertical.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	verticalRoadTile4 := NewTile(
		"vertical-road",
		LoadImage("tile/vertical.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	verticalRoadTile5 := NewTile(
		"vertical-road",
		LoadImage("tile/vertical.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	bottomLeftRoadTile := NewTile(
		"bottom-left-road",
		LoadImage("tile/bottom_left.png"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("road"),
		NewEdge("grass"),
	)

	bottomRightRoadTile := NewTile(
		"bottom-right-road",
		LoadImage("tile/bottom_right.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("road"),
		NewEdge("road"),
	)

	topLeftRoadTile := NewTile(
		"top-left-road",
		LoadImage("tile/top_left.png"),
		NewEdge("road"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	topRightRoadTile := NewTile(
		"top-right-road",
		LoadImage("tile/top_right.png"),
		NewEdge("road"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("road"),
	)

	tiles := []*Tile{
		&blankTile1,
		&blankTile2,
		&blankTile3,
		&blankTile4,
		&blankTile5,
		&blankTile6,
		&blankTile7,
		&blankTile8,
		&blankTile9,
		&blankTile10,
		&horizontalRoadTile1,
		&horizontalRoadTile2,
		&horizontalRoadTile3,
		&horizontalRoadTile4,
		&horizontalRoadTile5,
		&verticalRoadTile1,
		&verticalRoadTile2,
		&verticalRoadTile3,
		&verticalRoadTile4,
		&verticalRoadTile5,
		&bottomLeftRoadTile,
		&bottomRightRoadTile,
		&topLeftRoadTile,
		&topRightRoadTile,
	}

	return tiles
}

func LoadTiles() []*Tile {
	// grassBottomLeftTopRight := NewTile(
	// 	"grass-dirt-path-tl-br",
	// 	LoadImage("tile/cozy-farm/grass_bottom_left_top_right.png"),
	// 	NewEdge("sand-grass"),
	// 	NewEdge("grass-sand"),
	// 	NewEdge("sand-grass"),
	// 	NewEdge("grass-sand"),
	// )

	grassBottomLeft := NewTile(
		"grass-bottom-left",
		LoadImage("tile/cozy-farm/grass_bottom_left.png"),
		NewEdge("sand"),
		NewEdge("sand"),
		NewEdge("grass-sand"),
		NewEdge("sand-grass"),
	)

	grassBottomRight := NewTile(
		"grass-bottom-right",
		LoadImage("tile/cozy-farm/grass_bottom_right.png"),
		NewEdge("sand"),
		NewEdge("sand-grass"),
		NewEdge("sand-grass"),
		NewEdge("sand"),
	)

	grassBottom := NewTile(
		"grass-bottom",
		LoadImage("tile/cozy-farm/grass_bottom.png"),
		NewEdge("sand"),
		NewEdge("sand-grass"),
		NewEdge("grass"),
		NewEdge("sand-grass"),
	)

	grassLeft := NewTile(
		"grass-left",
		LoadImage("tile/cozy-farm/grass_left.png"),
		NewEdge("grass-sand"),
		NewEdge("sand"),
		NewEdge("grass-sand"),
		NewEdge("grass"),
	)

	grassRight := NewTile(
		"grass-right",
		LoadImage("tile/cozy-farm/grass_right.png"),
		NewEdge("sand-grass"),
		NewEdge("grass"),
		NewEdge("sand-grass"),
		NewEdge("sand"),
	)

	// grassTopLeftBottomRight := NewTile(
	// 	"grass-top-left-bottom-right",
	// 	LoadImage("tile/cozy-farm/grass_top_left_bottom_right.png"),
	// 	NewEdge("grass-sand"),
	// 	NewEdge("sand-grass"),
	// 	NewEdge("sand-grass"),
	// 	NewEdge("grass-sand"),
	// )

	grassTopLeft := NewTile(
		"grass-top-left",
		LoadImage("tile/cozy-farm/grass_top_left.png"),
		NewEdge("grass-sand"),
		NewEdge("sand"),
		NewEdge("sand"),
		NewEdge("grass-sand"),
	)

	grassTopRight := NewTile(
		"grass-top-right",
		LoadImage("tile/cozy-farm/grass_top_right.png"),
		NewEdge("sand-grass"),
		NewEdge("grass-sand"),
		NewEdge("sand"),
		NewEdge("sand"),
	)

	grassTop := NewTile(
		"grass-top",
		LoadImage("tile/cozy-farm/grass_top.png"),
		NewEdge("grass"),
		NewEdge("grass-sand"),
		NewEdge("sand"),
		NewEdge("grass-sand"),
	)

	grass := NewTile(
		"grass",
		LoadImage("tile/cozy-farm/grass.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	sandBottomLeft := NewTile(
		"sand-bottom-left",
		LoadImage("tile/cozy-farm/sand_bottom_left.png"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("sand-grass"),
		NewEdge("grass-sand"),
	)

	sandBottomRight := NewTile(
		"sand-bottom-right",
		LoadImage("tile/cozy-farm/sand_bottom_right.png"),
		NewEdge("grass"),
		NewEdge("grass-sand"),
		NewEdge("grass-sand"),
		NewEdge("grass"),
	)

	sandTopLeft := NewTile(
		"sand-top-left",
		LoadImage("tile/cozy-farm/sand_top_left.png"),
		NewEdge("sand-grass"),
		NewEdge("grass"),
		NewEdge("grass"),
		NewEdge("sand-grass"),
	)

	sandTopRight := NewTile(
		"sand-top-right",
		LoadImage("tile/cozy-farm/sand_top_right.png"),
		NewEdge("grass-sand"),
		NewEdge("sand-grass"),
		NewEdge("grass"),
		NewEdge("grass"),
	)

	sand := NewTile(
		"sand",
		LoadImage("tile/cozy-farm/sand.png"),
		NewEdge("sand"),
		NewEdge("sand"),
		NewEdge("sand"),
		NewEdge("sand"),
	)

	tiles := []*Tile{
		&grassBottomLeft,
		&grassBottomRight,
		&grassBottom,
		&grassLeft,
		&grassRight,
		&grassTopLeft,
		&grassTopRight,
		&grassTop,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&grass,
		&sandBottomLeft,
		&sandBottomRight,
		&sandTopLeft,
		&sandTopRight,
		&sand,
	}

	return tiles
}
