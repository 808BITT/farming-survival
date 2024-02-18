package main

import (
	"embed"
	"fs/lib/game"
)

//go:embed assets/*
var assets embed.FS

func main() {
	e := game.NewEngine(&assets)
	e.Run()
}
