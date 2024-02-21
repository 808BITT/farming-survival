package assets

import (
	"embed"
	"fs/lib/tilemap"
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

func LoadTestTiles() []*tilemap.Tile {
	blankProb := 1.0
	roadProb := 0.001
	cornerProb := 0.001

	var probability float64

	blank := tilemap.TileTexture{
		Texture: LoadImage("tile/blank.png"),
		Edge:    tilemap.NewTileEdges("grass", "grass", "grass", "grass"),
	}
	probability = blankProb
	blankTile := tilemap.NewTile(blank, probability)

	horizontalRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/horizontal.png"),
		Edge:    tilemap.NewTileEdges("grass", "road", "grass", "road"),
	}
	probability = roadProb
	horizontalRoadTile := tilemap.NewTile(horizontalRoad, probability)

	verticalRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/vertical.png"),
		Edge:    tilemap.NewTileEdges("road", "grass", "road", "grass"),
	}
	probability = roadProb
	verticalRoadTile := tilemap.NewTile(verticalRoad, probability)

	bottomLeftRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/bottom_left.png"),
		Edge:    tilemap.NewTileEdges("grass", "grass", "road", "road"),
	}
	probability = cornerProb
	bottomLeftRoadTile := tilemap.NewTile(bottomLeftRoad, probability)

	bottomRightRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/bottom_right.png"),
		Edge:    tilemap.NewTileEdges("grass", "road", "road", "grass"),
	}
	probability = cornerProb
	bottomRightRoadTile := tilemap.NewTile(bottomRightRoad, probability)

	topLeftRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/top_left.png"),
		Edge:    tilemap.NewTileEdges("road", "grass", "grass", "road"),
	}
	probability = cornerProb
	topLeftRoadTile := tilemap.NewTile(topLeftRoad, probability)

	topRightRoad := tilemap.TileTexture{
		Texture: LoadImage("tile/top_right.png"),
		Edge:    tilemap.NewTileEdges("road", "road", "grass", "grass"),
	}
	probability = cornerProb
	topRightRoadTile := tilemap.NewTile(topRightRoad, probability)

	tiles := []*tilemap.Tile{
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
