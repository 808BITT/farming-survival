package main

import (
	"fs/lib/game"
)

func main() {
	e := game.NewEngine(1920, 1080)
	e.Run()
}
